// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vorteil/direktiv/ent/workflowevents"
	"github.com/vorteil/direktiv/ent/workfloweventswait"
)

// WorkflowEventsWait is the model entity for the WorkflowEventsWait schema.
type WorkflowEventsWait struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Events holds the value of the "events" field.
	Events map[string]interface{} `json:"events,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowEventsWaitQuery when eager-loading is set.
	Edges                        WorkflowEventsWaitEdges `json:"edges"`
	workflow_events_wfeventswait *int
}

// WorkflowEventsWaitEdges holds the relations/edges for other nodes in the graph.
type WorkflowEventsWaitEdges struct {
	// Workflowevent holds the value of the workflowevent edge.
	Workflowevent *WorkflowEvents `json:"workflowevent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkfloweventOrErr returns the Workflowevent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowEventsWaitEdges) WorkfloweventOrErr() (*WorkflowEvents, error) {
	if e.loadedTypes[0] {
		if e.Workflowevent == nil {
			// The edge workflowevent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflowevents.Label}
		}
		return e.Workflowevent, nil
	}
	return nil, &NotLoadedError{edge: "workflowevent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowEventsWait) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workfloweventswait.FieldEvents:
			values[i] = new([]byte)
		case workfloweventswait.FieldID:
			values[i] = new(sql.NullInt64)
		case workfloweventswait.ForeignKeys[0]: // workflow_events_wfeventswait
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkflowEventsWait", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowEventsWait fields.
func (wew *WorkflowEventsWait) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workfloweventswait.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wew.ID = int(value.Int64)
		case workfloweventswait.FieldEvents:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field events", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wew.Events); err != nil {
					return fmt.Errorf("unmarshal field events: %w", err)
				}
			}
		case workfloweventswait.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workflow_events_wfeventswait", value)
			} else if value.Valid {
				wew.workflow_events_wfeventswait = new(int)
				*wew.workflow_events_wfeventswait = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWorkflowevent queries the "workflowevent" edge of the WorkflowEventsWait entity.
func (wew *WorkflowEventsWait) QueryWorkflowevent() *WorkflowEventsQuery {
	return (&WorkflowEventsWaitClient{config: wew.config}).QueryWorkflowevent(wew)
}

// Update returns a builder for updating this WorkflowEventsWait.
// Note that you need to call WorkflowEventsWait.Unwrap() before calling this method if this WorkflowEventsWait
// was returned from a transaction, and the transaction was committed or rolled back.
func (wew *WorkflowEventsWait) Update() *WorkflowEventsWaitUpdateOne {
	return (&WorkflowEventsWaitClient{config: wew.config}).UpdateOne(wew)
}

// Unwrap unwraps the WorkflowEventsWait entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wew *WorkflowEventsWait) Unwrap() *WorkflowEventsWait {
	tx, ok := wew.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkflowEventsWait is not a transactional entity")
	}
	wew.config.driver = tx.drv
	return wew
}

// String implements the fmt.Stringer.
func (wew *WorkflowEventsWait) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowEventsWait(")
	builder.WriteString(fmt.Sprintf("id=%v", wew.ID))
	builder.WriteString(", events=")
	builder.WriteString(fmt.Sprintf("%v", wew.Events))
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowEventsWaits is a parsable slice of WorkflowEventsWait.
type WorkflowEventsWaits []*WorkflowEventsWait

func (wew WorkflowEventsWaits) config(cfg config) {
	for _i := range wew {
		wew[_i].config = cfg
	}
}
