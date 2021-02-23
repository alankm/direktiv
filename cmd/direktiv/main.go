package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/vorteil/direktiv/pkg/direktiv"
	"github.com/vorteil/direktiv/pkg/dlog"
	"github.com/vorteil/direktiv/pkg/dlog/db"
	"github.com/vorteil/direktiv/pkg/dlog/dummy"
)

var (
	debug                  bool
	serverType, configFile string
)

var rootCmd = &cobra.Command{
	Use:   "direktiv",
	Short: "direktiv is a serverless, container workflow engine.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
			formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
				FullTimestamp: true,
			}}
			formatter.Line = true
			logrus.SetFormatter(&formatter)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		defer func() {
			// just in case, stop DNS server
			pv, err := ioutil.ReadFile("/proc/version")
			if err == nil {
				// this is a vorteil machine, so we press poweroff
				if strings.Contains(string(pv), "#vorteil") {
					logrus.Printf("vorteil machine, powering off")

					if err := exec.Command("/sbin/poweroff").Run(); err != nil {
						fmt.Println("error shutting down:", err)
					}

				}
			}

		}()

		c, err := direktiv.ReadConfig(configFile)
		if err != nil {
			logrus.Errorf("Failed to initialize server: %v", err)
			os.Exit(1)
		}

		server, err := direktiv.NewWorkflowServer(c, serverType)
		if err != nil {
			log.Fatalf("failed to create server: %v", err)
		}

		var logger dlog.Log

		switch c.InstanceLogging.Driver {
		case "database":
			logrus.Info("creating logger type database")
			l, err := db.NewLogger(c.Database.DB)
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}
			defer l.CloseConnection()
			logger = l
		default:
			logrus.Info("creating logger type default")
			logger, err = dummy.NewLogger()
		}

		server.SetInstanceLogger(logger)

		go func() {
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
			<-sig
			server.Stop()
			<-sig
			server.Kill()
		}()

		go func() {
			server.Run()
			if err != nil {
				log.Fatalf("unable to start server: %v", err)
			}
		}()

		<-server.Lifeline()

		log.Printf("server stopped\n")

		return

	},
}

func main() {

	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enabled debug output")
	rootCmd.Flags().StringVarP(&serverType, "type", "t", "wf", "can run worfklow engine (wfo), container runner (wfr) or both (wf)")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "configuration file to use")

	err := rootCmd.Execute()
	if err != nil {
		logrus.Errorf("%v", err)
		os.Exit(1)
	}

}
