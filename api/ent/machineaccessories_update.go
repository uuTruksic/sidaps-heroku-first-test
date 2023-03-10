// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineAccessoriesUpdate is the builder for updating MachineAccessories entities.
type MachineAccessoriesUpdate struct {
	config
	hooks    []Hook
	mutation *MachineAccessoriesMutation
}

// Where appends a list predicates to the MachineAccessoriesUpdate builder.
func (mau *MachineAccessoriesUpdate) Where(ps ...predicate.MachineAccessories) *MachineAccessoriesUpdate {
	mau.mutation.Where(ps...)
	return mau
}

// SetName1 sets the "name1" field.
func (mau *MachineAccessoriesUpdate) SetName1(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName1(s)
	return mau
}

// SetNillableName1 sets the "name1" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName1(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName1(*s)
	}
	return mau
}

// ClearName1 clears the value of the "name1" field.
func (mau *MachineAccessoriesUpdate) ClearName1() *MachineAccessoriesUpdate {
	mau.mutation.ClearName1()
	return mau
}

// SetName2 sets the "name2" field.
func (mau *MachineAccessoriesUpdate) SetName2(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName2(s)
	return mau
}

// SetNillableName2 sets the "name2" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName2(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName2(*s)
	}
	return mau
}

// ClearName2 clears the value of the "name2" field.
func (mau *MachineAccessoriesUpdate) ClearName2() *MachineAccessoriesUpdate {
	mau.mutation.ClearName2()
	return mau
}

// SetName3 sets the "name3" field.
func (mau *MachineAccessoriesUpdate) SetName3(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName3(s)
	return mau
}

// SetNillableName3 sets the "name3" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName3(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName3(*s)
	}
	return mau
}

// ClearName3 clears the value of the "name3" field.
func (mau *MachineAccessoriesUpdate) ClearName3() *MachineAccessoriesUpdate {
	mau.mutation.ClearName3()
	return mau
}

// SetName4 sets the "name4" field.
func (mau *MachineAccessoriesUpdate) SetName4(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName4(s)
	return mau
}

// SetNillableName4 sets the "name4" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName4(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName4(*s)
	}
	return mau
}

// ClearName4 clears the value of the "name4" field.
func (mau *MachineAccessoriesUpdate) ClearName4() *MachineAccessoriesUpdate {
	mau.mutation.ClearName4()
	return mau
}

// SetName5 sets the "name5" field.
func (mau *MachineAccessoriesUpdate) SetName5(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName5(s)
	return mau
}

// SetNillableName5 sets the "name5" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName5(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName5(*s)
	}
	return mau
}

// ClearName5 clears the value of the "name5" field.
func (mau *MachineAccessoriesUpdate) ClearName5() *MachineAccessoriesUpdate {
	mau.mutation.ClearName5()
	return mau
}

// SetName6 sets the "name6" field.
func (mau *MachineAccessoriesUpdate) SetName6(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName6(s)
	return mau
}

// SetNillableName6 sets the "name6" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName6(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName6(*s)
	}
	return mau
}

// ClearName6 clears the value of the "name6" field.
func (mau *MachineAccessoriesUpdate) ClearName6() *MachineAccessoriesUpdate {
	mau.mutation.ClearName6()
	return mau
}

// SetName7 sets the "name7" field.
func (mau *MachineAccessoriesUpdate) SetName7(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName7(s)
	return mau
}

// SetNillableName7 sets the "name7" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName7(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName7(*s)
	}
	return mau
}

// ClearName7 clears the value of the "name7" field.
func (mau *MachineAccessoriesUpdate) ClearName7() *MachineAccessoriesUpdate {
	mau.mutation.ClearName7()
	return mau
}

// SetName8 sets the "name8" field.
func (mau *MachineAccessoriesUpdate) SetName8(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName8(s)
	return mau
}

// SetNillableName8 sets the "name8" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName8(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName8(*s)
	}
	return mau
}

// ClearName8 clears the value of the "name8" field.
func (mau *MachineAccessoriesUpdate) ClearName8() *MachineAccessoriesUpdate {
	mau.mutation.ClearName8()
	return mau
}

// SetName9 sets the "name9" field.
func (mau *MachineAccessoriesUpdate) SetName9(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName9(s)
	return mau
}

// SetNillableName9 sets the "name9" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName9(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName9(*s)
	}
	return mau
}

// ClearName9 clears the value of the "name9" field.
func (mau *MachineAccessoriesUpdate) ClearName9() *MachineAccessoriesUpdate {
	mau.mutation.ClearName9()
	return mau
}

// SetName10 sets the "name10" field.
func (mau *MachineAccessoriesUpdate) SetName10(s string) *MachineAccessoriesUpdate {
	mau.mutation.SetName10(s)
	return mau
}

// SetNillableName10 sets the "name10" field if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableName10(s *string) *MachineAccessoriesUpdate {
	if s != nil {
		mau.SetName10(*s)
	}
	return mau
}

// ClearName10 clears the value of the "name10" field.
func (mau *MachineAccessoriesUpdate) ClearName10() *MachineAccessoriesUpdate {
	mau.mutation.ClearName10()
	return mau
}

// SetMachineID sets the "machine" edge to the Machine entity by ID.
func (mau *MachineAccessoriesUpdate) SetMachineID(id int) *MachineAccessoriesUpdate {
	mau.mutation.SetMachineID(id)
	return mau
}

// SetNillableMachineID sets the "machine" edge to the Machine entity by ID if the given value is not nil.
func (mau *MachineAccessoriesUpdate) SetNillableMachineID(id *int) *MachineAccessoriesUpdate {
	if id != nil {
		mau = mau.SetMachineID(*id)
	}
	return mau
}

// SetMachine sets the "machine" edge to the Machine entity.
func (mau *MachineAccessoriesUpdate) SetMachine(m *Machine) *MachineAccessoriesUpdate {
	return mau.SetMachineID(m.ID)
}

// Mutation returns the MachineAccessoriesMutation object of the builder.
func (mau *MachineAccessoriesUpdate) Mutation() *MachineAccessoriesMutation {
	return mau.mutation
}

// ClearMachine clears the "machine" edge to the Machine entity.
func (mau *MachineAccessoriesUpdate) ClearMachine() *MachineAccessoriesUpdate {
	mau.mutation.ClearMachine()
	return mau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mau *MachineAccessoriesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mau.hooks) == 0 {
		affected, err = mau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineAccessoriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mau.mutation = mutation
			affected, err = mau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mau.hooks) - 1; i >= 0; i-- {
			if mau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mau *MachineAccessoriesUpdate) SaveX(ctx context.Context) int {
	affected, err := mau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mau *MachineAccessoriesUpdate) Exec(ctx context.Context) error {
	_, err := mau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mau *MachineAccessoriesUpdate) ExecX(ctx context.Context) {
	if err := mau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mau *MachineAccessoriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machineaccessories.Table,
			Columns: machineaccessories.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machineaccessories.FieldID,
			},
		},
	}
	if ps := mau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mau.mutation.Name1(); ok {
		_spec.SetField(machineaccessories.FieldName1, field.TypeString, value)
	}
	if mau.mutation.Name1Cleared() {
		_spec.ClearField(machineaccessories.FieldName1, field.TypeString)
	}
	if value, ok := mau.mutation.Name2(); ok {
		_spec.SetField(machineaccessories.FieldName2, field.TypeString, value)
	}
	if mau.mutation.Name2Cleared() {
		_spec.ClearField(machineaccessories.FieldName2, field.TypeString)
	}
	if value, ok := mau.mutation.Name3(); ok {
		_spec.SetField(machineaccessories.FieldName3, field.TypeString, value)
	}
	if mau.mutation.Name3Cleared() {
		_spec.ClearField(machineaccessories.FieldName3, field.TypeString)
	}
	if value, ok := mau.mutation.Name4(); ok {
		_spec.SetField(machineaccessories.FieldName4, field.TypeString, value)
	}
	if mau.mutation.Name4Cleared() {
		_spec.ClearField(machineaccessories.FieldName4, field.TypeString)
	}
	if value, ok := mau.mutation.Name5(); ok {
		_spec.SetField(machineaccessories.FieldName5, field.TypeString, value)
	}
	if mau.mutation.Name5Cleared() {
		_spec.ClearField(machineaccessories.FieldName5, field.TypeString)
	}
	if value, ok := mau.mutation.Name6(); ok {
		_spec.SetField(machineaccessories.FieldName6, field.TypeString, value)
	}
	if mau.mutation.Name6Cleared() {
		_spec.ClearField(machineaccessories.FieldName6, field.TypeString)
	}
	if value, ok := mau.mutation.Name7(); ok {
		_spec.SetField(machineaccessories.FieldName7, field.TypeString, value)
	}
	if mau.mutation.Name7Cleared() {
		_spec.ClearField(machineaccessories.FieldName7, field.TypeString)
	}
	if value, ok := mau.mutation.Name8(); ok {
		_spec.SetField(machineaccessories.FieldName8, field.TypeString, value)
	}
	if mau.mutation.Name8Cleared() {
		_spec.ClearField(machineaccessories.FieldName8, field.TypeString)
	}
	if value, ok := mau.mutation.Name9(); ok {
		_spec.SetField(machineaccessories.FieldName9, field.TypeString, value)
	}
	if mau.mutation.Name9Cleared() {
		_spec.ClearField(machineaccessories.FieldName9, field.TypeString)
	}
	if value, ok := mau.mutation.Name10(); ok {
		_spec.SetField(machineaccessories.FieldName10, field.TypeString, value)
	}
	if mau.mutation.Name10Cleared() {
		_spec.ClearField(machineaccessories.FieldName10, field.TypeString)
	}
	if mau.mutation.MachineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   machineaccessories.MachineTable,
			Columns: []string{machineaccessories.MachineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mau.mutation.MachineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   machineaccessories.MachineTable,
			Columns: []string{machineaccessories.MachineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machineaccessories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MachineAccessoriesUpdateOne is the builder for updating a single MachineAccessories entity.
type MachineAccessoriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MachineAccessoriesMutation
}

// SetName1 sets the "name1" field.
func (mauo *MachineAccessoriesUpdateOne) SetName1(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName1(s)
	return mauo
}

// SetNillableName1 sets the "name1" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName1(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName1(*s)
	}
	return mauo
}

// ClearName1 clears the value of the "name1" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName1() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName1()
	return mauo
}

// SetName2 sets the "name2" field.
func (mauo *MachineAccessoriesUpdateOne) SetName2(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName2(s)
	return mauo
}

// SetNillableName2 sets the "name2" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName2(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName2(*s)
	}
	return mauo
}

// ClearName2 clears the value of the "name2" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName2() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName2()
	return mauo
}

// SetName3 sets the "name3" field.
func (mauo *MachineAccessoriesUpdateOne) SetName3(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName3(s)
	return mauo
}

// SetNillableName3 sets the "name3" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName3(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName3(*s)
	}
	return mauo
}

// ClearName3 clears the value of the "name3" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName3() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName3()
	return mauo
}

// SetName4 sets the "name4" field.
func (mauo *MachineAccessoriesUpdateOne) SetName4(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName4(s)
	return mauo
}

// SetNillableName4 sets the "name4" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName4(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName4(*s)
	}
	return mauo
}

// ClearName4 clears the value of the "name4" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName4() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName4()
	return mauo
}

// SetName5 sets the "name5" field.
func (mauo *MachineAccessoriesUpdateOne) SetName5(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName5(s)
	return mauo
}

// SetNillableName5 sets the "name5" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName5(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName5(*s)
	}
	return mauo
}

// ClearName5 clears the value of the "name5" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName5() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName5()
	return mauo
}

// SetName6 sets the "name6" field.
func (mauo *MachineAccessoriesUpdateOne) SetName6(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName6(s)
	return mauo
}

// SetNillableName6 sets the "name6" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName6(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName6(*s)
	}
	return mauo
}

// ClearName6 clears the value of the "name6" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName6() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName6()
	return mauo
}

// SetName7 sets the "name7" field.
func (mauo *MachineAccessoriesUpdateOne) SetName7(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName7(s)
	return mauo
}

// SetNillableName7 sets the "name7" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName7(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName7(*s)
	}
	return mauo
}

// ClearName7 clears the value of the "name7" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName7() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName7()
	return mauo
}

// SetName8 sets the "name8" field.
func (mauo *MachineAccessoriesUpdateOne) SetName8(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName8(s)
	return mauo
}

// SetNillableName8 sets the "name8" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName8(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName8(*s)
	}
	return mauo
}

// ClearName8 clears the value of the "name8" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName8() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName8()
	return mauo
}

// SetName9 sets the "name9" field.
func (mauo *MachineAccessoriesUpdateOne) SetName9(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName9(s)
	return mauo
}

// SetNillableName9 sets the "name9" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName9(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName9(*s)
	}
	return mauo
}

// ClearName9 clears the value of the "name9" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName9() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName9()
	return mauo
}

// SetName10 sets the "name10" field.
func (mauo *MachineAccessoriesUpdateOne) SetName10(s string) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetName10(s)
	return mauo
}

// SetNillableName10 sets the "name10" field if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableName10(s *string) *MachineAccessoriesUpdateOne {
	if s != nil {
		mauo.SetName10(*s)
	}
	return mauo
}

// ClearName10 clears the value of the "name10" field.
func (mauo *MachineAccessoriesUpdateOne) ClearName10() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearName10()
	return mauo
}

// SetMachineID sets the "machine" edge to the Machine entity by ID.
func (mauo *MachineAccessoriesUpdateOne) SetMachineID(id int) *MachineAccessoriesUpdateOne {
	mauo.mutation.SetMachineID(id)
	return mauo
}

// SetNillableMachineID sets the "machine" edge to the Machine entity by ID if the given value is not nil.
func (mauo *MachineAccessoriesUpdateOne) SetNillableMachineID(id *int) *MachineAccessoriesUpdateOne {
	if id != nil {
		mauo = mauo.SetMachineID(*id)
	}
	return mauo
}

// SetMachine sets the "machine" edge to the Machine entity.
func (mauo *MachineAccessoriesUpdateOne) SetMachine(m *Machine) *MachineAccessoriesUpdateOne {
	return mauo.SetMachineID(m.ID)
}

// Mutation returns the MachineAccessoriesMutation object of the builder.
func (mauo *MachineAccessoriesUpdateOne) Mutation() *MachineAccessoriesMutation {
	return mauo.mutation
}

// ClearMachine clears the "machine" edge to the Machine entity.
func (mauo *MachineAccessoriesUpdateOne) ClearMachine() *MachineAccessoriesUpdateOne {
	mauo.mutation.ClearMachine()
	return mauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mauo *MachineAccessoriesUpdateOne) Select(field string, fields ...string) *MachineAccessoriesUpdateOne {
	mauo.fields = append([]string{field}, fields...)
	return mauo
}

// Save executes the query and returns the updated MachineAccessories entity.
func (mauo *MachineAccessoriesUpdateOne) Save(ctx context.Context) (*MachineAccessories, error) {
	var (
		err  error
		node *MachineAccessories
	)
	if len(mauo.hooks) == 0 {
		node, err = mauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineAccessoriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mauo.mutation = mutation
			node, err = mauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mauo.hooks) - 1; i >= 0; i-- {
			if mauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MachineAccessories)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MachineAccessoriesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mauo *MachineAccessoriesUpdateOne) SaveX(ctx context.Context) *MachineAccessories {
	node, err := mauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mauo *MachineAccessoriesUpdateOne) Exec(ctx context.Context) error {
	_, err := mauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mauo *MachineAccessoriesUpdateOne) ExecX(ctx context.Context) {
	if err := mauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mauo *MachineAccessoriesUpdateOne) sqlSave(ctx context.Context) (_node *MachineAccessories, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machineaccessories.Table,
			Columns: machineaccessories.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machineaccessories.FieldID,
			},
		},
	}
	id, ok := mauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MachineAccessories.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, machineaccessories.FieldID)
		for _, f := range fields {
			if !machineaccessories.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != machineaccessories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mauo.mutation.Name1(); ok {
		_spec.SetField(machineaccessories.FieldName1, field.TypeString, value)
	}
	if mauo.mutation.Name1Cleared() {
		_spec.ClearField(machineaccessories.FieldName1, field.TypeString)
	}
	if value, ok := mauo.mutation.Name2(); ok {
		_spec.SetField(machineaccessories.FieldName2, field.TypeString, value)
	}
	if mauo.mutation.Name2Cleared() {
		_spec.ClearField(machineaccessories.FieldName2, field.TypeString)
	}
	if value, ok := mauo.mutation.Name3(); ok {
		_spec.SetField(machineaccessories.FieldName3, field.TypeString, value)
	}
	if mauo.mutation.Name3Cleared() {
		_spec.ClearField(machineaccessories.FieldName3, field.TypeString)
	}
	if value, ok := mauo.mutation.Name4(); ok {
		_spec.SetField(machineaccessories.FieldName4, field.TypeString, value)
	}
	if mauo.mutation.Name4Cleared() {
		_spec.ClearField(machineaccessories.FieldName4, field.TypeString)
	}
	if value, ok := mauo.mutation.Name5(); ok {
		_spec.SetField(machineaccessories.FieldName5, field.TypeString, value)
	}
	if mauo.mutation.Name5Cleared() {
		_spec.ClearField(machineaccessories.FieldName5, field.TypeString)
	}
	if value, ok := mauo.mutation.Name6(); ok {
		_spec.SetField(machineaccessories.FieldName6, field.TypeString, value)
	}
	if mauo.mutation.Name6Cleared() {
		_spec.ClearField(machineaccessories.FieldName6, field.TypeString)
	}
	if value, ok := mauo.mutation.Name7(); ok {
		_spec.SetField(machineaccessories.FieldName7, field.TypeString, value)
	}
	if mauo.mutation.Name7Cleared() {
		_spec.ClearField(machineaccessories.FieldName7, field.TypeString)
	}
	if value, ok := mauo.mutation.Name8(); ok {
		_spec.SetField(machineaccessories.FieldName8, field.TypeString, value)
	}
	if mauo.mutation.Name8Cleared() {
		_spec.ClearField(machineaccessories.FieldName8, field.TypeString)
	}
	if value, ok := mauo.mutation.Name9(); ok {
		_spec.SetField(machineaccessories.FieldName9, field.TypeString, value)
	}
	if mauo.mutation.Name9Cleared() {
		_spec.ClearField(machineaccessories.FieldName9, field.TypeString)
	}
	if value, ok := mauo.mutation.Name10(); ok {
		_spec.SetField(machineaccessories.FieldName10, field.TypeString, value)
	}
	if mauo.mutation.Name10Cleared() {
		_spec.ClearField(machineaccessories.FieldName10, field.TypeString)
	}
	if mauo.mutation.MachineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   machineaccessories.MachineTable,
			Columns: []string{machineaccessories.MachineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mauo.mutation.MachineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   machineaccessories.MachineTable,
			Columns: []string{machineaccessories.MachineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MachineAccessories{config: mauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machineaccessories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
