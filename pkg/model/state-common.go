package model

import "errors"

type RetryDefinition struct {
	MaxAttempts int     `yaml:"max_attempts"`
	Delay       string  `yaml:"delay,omitempty"`
	Multiplier  float64 `yaml:"multiplier,omitempty"`
}

func (o *RetryDefinition) Validate() error {
	if o == nil {
		return nil
	}

	if o.MaxAttempts == 0 {
		return errors.New("maxAttempts required to be more than 0")
	}

	if o.Delay != "" && !isISO8601(o.Delay) {
		return errors.New("delay is not a ISO8601 string")
	}

	return nil
}

type ErrorDefinition struct {
	Error      string           `yaml:"error"`
	Retry      *RetryDefinition `yaml:"retry,omitempty"`
	Transition string           `yaml:"transition,omitempty"`
}

func (o *ErrorDefinition) Validate() error {
	if o.Error == "" {
		return errors.New("error required")
	}

	if err := o.Retry.Validate(); err != nil {
		return err
	}

	return nil
}

type State interface {
	GetID() string
	GetType() StateType
	Validate() error
	ErrorDefinitions() []ErrorDefinition
	GetTransitions() []string
	getTransitions() map[string]string
}

type ConsumeEventDefinition struct {
	Type    string                 `yaml:"type"`
	Context map[string]interface{} `yaml:"context,omitempty"`
}

func (o *ConsumeEventDefinition) Validate() error {
	if o.Type == "" {
		return errors.New("type required")
	}

	return nil

}

type ProduceEventDefinition struct {
	Type    string                 `yaml:"type,omitempty"`
	Source  string                 `yaml:"source,omitempty"`
	Data    string                 `yaml:"data"`
	Context map[string]interface{} `yaml:"context"`
}

func (o *ProduceEventDefinition) Validate() error {
	if o.Source == "" {
		return errors.New("source required")
	}

	if o.Type == "" {
		return errors.New("type required")
	}

	return nil

}

type StateCommon struct {
	ID   string    `yaml:"id"`
	Type StateType `yaml:"type"`
}

func (o *StateCommon) GetType() StateType {
	return o.Type
}

func (o *StateCommon) commonValidate() error {
	if o.ID == "" {
		return errors.New("id required")
	}

	return nil
}
