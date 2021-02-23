// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/pkg/secrets/ent/bucketsecret"
	"github.com/vorteil/direktiv/pkg/secrets/ent/predicate"
)

// BucketSecretDelete is the builder for deleting a BucketSecret entity.
type BucketSecretDelete struct {
	config
	hooks    []Hook
	mutation *BucketSecretMutation
}

// Where adds a new predicate to the BucketSecretDelete builder.
func (bsd *BucketSecretDelete) Where(ps ...predicate.BucketSecret) *BucketSecretDelete {
	bsd.mutation.predicates = append(bsd.mutation.predicates, ps...)
	return bsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bsd *BucketSecretDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bsd.hooks) == 0 {
		affected, err = bsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BucketSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bsd.mutation = mutation
			affected, err = bsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bsd.hooks) - 1; i >= 0; i-- {
			mut = bsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsd *BucketSecretDelete) ExecX(ctx context.Context) int {
	n, err := bsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bsd *BucketSecretDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bucketsecret.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bucketsecret.FieldID,
			},
		},
	}
	if ps := bsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bsd.driver, _spec)
}

// BucketSecretDeleteOne is the builder for deleting a single BucketSecret entity.
type BucketSecretDeleteOne struct {
	bsd *BucketSecretDelete
}

// Exec executes the deletion query.
func (bsdo *BucketSecretDeleteOne) Exec(ctx context.Context) error {
	n, err := bsdo.bsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bucketsecret.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bsdo *BucketSecretDeleteOne) ExecX(ctx context.Context) {
	bsdo.bsd.ExecX(ctx)
}
