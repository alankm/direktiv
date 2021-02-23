// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/ent/server"
)

// ServerCreate is the builder for creating a Server entity.
type ServerCreate struct {
	config
	mutation *ServerMutation
	hooks    []Hook
}

// SetIP sets the "ip" field.
func (sc *ServerCreate) SetIP(s string) *ServerCreate {
	sc.mutation.SetIP(s)
	return sc
}

// SetExtIP sets the "extIP" field.
func (sc *ServerCreate) SetExtIP(s string) *ServerCreate {
	sc.mutation.SetExtIP(s)
	return sc
}

// SetNatsPort sets the "natsPort" field.
func (sc *ServerCreate) SetNatsPort(i int) *ServerCreate {
	sc.mutation.SetNatsPort(i)
	return sc
}

// SetMemberPort sets the "memberPort" field.
func (sc *ServerCreate) SetMemberPort(i int) *ServerCreate {
	sc.mutation.SetMemberPort(i)
	return sc
}

// SetAdded sets the "added" field.
func (sc *ServerCreate) SetAdded(t time.Time) *ServerCreate {
	sc.mutation.SetAdded(t)
	return sc
}

// SetNillableAdded sets the "added" field if the given value is not nil.
func (sc *ServerCreate) SetNillableAdded(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetAdded(*t)
	}
	return sc
}

// Mutation returns the ServerMutation object of the builder.
func (sc *ServerCreate) Mutation() *ServerMutation {
	return sc.mutation
}

// Save creates the Server in the database.
func (sc *ServerCreate) Save(ctx context.Context) (*Server, error) {
	var (
		err  error
		node *Server
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServerCreate) SaveX(ctx context.Context) *Server {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sc *ServerCreate) defaults() {
	if _, ok := sc.mutation.Added(); !ok {
		v := server.DefaultAdded()
		sc.mutation.SetAdded(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServerCreate) check() error {
	if _, ok := sc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New("ent: missing required field \"ip\"")}
	}
	if v, ok := sc.mutation.IP(); ok {
		if err := server.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf("ent: validator failed for field \"ip\": %w", err)}
		}
	}
	if _, ok := sc.mutation.ExtIP(); !ok {
		return &ValidationError{Name: "extIP", err: errors.New("ent: missing required field \"extIP\"")}
	}
	if v, ok := sc.mutation.ExtIP(); ok {
		if err := server.ExtIPValidator(v); err != nil {
			return &ValidationError{Name: "extIP", err: fmt.Errorf("ent: validator failed for field \"extIP\": %w", err)}
		}
	}
	if _, ok := sc.mutation.NatsPort(); !ok {
		return &ValidationError{Name: "natsPort", err: errors.New("ent: missing required field \"natsPort\"")}
	}
	if _, ok := sc.mutation.MemberPort(); !ok {
		return &ValidationError{Name: "memberPort", err: errors.New("ent: missing required field \"memberPort\"")}
	}
	if _, ok := sc.mutation.Added(); !ok {
		return &ValidationError{Name: "added", err: errors.New("ent: missing required field \"added\"")}
	}
	return nil
}

func (sc *ServerCreate) sqlSave(ctx context.Context) (*Server, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ServerCreate) createSpec() (*Server, *sqlgraph.CreateSpec) {
	var (
		_node = &Server{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: server.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: server.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: server.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := sc.mutation.ExtIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: server.FieldExtIP,
		})
		_node.ExtIP = value
	}
	if value, ok := sc.mutation.NatsPort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldNatsPort,
		})
		_node.NatsPort = value
	}
	if value, ok := sc.mutation.MemberPort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldMemberPort,
		})
		_node.MemberPort = value
	}
	if value, ok := sc.mutation.Added(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldAdded,
		})
		_node.Added = value
	}
	return _node, _spec
}

// ServerCreateBulk is the builder for creating many Server entities in bulk.
type ServerCreateBulk struct {
	config
	builders []*ServerCreate
}

// Save creates the Server entities in the database.
func (scb *ServerCreateBulk) Save(ctx context.Context) ([]*Server, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Server, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServerCreateBulk) SaveX(ctx context.Context) []*Server {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
