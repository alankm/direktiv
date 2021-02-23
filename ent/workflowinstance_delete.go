// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/ent/predicate"
	"github.com/vorteil/direktiv/ent/workflowinstance"
)

// WorkflowInstanceDelete is the builder for deleting a WorkflowInstance entity.
type WorkflowInstanceDelete struct {
	config
	hooks    []Hook
	mutation *WorkflowInstanceMutation
}

// Where adds a new predicate to the WorkflowInstanceDelete builder.
func (wid *WorkflowInstanceDelete) Where(ps ...predicate.WorkflowInstance) *WorkflowInstanceDelete {
	wid.mutation.predicates = append(wid.mutation.predicates, ps...)
	return wid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wid *WorkflowInstanceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wid.hooks) == 0 {
		affected, err = wid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowInstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wid.mutation = mutation
			affected, err = wid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wid.hooks) - 1; i >= 0; i-- {
			mut = wid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wid *WorkflowInstanceDelete) ExecX(ctx context.Context) int {
	n, err := wid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wid *WorkflowInstanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workflowinstance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflowinstance.FieldID,
			},
		},
	}
	if ps := wid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wid.driver, _spec)
}

// WorkflowInstanceDeleteOne is the builder for deleting a single WorkflowInstance entity.
type WorkflowInstanceDeleteOne struct {
	wid *WorkflowInstanceDelete
}

// Exec executes the deletion query.
func (wido *WorkflowInstanceDeleteOne) Exec(ctx context.Context) error {
	n, err := wido.wid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workflowinstance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wido *WorkflowInstanceDeleteOne) ExecX(ctx context.Context) {
	wido.wid.ExecX(ctx)
}
