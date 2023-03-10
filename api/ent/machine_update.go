// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/machinespecification"
	"api/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineUpdate is the builder for updating Machine entities.
type MachineUpdate struct {
	config
	hooks    []Hook
	mutation *MachineMutation
}

// Where appends a list predicates to the MachineUpdate builder.
func (mu *MachineUpdate) Where(ps ...predicate.Machine) *MachineUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MachineUpdate) SetName(s string) *MachineUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetYear sets the "year" field.
func (mu *MachineUpdate) SetYear(u uint) *MachineUpdate {
	mu.mutation.ResetYear()
	mu.mutation.SetYear(u)
	return mu
}

// AddYear adds u to the "year" field.
func (mu *MachineUpdate) AddYear(u int) *MachineUpdate {
	mu.mutation.AddYear(u)
	return mu
}

// SetLiftHeight sets the "liftHeight" field.
func (mu *MachineUpdate) SetLiftHeight(u uint) *MachineUpdate {
	mu.mutation.ResetLiftHeight()
	mu.mutation.SetLiftHeight(u)
	return mu
}

// AddLiftHeight adds u to the "liftHeight" field.
func (mu *MachineUpdate) AddLiftHeight(u int) *MachineUpdate {
	mu.mutation.AddLiftHeight(u)
	return mu
}

// SetMotorHours sets the "motorHours" field.
func (mu *MachineUpdate) SetMotorHours(u uint) *MachineUpdate {
	mu.mutation.ResetMotorHours()
	mu.mutation.SetMotorHours(u)
	return mu
}

// AddMotorHours adds u to the "motorHours" field.
func (mu *MachineUpdate) AddMotorHours(u int) *MachineUpdate {
	mu.mutation.AddMotorHours(u)
	return mu
}

// SetDrive sets the "drive" field.
func (mu *MachineUpdate) SetDrive(s string) *MachineUpdate {
	mu.mutation.SetDrive(s)
	return mu
}

// SetWeight sets the "weight" field.
func (mu *MachineUpdate) SetWeight(u uint) *MachineUpdate {
	mu.mutation.ResetWeight()
	mu.mutation.SetWeight(u)
	return mu
}

// AddWeight adds u to the "weight" field.
func (mu *MachineUpdate) AddWeight(u int) *MachineUpdate {
	mu.mutation.AddWeight(u)
	return mu
}

// SetLoadCapacity sets the "loadCapacity" field.
func (mu *MachineUpdate) SetLoadCapacity(u uint) *MachineUpdate {
	mu.mutation.ResetLoadCapacity()
	mu.mutation.SetLoadCapacity(u)
	return mu
}

// AddLoadCapacity adds u to the "loadCapacity" field.
func (mu *MachineUpdate) AddLoadCapacity(u int) *MachineUpdate {
	mu.mutation.AddLoadCapacity(u)
	return mu
}

// SetRentPrice sets the "rentPrice" field.
func (mu *MachineUpdate) SetRentPrice(u uint) *MachineUpdate {
	mu.mutation.ResetRentPrice()
	mu.mutation.SetRentPrice(u)
	return mu
}

// SetNillableRentPrice sets the "rentPrice" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableRentPrice(u *uint) *MachineUpdate {
	if u != nil {
		mu.SetRentPrice(*u)
	}
	return mu
}

// AddRentPrice adds u to the "rentPrice" field.
func (mu *MachineUpdate) AddRentPrice(u int) *MachineUpdate {
	mu.mutation.AddRentPrice(u)
	return mu
}

// ClearRentPrice clears the value of the "rentPrice" field.
func (mu *MachineUpdate) ClearRentPrice() *MachineUpdate {
	mu.mutation.ClearRentPrice()
	return mu
}

// SetSalePrice sets the "salePrice" field.
func (mu *MachineUpdate) SetSalePrice(u uint) *MachineUpdate {
	mu.mutation.ResetSalePrice()
	mu.mutation.SetSalePrice(u)
	return mu
}

// AddSalePrice adds u to the "salePrice" field.
func (mu *MachineUpdate) AddSalePrice(u int) *MachineUpdate {
	mu.mutation.AddSalePrice(u)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MachineUpdate) SetDescription(s string) *MachineUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetHidden sets the "hidden" field.
func (mu *MachineUpdate) SetHidden(b bool) *MachineUpdate {
	mu.mutation.SetHidden(b)
	return mu
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableHidden(b *bool) *MachineUpdate {
	if b != nil {
		mu.SetHidden(*b)
	}
	return mu
}

// SetDeleted sets the "deleted" field.
func (mu *MachineUpdate) SetDeleted(b bool) *MachineUpdate {
	mu.mutation.SetDeleted(b)
	return mu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableDeleted(b *bool) *MachineUpdate {
	if b != nil {
		mu.SetDeleted(*b)
	}
	return mu
}

// SetNumberOfImages sets the "numberOfImages" field.
func (mu *MachineUpdate) SetNumberOfImages(i int) *MachineUpdate {
	mu.mutation.ResetNumberOfImages()
	mu.mutation.SetNumberOfImages(i)
	return mu
}

// SetNillableNumberOfImages sets the "numberOfImages" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableNumberOfImages(i *int) *MachineUpdate {
	if i != nil {
		mu.SetNumberOfImages(*i)
	}
	return mu
}

// AddNumberOfImages adds i to the "numberOfImages" field.
func (mu *MachineUpdate) AddNumberOfImages(i int) *MachineUpdate {
	mu.mutation.AddNumberOfImages(i)
	return mu
}

// ClearNumberOfImages clears the value of the "numberOfImages" field.
func (mu *MachineUpdate) ClearNumberOfImages() *MachineUpdate {
	mu.mutation.ClearNumberOfImages()
	return mu
}

// SetInternalDescription sets the "internalDescription" field.
func (mu *MachineUpdate) SetInternalDescription(s string) *MachineUpdate {
	mu.mutation.SetInternalDescription(s)
	return mu
}

// SetNillableInternalDescription sets the "internalDescription" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableInternalDescription(s *string) *MachineUpdate {
	if s != nil {
		mu.SetInternalDescription(*s)
	}
	return mu
}

// ClearInternalDescription clears the value of the "internalDescription" field.
func (mu *MachineUpdate) ClearInternalDescription() *MachineUpdate {
	mu.mutation.ClearInternalDescription()
	return mu
}

// AddMachineSpecificationIDs adds the "machineSpecification" edge to the MachineSpecification entity by IDs.
func (mu *MachineUpdate) AddMachineSpecificationIDs(ids ...int) *MachineUpdate {
	mu.mutation.AddMachineSpecificationIDs(ids...)
	return mu
}

// AddMachineSpecification adds the "machineSpecification" edges to the MachineSpecification entity.
func (mu *MachineUpdate) AddMachineSpecification(m ...*MachineSpecification) *MachineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMachineSpecificationIDs(ids...)
}

// AddMachineAccessoryIDs adds the "machineAccessories" edge to the MachineAccessories entity by IDs.
func (mu *MachineUpdate) AddMachineAccessoryIDs(ids ...int) *MachineUpdate {
	mu.mutation.AddMachineAccessoryIDs(ids...)
	return mu
}

// AddMachineAccessories adds the "machineAccessories" edges to the MachineAccessories entity.
func (mu *MachineUpdate) AddMachineAccessories(m ...*MachineAccessories) *MachineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMachineAccessoryIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (mu *MachineUpdate) Mutation() *MachineMutation {
	return mu.mutation
}

// ClearMachineSpecification clears all "machineSpecification" edges to the MachineSpecification entity.
func (mu *MachineUpdate) ClearMachineSpecification() *MachineUpdate {
	mu.mutation.ClearMachineSpecification()
	return mu
}

// RemoveMachineSpecificationIDs removes the "machineSpecification" edge to MachineSpecification entities by IDs.
func (mu *MachineUpdate) RemoveMachineSpecificationIDs(ids ...int) *MachineUpdate {
	mu.mutation.RemoveMachineSpecificationIDs(ids...)
	return mu
}

// RemoveMachineSpecification removes "machineSpecification" edges to MachineSpecification entities.
func (mu *MachineUpdate) RemoveMachineSpecification(m ...*MachineSpecification) *MachineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMachineSpecificationIDs(ids...)
}

// ClearMachineAccessories clears all "machineAccessories" edges to the MachineAccessories entity.
func (mu *MachineUpdate) ClearMachineAccessories() *MachineUpdate {
	mu.mutation.ClearMachineAccessories()
	return mu
}

// RemoveMachineAccessoryIDs removes the "machineAccessories" edge to MachineAccessories entities by IDs.
func (mu *MachineUpdate) RemoveMachineAccessoryIDs(ids ...int) *MachineUpdate {
	mu.mutation.RemoveMachineAccessoryIDs(ids...)
	return mu
}

// RemoveMachineAccessories removes "machineAccessories" edges to MachineAccessories entities.
func (mu *MachineUpdate) RemoveMachineAccessories(m ...*MachineAccessories) *MachineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMachineAccessoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MachineUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MachineUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MachineUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MachineUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MachineUpdate) check() error {
	if v, ok := mu.mutation.Description(); ok {
		if err := machine.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Machine.description": %w`, err)}
		}
	}
	if v, ok := mu.mutation.InternalDescription(); ok {
		if err := machine.InternalDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "internalDescription", err: fmt.Errorf(`ent: validator failed for field "Machine.internalDescription": %w`, err)}
		}
	}
	return nil
}

func (mu *MachineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machine.Table,
			Columns: machine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(machine.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Year(); ok {
		_spec.SetField(machine.FieldYear, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedYear(); ok {
		_spec.AddField(machine.FieldYear, field.TypeUint, value)
	}
	if value, ok := mu.mutation.LiftHeight(); ok {
		_spec.SetField(machine.FieldLiftHeight, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedLiftHeight(); ok {
		_spec.AddField(machine.FieldLiftHeight, field.TypeUint, value)
	}
	if value, ok := mu.mutation.MotorHours(); ok {
		_spec.SetField(machine.FieldMotorHours, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedMotorHours(); ok {
		_spec.AddField(machine.FieldMotorHours, field.TypeUint, value)
	}
	if value, ok := mu.mutation.Drive(); ok {
		_spec.SetField(machine.FieldDrive, field.TypeString, value)
	}
	if value, ok := mu.mutation.Weight(); ok {
		_spec.SetField(machine.FieldWeight, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedWeight(); ok {
		_spec.AddField(machine.FieldWeight, field.TypeUint, value)
	}
	if value, ok := mu.mutation.LoadCapacity(); ok {
		_spec.SetField(machine.FieldLoadCapacity, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedLoadCapacity(); ok {
		_spec.AddField(machine.FieldLoadCapacity, field.TypeUint, value)
	}
	if value, ok := mu.mutation.RentPrice(); ok {
		_spec.SetField(machine.FieldRentPrice, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedRentPrice(); ok {
		_spec.AddField(machine.FieldRentPrice, field.TypeUint, value)
	}
	if mu.mutation.RentPriceCleared() {
		_spec.ClearField(machine.FieldRentPrice, field.TypeUint)
	}
	if value, ok := mu.mutation.SalePrice(); ok {
		_spec.SetField(machine.FieldSalePrice, field.TypeUint, value)
	}
	if value, ok := mu.mutation.AddedSalePrice(); ok {
		_spec.AddField(machine.FieldSalePrice, field.TypeUint, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(machine.FieldDescription, field.TypeString, value)
	}
	if value, ok := mu.mutation.Hidden(); ok {
		_spec.SetField(machine.FieldHidden, field.TypeBool, value)
	}
	if value, ok := mu.mutation.Deleted(); ok {
		_spec.SetField(machine.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := mu.mutation.NumberOfImages(); ok {
		_spec.SetField(machine.FieldNumberOfImages, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedNumberOfImages(); ok {
		_spec.AddField(machine.FieldNumberOfImages, field.TypeInt, value)
	}
	if mu.mutation.NumberOfImagesCleared() {
		_spec.ClearField(machine.FieldNumberOfImages, field.TypeInt)
	}
	if value, ok := mu.mutation.InternalDescription(); ok {
		_spec.SetField(machine.FieldInternalDescription, field.TypeString, value)
	}
	if mu.mutation.InternalDescriptionCleared() {
		_spec.ClearField(machine.FieldInternalDescription, field.TypeString)
	}
	if mu.mutation.MachineSpecificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMachineSpecificationIDs(); len(nodes) > 0 && !mu.mutation.MachineSpecificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MachineSpecificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MachineAccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMachineAccessoriesIDs(); len(nodes) > 0 && !mu.mutation.MachineAccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MachineAccessoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MachineUpdateOne is the builder for updating a single Machine entity.
type MachineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MachineMutation
}

// SetName sets the "name" field.
func (muo *MachineUpdateOne) SetName(s string) *MachineUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetYear sets the "year" field.
func (muo *MachineUpdateOne) SetYear(u uint) *MachineUpdateOne {
	muo.mutation.ResetYear()
	muo.mutation.SetYear(u)
	return muo
}

// AddYear adds u to the "year" field.
func (muo *MachineUpdateOne) AddYear(u int) *MachineUpdateOne {
	muo.mutation.AddYear(u)
	return muo
}

// SetLiftHeight sets the "liftHeight" field.
func (muo *MachineUpdateOne) SetLiftHeight(u uint) *MachineUpdateOne {
	muo.mutation.ResetLiftHeight()
	muo.mutation.SetLiftHeight(u)
	return muo
}

// AddLiftHeight adds u to the "liftHeight" field.
func (muo *MachineUpdateOne) AddLiftHeight(u int) *MachineUpdateOne {
	muo.mutation.AddLiftHeight(u)
	return muo
}

// SetMotorHours sets the "motorHours" field.
func (muo *MachineUpdateOne) SetMotorHours(u uint) *MachineUpdateOne {
	muo.mutation.ResetMotorHours()
	muo.mutation.SetMotorHours(u)
	return muo
}

// AddMotorHours adds u to the "motorHours" field.
func (muo *MachineUpdateOne) AddMotorHours(u int) *MachineUpdateOne {
	muo.mutation.AddMotorHours(u)
	return muo
}

// SetDrive sets the "drive" field.
func (muo *MachineUpdateOne) SetDrive(s string) *MachineUpdateOne {
	muo.mutation.SetDrive(s)
	return muo
}

// SetWeight sets the "weight" field.
func (muo *MachineUpdateOne) SetWeight(u uint) *MachineUpdateOne {
	muo.mutation.ResetWeight()
	muo.mutation.SetWeight(u)
	return muo
}

// AddWeight adds u to the "weight" field.
func (muo *MachineUpdateOne) AddWeight(u int) *MachineUpdateOne {
	muo.mutation.AddWeight(u)
	return muo
}

// SetLoadCapacity sets the "loadCapacity" field.
func (muo *MachineUpdateOne) SetLoadCapacity(u uint) *MachineUpdateOne {
	muo.mutation.ResetLoadCapacity()
	muo.mutation.SetLoadCapacity(u)
	return muo
}

// AddLoadCapacity adds u to the "loadCapacity" field.
func (muo *MachineUpdateOne) AddLoadCapacity(u int) *MachineUpdateOne {
	muo.mutation.AddLoadCapacity(u)
	return muo
}

// SetRentPrice sets the "rentPrice" field.
func (muo *MachineUpdateOne) SetRentPrice(u uint) *MachineUpdateOne {
	muo.mutation.ResetRentPrice()
	muo.mutation.SetRentPrice(u)
	return muo
}

// SetNillableRentPrice sets the "rentPrice" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableRentPrice(u *uint) *MachineUpdateOne {
	if u != nil {
		muo.SetRentPrice(*u)
	}
	return muo
}

// AddRentPrice adds u to the "rentPrice" field.
func (muo *MachineUpdateOne) AddRentPrice(u int) *MachineUpdateOne {
	muo.mutation.AddRentPrice(u)
	return muo
}

// ClearRentPrice clears the value of the "rentPrice" field.
func (muo *MachineUpdateOne) ClearRentPrice() *MachineUpdateOne {
	muo.mutation.ClearRentPrice()
	return muo
}

// SetSalePrice sets the "salePrice" field.
func (muo *MachineUpdateOne) SetSalePrice(u uint) *MachineUpdateOne {
	muo.mutation.ResetSalePrice()
	muo.mutation.SetSalePrice(u)
	return muo
}

// AddSalePrice adds u to the "salePrice" field.
func (muo *MachineUpdateOne) AddSalePrice(u int) *MachineUpdateOne {
	muo.mutation.AddSalePrice(u)
	return muo
}

// SetDescription sets the "description" field.
func (muo *MachineUpdateOne) SetDescription(s string) *MachineUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetHidden sets the "hidden" field.
func (muo *MachineUpdateOne) SetHidden(b bool) *MachineUpdateOne {
	muo.mutation.SetHidden(b)
	return muo
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableHidden(b *bool) *MachineUpdateOne {
	if b != nil {
		muo.SetHidden(*b)
	}
	return muo
}

// SetDeleted sets the "deleted" field.
func (muo *MachineUpdateOne) SetDeleted(b bool) *MachineUpdateOne {
	muo.mutation.SetDeleted(b)
	return muo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableDeleted(b *bool) *MachineUpdateOne {
	if b != nil {
		muo.SetDeleted(*b)
	}
	return muo
}

// SetNumberOfImages sets the "numberOfImages" field.
func (muo *MachineUpdateOne) SetNumberOfImages(i int) *MachineUpdateOne {
	muo.mutation.ResetNumberOfImages()
	muo.mutation.SetNumberOfImages(i)
	return muo
}

// SetNillableNumberOfImages sets the "numberOfImages" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableNumberOfImages(i *int) *MachineUpdateOne {
	if i != nil {
		muo.SetNumberOfImages(*i)
	}
	return muo
}

// AddNumberOfImages adds i to the "numberOfImages" field.
func (muo *MachineUpdateOne) AddNumberOfImages(i int) *MachineUpdateOne {
	muo.mutation.AddNumberOfImages(i)
	return muo
}

// ClearNumberOfImages clears the value of the "numberOfImages" field.
func (muo *MachineUpdateOne) ClearNumberOfImages() *MachineUpdateOne {
	muo.mutation.ClearNumberOfImages()
	return muo
}

// SetInternalDescription sets the "internalDescription" field.
func (muo *MachineUpdateOne) SetInternalDescription(s string) *MachineUpdateOne {
	muo.mutation.SetInternalDescription(s)
	return muo
}

// SetNillableInternalDescription sets the "internalDescription" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableInternalDescription(s *string) *MachineUpdateOne {
	if s != nil {
		muo.SetInternalDescription(*s)
	}
	return muo
}

// ClearInternalDescription clears the value of the "internalDescription" field.
func (muo *MachineUpdateOne) ClearInternalDescription() *MachineUpdateOne {
	muo.mutation.ClearInternalDescription()
	return muo
}

// AddMachineSpecificationIDs adds the "machineSpecification" edge to the MachineSpecification entity by IDs.
func (muo *MachineUpdateOne) AddMachineSpecificationIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.AddMachineSpecificationIDs(ids...)
	return muo
}

// AddMachineSpecification adds the "machineSpecification" edges to the MachineSpecification entity.
func (muo *MachineUpdateOne) AddMachineSpecification(m ...*MachineSpecification) *MachineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMachineSpecificationIDs(ids...)
}

// AddMachineAccessoryIDs adds the "machineAccessories" edge to the MachineAccessories entity by IDs.
func (muo *MachineUpdateOne) AddMachineAccessoryIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.AddMachineAccessoryIDs(ids...)
	return muo
}

// AddMachineAccessories adds the "machineAccessories" edges to the MachineAccessories entity.
func (muo *MachineUpdateOne) AddMachineAccessories(m ...*MachineAccessories) *MachineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMachineAccessoryIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (muo *MachineUpdateOne) Mutation() *MachineMutation {
	return muo.mutation
}

// ClearMachineSpecification clears all "machineSpecification" edges to the MachineSpecification entity.
func (muo *MachineUpdateOne) ClearMachineSpecification() *MachineUpdateOne {
	muo.mutation.ClearMachineSpecification()
	return muo
}

// RemoveMachineSpecificationIDs removes the "machineSpecification" edge to MachineSpecification entities by IDs.
func (muo *MachineUpdateOne) RemoveMachineSpecificationIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.RemoveMachineSpecificationIDs(ids...)
	return muo
}

// RemoveMachineSpecification removes "machineSpecification" edges to MachineSpecification entities.
func (muo *MachineUpdateOne) RemoveMachineSpecification(m ...*MachineSpecification) *MachineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMachineSpecificationIDs(ids...)
}

// ClearMachineAccessories clears all "machineAccessories" edges to the MachineAccessories entity.
func (muo *MachineUpdateOne) ClearMachineAccessories() *MachineUpdateOne {
	muo.mutation.ClearMachineAccessories()
	return muo
}

// RemoveMachineAccessoryIDs removes the "machineAccessories" edge to MachineAccessories entities by IDs.
func (muo *MachineUpdateOne) RemoveMachineAccessoryIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.RemoveMachineAccessoryIDs(ids...)
	return muo
}

// RemoveMachineAccessories removes "machineAccessories" edges to MachineAccessories entities.
func (muo *MachineUpdateOne) RemoveMachineAccessories(m ...*MachineAccessories) *MachineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMachineAccessoryIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MachineUpdateOne) Select(field string, fields ...string) *MachineUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Machine entity.
func (muo *MachineUpdateOne) Save(ctx context.Context) (*Machine, error) {
	var (
		err  error
		node *Machine
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Machine)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MachineMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MachineUpdateOne) SaveX(ctx context.Context) *Machine {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MachineUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MachineUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MachineUpdateOne) check() error {
	if v, ok := muo.mutation.Description(); ok {
		if err := machine.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Machine.description": %w`, err)}
		}
	}
	if v, ok := muo.mutation.InternalDescription(); ok {
		if err := machine.InternalDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "internalDescription", err: fmt.Errorf(`ent: validator failed for field "Machine.internalDescription": %w`, err)}
		}
	}
	return nil
}

func (muo *MachineUpdateOne) sqlSave(ctx context.Context) (_node *Machine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machine.Table,
			Columns: machine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Machine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, machine.FieldID)
		for _, f := range fields {
			if !machine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != machine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(machine.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Year(); ok {
		_spec.SetField(machine.FieldYear, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedYear(); ok {
		_spec.AddField(machine.FieldYear, field.TypeUint, value)
	}
	if value, ok := muo.mutation.LiftHeight(); ok {
		_spec.SetField(machine.FieldLiftHeight, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedLiftHeight(); ok {
		_spec.AddField(machine.FieldLiftHeight, field.TypeUint, value)
	}
	if value, ok := muo.mutation.MotorHours(); ok {
		_spec.SetField(machine.FieldMotorHours, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedMotorHours(); ok {
		_spec.AddField(machine.FieldMotorHours, field.TypeUint, value)
	}
	if value, ok := muo.mutation.Drive(); ok {
		_spec.SetField(machine.FieldDrive, field.TypeString, value)
	}
	if value, ok := muo.mutation.Weight(); ok {
		_spec.SetField(machine.FieldWeight, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedWeight(); ok {
		_spec.AddField(machine.FieldWeight, field.TypeUint, value)
	}
	if value, ok := muo.mutation.LoadCapacity(); ok {
		_spec.SetField(machine.FieldLoadCapacity, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedLoadCapacity(); ok {
		_spec.AddField(machine.FieldLoadCapacity, field.TypeUint, value)
	}
	if value, ok := muo.mutation.RentPrice(); ok {
		_spec.SetField(machine.FieldRentPrice, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedRentPrice(); ok {
		_spec.AddField(machine.FieldRentPrice, field.TypeUint, value)
	}
	if muo.mutation.RentPriceCleared() {
		_spec.ClearField(machine.FieldRentPrice, field.TypeUint)
	}
	if value, ok := muo.mutation.SalePrice(); ok {
		_spec.SetField(machine.FieldSalePrice, field.TypeUint, value)
	}
	if value, ok := muo.mutation.AddedSalePrice(); ok {
		_spec.AddField(machine.FieldSalePrice, field.TypeUint, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(machine.FieldDescription, field.TypeString, value)
	}
	if value, ok := muo.mutation.Hidden(); ok {
		_spec.SetField(machine.FieldHidden, field.TypeBool, value)
	}
	if value, ok := muo.mutation.Deleted(); ok {
		_spec.SetField(machine.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := muo.mutation.NumberOfImages(); ok {
		_spec.SetField(machine.FieldNumberOfImages, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedNumberOfImages(); ok {
		_spec.AddField(machine.FieldNumberOfImages, field.TypeInt, value)
	}
	if muo.mutation.NumberOfImagesCleared() {
		_spec.ClearField(machine.FieldNumberOfImages, field.TypeInt)
	}
	if value, ok := muo.mutation.InternalDescription(); ok {
		_spec.SetField(machine.FieldInternalDescription, field.TypeString, value)
	}
	if muo.mutation.InternalDescriptionCleared() {
		_spec.ClearField(machine.FieldInternalDescription, field.TypeString)
	}
	if muo.mutation.MachineSpecificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMachineSpecificationIDs(); len(nodes) > 0 && !muo.mutation.MachineSpecificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MachineSpecificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineSpecificationTable,
			Columns: []string{machine.MachineSpecificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machinespecification.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MachineAccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMachineAccessoriesIDs(); len(nodes) > 0 && !muo.mutation.MachineAccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MachineAccessoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.MachineAccessoriesTable,
			Columns: []string{machine.MachineAccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: machineaccessories.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Machine{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
