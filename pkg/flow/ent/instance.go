// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/pkg/flow/ent/instance"
	"github.com/vorteil/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/vorteil/direktiv/pkg/flow/ent/namespace"
	"github.com/vorteil/direktiv/pkg/flow/ent/revision"
	"github.com/vorteil/direktiv/pkg/flow/ent/workflow"
)

// Instance is the model entity for the Instance schema.
type Instance struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt time.Time `json:"end_at,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// As holds the value of the "as" field.
	As string `json:"as,omitempty"`
	// ErrorCode holds the value of the "errorCode" field.
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorMessage holds the value of the "errorMessage" field.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceQuery when eager-loading is set.
	Edges               InstanceEdges `json:"edges"`
	namespace_instances *uuid.UUID
	revision_instances  *uuid.UUID
	workflow_instances  *uuid.UUID
}

// InstanceEdges holds the relations/edges for other nodes in the graph.
type InstanceEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// Revision holds the value of the revision edge.
	Revision *Revision `json:"revision,omitempty"`
	// Logs holds the value of the logs edge.
	Logs []*LogMsg `json:"logs,omitempty"`
	// Vars holds the value of the vars edge.
	Vars []*VarRef `json:"vars,omitempty"`
	// Runtime holds the value of the runtime edge.
	Runtime *InstanceRuntime `json:"runtime,omitempty"`
	// Children holds the value of the children edge.
	Children []*InstanceRuntime `json:"children,omitempty"`
	// Eventlisteners holds the value of the eventlisteners edge.
	Eventlisteners []*Events `json:"eventlisteners,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// The edge namespace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[1] {
		if e.Workflow == nil {
			// The edge workflow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// RevisionOrErr returns the Revision value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) RevisionOrErr() (*Revision, error) {
	if e.loadedTypes[2] {
		if e.Revision == nil {
			// The edge revision was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: revision.Label}
		}
		return e.Revision, nil
	}
	return nil, &NotLoadedError{edge: "revision"}
}

// LogsOrErr returns the Logs value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) LogsOrErr() ([]*LogMsg, error) {
	if e.loadedTypes[3] {
		return e.Logs, nil
	}
	return nil, &NotLoadedError{edge: "logs"}
}

// VarsOrErr returns the Vars value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) VarsOrErr() ([]*VarRef, error) {
	if e.loadedTypes[4] {
		return e.Vars, nil
	}
	return nil, &NotLoadedError{edge: "vars"}
}

// RuntimeOrErr returns the Runtime value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) RuntimeOrErr() (*InstanceRuntime, error) {
	if e.loadedTypes[5] {
		if e.Runtime == nil {
			// The edge runtime was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instanceruntime.Label}
		}
		return e.Runtime, nil
	}
	return nil, &NotLoadedError{edge: "runtime"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) ChildrenOrErr() ([]*InstanceRuntime, error) {
	if e.loadedTypes[6] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// EventlistenersOrErr returns the Eventlisteners value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) EventlistenersOrErr() ([]*Events, error) {
	if e.loadedTypes[7] {
		return e.Eventlisteners, nil
	}
	return nil, &NotLoadedError{edge: "eventlisteners"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Instance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case instance.FieldStatus, instance.FieldAs, instance.FieldErrorCode, instance.FieldErrorMessage:
			values[i] = new(sql.NullString)
		case instance.FieldCreatedAt, instance.FieldUpdatedAt, instance.FieldEndAt:
			values[i] = new(sql.NullTime)
		case instance.FieldID:
			values[i] = new(uuid.UUID)
		case instance.ForeignKeys[0]: // namespace_instances
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case instance.ForeignKeys[1]: // revision_instances
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case instance.ForeignKeys[2]: // workflow_instances
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Instance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Instance fields.
func (i *Instance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case instance.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case instance.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case instance.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case instance.FieldEndAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[j])
			} else if value.Valid {
				i.EndAt = value.Time
			}
		case instance.FieldStatus:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[j])
			} else if value.Valid {
				i.Status = value.String
			}
		case instance.FieldAs:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field as", values[j])
			} else if value.Valid {
				i.As = value.String
			}
		case instance.FieldErrorCode:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errorCode", values[j])
			} else if value.Valid {
				i.ErrorCode = value.String
			}
		case instance.FieldErrorMessage:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errorMessage", values[j])
			} else if value.Valid {
				i.ErrorMessage = value.String
			}
		case instance.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_instances", values[j])
			} else if value.Valid {
				i.namespace_instances = new(uuid.UUID)
				*i.namespace_instances = *value.S.(*uuid.UUID)
			}
		case instance.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field revision_instances", values[j])
			} else if value.Valid {
				i.revision_instances = new(uuid.UUID)
				*i.revision_instances = *value.S.(*uuid.UUID)
			}
		case instance.ForeignKeys[2]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_instances", values[j])
			} else if value.Valid {
				i.workflow_instances = new(uuid.UUID)
				*i.workflow_instances = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the Instance entity.
func (i *Instance) QueryNamespace() *NamespaceQuery {
	return (&InstanceClient{config: i.config}).QueryNamespace(i)
}

// QueryWorkflow queries the "workflow" edge of the Instance entity.
func (i *Instance) QueryWorkflow() *WorkflowQuery {
	return (&InstanceClient{config: i.config}).QueryWorkflow(i)
}

// QueryRevision queries the "revision" edge of the Instance entity.
func (i *Instance) QueryRevision() *RevisionQuery {
	return (&InstanceClient{config: i.config}).QueryRevision(i)
}

// QueryLogs queries the "logs" edge of the Instance entity.
func (i *Instance) QueryLogs() *LogMsgQuery {
	return (&InstanceClient{config: i.config}).QueryLogs(i)
}

// QueryVars queries the "vars" edge of the Instance entity.
func (i *Instance) QueryVars() *VarRefQuery {
	return (&InstanceClient{config: i.config}).QueryVars(i)
}

// QueryRuntime queries the "runtime" edge of the Instance entity.
func (i *Instance) QueryRuntime() *InstanceRuntimeQuery {
	return (&InstanceClient{config: i.config}).QueryRuntime(i)
}

// QueryChildren queries the "children" edge of the Instance entity.
func (i *Instance) QueryChildren() *InstanceRuntimeQuery {
	return (&InstanceClient{config: i.config}).QueryChildren(i)
}

// QueryEventlisteners queries the "eventlisteners" edge of the Instance entity.
func (i *Instance) QueryEventlisteners() *EventsQuery {
	return (&InstanceClient{config: i.config}).QueryEventlisteners(i)
}

// Update returns a builder for updating this Instance.
// Note that you need to call Instance.Unwrap() before calling this method if this Instance
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Instance) Update() *InstanceUpdateOne {
	return (&InstanceClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Instance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Instance) Unwrap() *Instance {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Instance is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Instance) String() string {
	var builder strings.Builder
	builder.WriteString("Instance(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", end_at=")
	builder.WriteString(i.EndAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(i.Status)
	builder.WriteString(", as=")
	builder.WriteString(i.As)
	builder.WriteString(", errorCode=")
	builder.WriteString(i.ErrorCode)
	builder.WriteString(", errorMessage=")
	builder.WriteString(i.ErrorMessage)
	builder.WriteByte(')')
	return builder.String()
}

// Instances is a parsable slice of Instance.
type Instances []*Instance

func (i Instances) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}