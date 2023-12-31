// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carenaggio/hermes/ent/system"
	"github.com/google/uuid"
)

// SystemCreate is the builder for creating a System entity.
type SystemCreate struct {
	config
	mutation *SystemMutation
	hooks    []Hook
}

// SetSystemID sets the "system_id" field.
func (sc *SystemCreate) SetSystemID(u uuid.UUID) *SystemCreate {
	sc.mutation.SetSystemID(u)
	return sc
}

// SetPublicKey sets the "public_key" field.
func (sc *SystemCreate) SetPublicKey(b []byte) *SystemCreate {
	sc.mutation.SetPublicKey(b)
	return sc
}

// SetApproved sets the "approved" field.
func (sc *SystemCreate) SetApproved(b bool) *SystemCreate {
	sc.mutation.SetApproved(b)
	return sc
}

// SetLastLogin sets the "last_login" field.
func (sc *SystemCreate) SetLastLogin(i int64) *SystemCreate {
	sc.mutation.SetLastLogin(i)
	return sc
}

// Mutation returns the SystemMutation object of the builder.
func (sc *SystemCreate) Mutation() *SystemMutation {
	return sc.mutation
}

// Save creates the System in the database.
func (sc *SystemCreate) Save(ctx context.Context) (*System, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SystemCreate) SaveX(ctx context.Context) *System {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SystemCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SystemCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SystemCreate) check() error {
	if _, ok := sc.mutation.SystemID(); !ok {
		return &ValidationError{Name: "system_id", err: errors.New(`ent: missing required field "System.system_id"`)}
	}
	if _, ok := sc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "System.public_key"`)}
	}
	if _, ok := sc.mutation.Approved(); !ok {
		return &ValidationError{Name: "approved", err: errors.New(`ent: missing required field "System.approved"`)}
	}
	if _, ok := sc.mutation.LastLogin(); !ok {
		return &ValidationError{Name: "last_login", err: errors.New(`ent: missing required field "System.last_login"`)}
	}
	return nil
}

func (sc *SystemCreate) sqlSave(ctx context.Context) (*System, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SystemCreate) createSpec() (*System, *sqlgraph.CreateSpec) {
	var (
		_node = &System{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(system.Table, sqlgraph.NewFieldSpec(system.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.SystemID(); ok {
		_spec.SetField(system.FieldSystemID, field.TypeUUID, value)
		_node.SystemID = value
	}
	if value, ok := sc.mutation.PublicKey(); ok {
		_spec.SetField(system.FieldPublicKey, field.TypeBytes, value)
		_node.PublicKey = value
	}
	if value, ok := sc.mutation.Approved(); ok {
		_spec.SetField(system.FieldApproved, field.TypeBool, value)
		_node.Approved = value
	}
	if value, ok := sc.mutation.LastLogin(); ok {
		_spec.SetField(system.FieldLastLogin, field.TypeInt64, value)
		_node.LastLogin = value
	}
	return _node, _spec
}

// SystemCreateBulk is the builder for creating many System entities in bulk.
type SystemCreateBulk struct {
	config
	err      error
	builders []*SystemCreate
}

// Save creates the System entities in the database.
func (scb *SystemCreateBulk) Save(ctx context.Context) ([]*System, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*System, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SystemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
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
func (scb *SystemCreateBulk) SaveX(ctx context.Context) []*System {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SystemCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SystemCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
