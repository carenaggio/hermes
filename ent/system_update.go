// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carenaggio/hermes/ent/predicate"
	"github.com/carenaggio/hermes/ent/system"
	"github.com/google/uuid"
)

// SystemUpdate is the builder for updating System entities.
type SystemUpdate struct {
	config
	hooks    []Hook
	mutation *SystemMutation
}

// Where appends a list predicates to the SystemUpdate builder.
func (su *SystemUpdate) Where(ps ...predicate.System) *SystemUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSystemID sets the "system_id" field.
func (su *SystemUpdate) SetSystemID(u uuid.UUID) *SystemUpdate {
	su.mutation.SetSystemID(u)
	return su
}

// SetPublicKey sets the "public_key" field.
func (su *SystemUpdate) SetPublicKey(b []byte) *SystemUpdate {
	su.mutation.SetPublicKey(b)
	return su
}

// SetApproved sets the "approved" field.
func (su *SystemUpdate) SetApproved(b bool) *SystemUpdate {
	su.mutation.SetApproved(b)
	return su
}

// SetLastLogin sets the "last_login" field.
func (su *SystemUpdate) SetLastLogin(i int64) *SystemUpdate {
	su.mutation.ResetLastLogin()
	su.mutation.SetLastLogin(i)
	return su
}

// AddLastLogin adds i to the "last_login" field.
func (su *SystemUpdate) AddLastLogin(i int64) *SystemUpdate {
	su.mutation.AddLastLogin(i)
	return su
}

// Mutation returns the SystemMutation object of the builder.
func (su *SystemUpdate) Mutation() *SystemMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SystemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SystemUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SystemUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SystemUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SystemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(system.Table, system.Columns, sqlgraph.NewFieldSpec(system.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SystemID(); ok {
		_spec.SetField(system.FieldSystemID, field.TypeUUID, value)
	}
	if value, ok := su.mutation.PublicKey(); ok {
		_spec.SetField(system.FieldPublicKey, field.TypeBytes, value)
	}
	if value, ok := su.mutation.Approved(); ok {
		_spec.SetField(system.FieldApproved, field.TypeBool, value)
	}
	if value, ok := su.mutation.LastLogin(); ok {
		_spec.SetField(system.FieldLastLogin, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedLastLogin(); ok {
		_spec.AddField(system.FieldLastLogin, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SystemUpdateOne is the builder for updating a single System entity.
type SystemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SystemMutation
}

// SetSystemID sets the "system_id" field.
func (suo *SystemUpdateOne) SetSystemID(u uuid.UUID) *SystemUpdateOne {
	suo.mutation.SetSystemID(u)
	return suo
}

// SetPublicKey sets the "public_key" field.
func (suo *SystemUpdateOne) SetPublicKey(b []byte) *SystemUpdateOne {
	suo.mutation.SetPublicKey(b)
	return suo
}

// SetApproved sets the "approved" field.
func (suo *SystemUpdateOne) SetApproved(b bool) *SystemUpdateOne {
	suo.mutation.SetApproved(b)
	return suo
}

// SetLastLogin sets the "last_login" field.
func (suo *SystemUpdateOne) SetLastLogin(i int64) *SystemUpdateOne {
	suo.mutation.ResetLastLogin()
	suo.mutation.SetLastLogin(i)
	return suo
}

// AddLastLogin adds i to the "last_login" field.
func (suo *SystemUpdateOne) AddLastLogin(i int64) *SystemUpdateOne {
	suo.mutation.AddLastLogin(i)
	return suo
}

// Mutation returns the SystemMutation object of the builder.
func (suo *SystemUpdateOne) Mutation() *SystemMutation {
	return suo.mutation
}

// Where appends a list predicates to the SystemUpdate builder.
func (suo *SystemUpdateOne) Where(ps ...predicate.System) *SystemUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SystemUpdateOne) Select(field string, fields ...string) *SystemUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated System entity.
func (suo *SystemUpdateOne) Save(ctx context.Context) (*System, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SystemUpdateOne) SaveX(ctx context.Context) *System {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SystemUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SystemUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SystemUpdateOne) sqlSave(ctx context.Context) (_node *System, err error) {
	_spec := sqlgraph.NewUpdateSpec(system.Table, system.Columns, sqlgraph.NewFieldSpec(system.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "System.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, system.FieldID)
		for _, f := range fields {
			if !system.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != system.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.SystemID(); ok {
		_spec.SetField(system.FieldSystemID, field.TypeUUID, value)
	}
	if value, ok := suo.mutation.PublicKey(); ok {
		_spec.SetField(system.FieldPublicKey, field.TypeBytes, value)
	}
	if value, ok := suo.mutation.Approved(); ok {
		_spec.SetField(system.FieldApproved, field.TypeBool, value)
	}
	if value, ok := suo.mutation.LastLogin(); ok {
		_spec.SetField(system.FieldLastLogin, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedLastLogin(); ok {
		_spec.AddField(system.FieldLastLogin, field.TypeInt64, value)
	}
	_node = &System{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
