// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/machinespecification"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineCreate is the builder for creating a Machine entity.
type MachineCreate struct {
	config
	mutation *MachineMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MachineCreate) SetName(s string) *MachineCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetYear sets the "year" field.
func (mc *MachineCreate) SetYear(u uint) *MachineCreate {
	mc.mutation.SetYear(u)
	return mc
}

// SetLiftHeight sets the "liftHeight" field.
func (mc *MachineCreate) SetLiftHeight(u uint) *MachineCreate {
	mc.mutation.SetLiftHeight(u)
	return mc
}

// SetMotorHours sets the "motorHours" field.
func (mc *MachineCreate) SetMotorHours(u uint) *MachineCreate {
	mc.mutation.SetMotorHours(u)
	return mc
}

// SetDrive sets the "drive" field.
func (mc *MachineCreate) SetDrive(s string) *MachineCreate {
	mc.mutation.SetDrive(s)
	return mc
}

// SetWeight sets the "weight" field.
func (mc *MachineCreate) SetWeight(u uint) *MachineCreate {
	mc.mutation.SetWeight(u)
	return mc
}

// SetLoadCapacity sets the "loadCapacity" field.
func (mc *MachineCreate) SetLoadCapacity(u uint) *MachineCreate {
	mc.mutation.SetLoadCapacity(u)
	return mc
}

// SetRentPrice sets the "rentPrice" field.
func (mc *MachineCreate) SetRentPrice(u uint) *MachineCreate {
	mc.mutation.SetRentPrice(u)
	return mc
}

// SetNillableRentPrice sets the "rentPrice" field if the given value is not nil.
func (mc *MachineCreate) SetNillableRentPrice(u *uint) *MachineCreate {
	if u != nil {
		mc.SetRentPrice(*u)
	}
	return mc
}

// SetSalePrice sets the "salePrice" field.
func (mc *MachineCreate) SetSalePrice(u uint) *MachineCreate {
	mc.mutation.SetSalePrice(u)
	return mc
}

// SetDescription sets the "description" field.
func (mc *MachineCreate) SetDescription(s string) *MachineCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetHidden sets the "hidden" field.
func (mc *MachineCreate) SetHidden(b bool) *MachineCreate {
	mc.mutation.SetHidden(b)
	return mc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (mc *MachineCreate) SetNillableHidden(b *bool) *MachineCreate {
	if b != nil {
		mc.SetHidden(*b)
	}
	return mc
}

// SetDeleted sets the "deleted" field.
func (mc *MachineCreate) SetDeleted(b bool) *MachineCreate {
	mc.mutation.SetDeleted(b)
	return mc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (mc *MachineCreate) SetNillableDeleted(b *bool) *MachineCreate {
	if b != nil {
		mc.SetDeleted(*b)
	}
	return mc
}

// SetNumberOfImages sets the "numberOfImages" field.
func (mc *MachineCreate) SetNumberOfImages(i int) *MachineCreate {
	mc.mutation.SetNumberOfImages(i)
	return mc
}

// SetNillableNumberOfImages sets the "numberOfImages" field if the given value is not nil.
func (mc *MachineCreate) SetNillableNumberOfImages(i *int) *MachineCreate {
	if i != nil {
		mc.SetNumberOfImages(*i)
	}
	return mc
}

// SetInternalDescription sets the "internalDescription" field.
func (mc *MachineCreate) SetInternalDescription(s string) *MachineCreate {
	mc.mutation.SetInternalDescription(s)
	return mc
}

// SetNillableInternalDescription sets the "internalDescription" field if the given value is not nil.
func (mc *MachineCreate) SetNillableInternalDescription(s *string) *MachineCreate {
	if s != nil {
		mc.SetInternalDescription(*s)
	}
	return mc
}

// AddMachineSpecificationIDs adds the "machineSpecification" edge to the MachineSpecification entity by IDs.
func (mc *MachineCreate) AddMachineSpecificationIDs(ids ...int) *MachineCreate {
	mc.mutation.AddMachineSpecificationIDs(ids...)
	return mc
}

// AddMachineSpecification adds the "machineSpecification" edges to the MachineSpecification entity.
func (mc *MachineCreate) AddMachineSpecification(m ...*MachineSpecification) *MachineCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMachineSpecificationIDs(ids...)
}

// AddMachineAccessoryIDs adds the "machineAccessories" edge to the MachineAccessories entity by IDs.
func (mc *MachineCreate) AddMachineAccessoryIDs(ids ...int) *MachineCreate {
	mc.mutation.AddMachineAccessoryIDs(ids...)
	return mc
}

// AddMachineAccessories adds the "machineAccessories" edges to the MachineAccessories entity.
func (mc *MachineCreate) AddMachineAccessories(m ...*MachineAccessories) *MachineCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMachineAccessoryIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (mc *MachineCreate) Mutation() *MachineMutation {
	return mc.mutation
}

// Save creates the Machine in the database.
func (mc *MachineCreate) Save(ctx context.Context) (*Machine, error) {
	var (
		err  error
		node *Machine
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (mc *MachineCreate) SaveX(ctx context.Context) *Machine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MachineCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MachineCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MachineCreate) defaults() {
	if _, ok := mc.mutation.Hidden(); !ok {
		v := machine.DefaultHidden
		mc.mutation.SetHidden(v)
	}
	if _, ok := mc.mutation.Deleted(); !ok {
		v := machine.DefaultDeleted
		mc.mutation.SetDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MachineCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Machine.name"`)}
	}
	if _, ok := mc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Machine.year"`)}
	}
	if _, ok := mc.mutation.LiftHeight(); !ok {
		return &ValidationError{Name: "liftHeight", err: errors.New(`ent: missing required field "Machine.liftHeight"`)}
	}
	if _, ok := mc.mutation.MotorHours(); !ok {
		return &ValidationError{Name: "motorHours", err: errors.New(`ent: missing required field "Machine.motorHours"`)}
	}
	if _, ok := mc.mutation.Drive(); !ok {
		return &ValidationError{Name: "drive", err: errors.New(`ent: missing required field "Machine.drive"`)}
	}
	if _, ok := mc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Machine.weight"`)}
	}
	if _, ok := mc.mutation.LoadCapacity(); !ok {
		return &ValidationError{Name: "loadCapacity", err: errors.New(`ent: missing required field "Machine.loadCapacity"`)}
	}
	if _, ok := mc.mutation.SalePrice(); !ok {
		return &ValidationError{Name: "salePrice", err: errors.New(`ent: missing required field "Machine.salePrice"`)}
	}
	if _, ok := mc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Machine.description"`)}
	}
	if v, ok := mc.mutation.Description(); ok {
		if err := machine.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Machine.description": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Hidden(); !ok {
		return &ValidationError{Name: "hidden", err: errors.New(`ent: missing required field "Machine.hidden"`)}
	}
	if _, ok := mc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Machine.deleted"`)}
	}
	if v, ok := mc.mutation.InternalDescription(); ok {
		if err := machine.InternalDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "internalDescription", err: fmt.Errorf(`ent: validator failed for field "Machine.internalDescription": %w`, err)}
		}
	}
	return nil
}

func (mc *MachineCreate) sqlSave(ctx context.Context) (*Machine, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MachineCreate) createSpec() (*Machine, *sqlgraph.CreateSpec) {
	var (
		_node = &Machine{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: machine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(machine.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Year(); ok {
		_spec.SetField(machine.FieldYear, field.TypeUint, value)
		_node.Year = value
	}
	if value, ok := mc.mutation.LiftHeight(); ok {
		_spec.SetField(machine.FieldLiftHeight, field.TypeUint, value)
		_node.LiftHeight = value
	}
	if value, ok := mc.mutation.MotorHours(); ok {
		_spec.SetField(machine.FieldMotorHours, field.TypeUint, value)
		_node.MotorHours = value
	}
	if value, ok := mc.mutation.Drive(); ok {
		_spec.SetField(machine.FieldDrive, field.TypeString, value)
		_node.Drive = value
	}
	if value, ok := mc.mutation.Weight(); ok {
		_spec.SetField(machine.FieldWeight, field.TypeUint, value)
		_node.Weight = value
	}
	if value, ok := mc.mutation.LoadCapacity(); ok {
		_spec.SetField(machine.FieldLoadCapacity, field.TypeUint, value)
		_node.LoadCapacity = value
	}
	if value, ok := mc.mutation.RentPrice(); ok {
		_spec.SetField(machine.FieldRentPrice, field.TypeUint, value)
		_node.RentPrice = value
	}
	if value, ok := mc.mutation.SalePrice(); ok {
		_spec.SetField(machine.FieldSalePrice, field.TypeUint, value)
		_node.SalePrice = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(machine.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.Hidden(); ok {
		_spec.SetField(machine.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := mc.mutation.Deleted(); ok {
		_spec.SetField(machine.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := mc.mutation.NumberOfImages(); ok {
		_spec.SetField(machine.FieldNumberOfImages, field.TypeInt, value)
		_node.NumberOfImages = value
	}
	if value, ok := mc.mutation.InternalDescription(); ok {
		_spec.SetField(machine.FieldInternalDescription, field.TypeString, value)
		_node.InternalDescription = value
	}
	if nodes := mc.mutation.MachineSpecificationIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MachineAccessoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MachineCreateBulk is the builder for creating many Machine entities in bulk.
type MachineCreateBulk struct {
	config
	builders []*MachineCreate
}

// Save creates the Machine entities in the database.
func (mcb *MachineCreateBulk) Save(ctx context.Context) ([]*Machine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Machine, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MachineMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MachineCreateBulk) SaveX(ctx context.Context) []*Machine {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MachineCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MachineCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
