// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/predicate"
	"api/ent/session"
	"api/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetToken sets the "token" field.
func (su *SessionUpdate) SetToken(s string) *SessionUpdate {
	su.mutation.SetToken(s)
	return su
}

// SetIP sets the "ip" field.
func (su *SessionUpdate) SetIP(s string) *SessionUpdate {
	su.mutation.SetIP(s)
	return su
}

// SetDevice sets the "device" field.
func (su *SessionUpdate) SetDevice(s string) *SessionUpdate {
	su.mutation.SetDevice(s)
	return su
}

// SetUsed sets the "used" field.
func (su *SessionUpdate) SetUsed(i int) *SessionUpdate {
	su.mutation.ResetUsed()
	su.mutation.SetUsed(i)
	return su
}

// AddUsed adds i to the "used" field.
func (su *SessionUpdate) AddUsed(i int) *SessionUpdate {
	su.mutation.AddUsed(i)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUpdatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SessionUpdate) SetUserID(id int) *SessionUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(id *int) *SessionUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if value, ok := su.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if value, ok := su.mutation.Device(); ok {
		_spec.SetField(session.FieldDevice, field.TypeString, value)
	}
	if value, ok := su.mutation.Used(); ok {
		_spec.SetField(session.FieldUsed, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedUsed(); ok {
		_spec.AddField(session.FieldUsed, field.TypeInt, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetToken sets the "token" field.
func (suo *SessionUpdateOne) SetToken(s string) *SessionUpdateOne {
	suo.mutation.SetToken(s)
	return suo
}

// SetIP sets the "ip" field.
func (suo *SessionUpdateOne) SetIP(s string) *SessionUpdateOne {
	suo.mutation.SetIP(s)
	return suo
}

// SetDevice sets the "device" field.
func (suo *SessionUpdateOne) SetDevice(s string) *SessionUpdateOne {
	suo.mutation.SetDevice(s)
	return suo
}

// SetUsed sets the "used" field.
func (suo *SessionUpdateOne) SetUsed(i int) *SessionUpdateOne {
	suo.mutation.ResetUsed()
	suo.mutation.SetUsed(i)
	return suo
}

// AddUsed adds i to the "used" field.
func (suo *SessionUpdateOne) AddUsed(i int) *SessionUpdateOne {
	suo.mutation.AddUsed(i)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUpdatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SessionUpdateOne) SetUserID(id int) *SessionUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(id *int) *SessionUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
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
	if value, ok := suo.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if value, ok := suo.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if value, ok := suo.mutation.Device(); ok {
		_spec.SetField(session.FieldDevice, field.TypeString, value)
	}
	if value, ok := suo.mutation.Used(); ok {
		_spec.SetField(session.FieldUsed, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedUsed(); ok {
		_spec.AddField(session.FieldUsed, field.TypeInt, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
