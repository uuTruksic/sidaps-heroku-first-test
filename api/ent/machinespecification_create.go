// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machinespecification"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineSpecificationCreate is the builder for creating a MachineSpecification entity.
type MachineSpecificationCreate struct {
	config
	mutation *MachineSpecificationMutation
	hooks    []Hook
}

// SetMachineNumber sets the "machineNumber" field.
func (msc *MachineSpecificationCreate) SetMachineNumber(i int) *MachineSpecificationCreate {
	msc.mutation.SetMachineNumber(i)
	return msc
}

// SetMachineType sets the "machineType" field.
func (msc *MachineSpecificationCreate) SetMachineType(s string) *MachineSpecificationCreate {
	msc.mutation.SetMachineType(s)
	return msc
}

// SetDrive sets the "drive" field.
func (msc *MachineSpecificationCreate) SetDrive(s string) *MachineSpecificationCreate {
	msc.mutation.SetDrive(s)
	return msc
}

// SetLoadCapacity sets the "loadCapacity" field.
func (msc *MachineSpecificationCreate) SetLoadCapacity(u uint) *MachineSpecificationCreate {
	msc.mutation.SetLoadCapacity(u)
	return msc
}

// SetYear sets the "year" field.
func (msc *MachineSpecificationCreate) SetYear(u uint) *MachineSpecificationCreate {
	msc.mutation.SetYear(u)
	return msc
}

// SetMotorHours sets the "motorHours" field.
func (msc *MachineSpecificationCreate) SetMotorHours(u uint) *MachineSpecificationCreate {
	msc.mutation.SetMotorHours(u)
	return msc
}

// SetPassingHeight sets the "passingHeight" field.
func (msc *MachineSpecificationCreate) SetPassingHeight(u uint) *MachineSpecificationCreate {
	msc.mutation.SetPassingHeight(u)
	return msc
}

// SetNillablePassingHeight sets the "passingHeight" field if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillablePassingHeight(u *uint) *MachineSpecificationCreate {
	if u != nil {
		msc.SetPassingHeight(*u)
	}
	return msc
}

// SetLiftHeight sets the "liftHeight" field.
func (msc *MachineSpecificationCreate) SetLiftHeight(u uint) *MachineSpecificationCreate {
	msc.mutation.SetLiftHeight(u)
	return msc
}

// SetWeight sets the "weight" field.
func (msc *MachineSpecificationCreate) SetWeight(u uint) *MachineSpecificationCreate {
	msc.mutation.SetWeight(u)
	return msc
}

// SetForks sets the "forks" field.
func (msc *MachineSpecificationCreate) SetForks(s string) *MachineSpecificationCreate {
	msc.mutation.SetForks(s)
	return msc
}

// SetMastType sets the "mastType" field.
func (msc *MachineSpecificationCreate) SetMastType(s string) *MachineSpecificationCreate {
	msc.mutation.SetMastType(s)
	return msc
}

// SetNillableMastType sets the "mastType" field if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillableMastType(s *string) *MachineSpecificationCreate {
	if s != nil {
		msc.SetMastType(*s)
	}
	return msc
}

// SetAdditionalAttributes sets the "additionalAttributes" field.
func (msc *MachineSpecificationCreate) SetAdditionalAttributes(s string) *MachineSpecificationCreate {
	msc.mutation.SetAdditionalAttributes(s)
	return msc
}

// SetNillableAdditionalAttributes sets the "additionalAttributes" field if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillableAdditionalAttributes(s *string) *MachineSpecificationCreate {
	if s != nil {
		msc.SetAdditionalAttributes(*s)
	}
	return msc
}

// SetEquipment sets the "equipment" field.
func (msc *MachineSpecificationCreate) SetEquipment(s string) *MachineSpecificationCreate {
	msc.mutation.SetEquipment(s)
	return msc
}

// SetNillableEquipment sets the "equipment" field if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillableEquipment(s *string) *MachineSpecificationCreate {
	if s != nil {
		msc.SetEquipment(*s)
	}
	return msc
}

// SetEngine sets the "engine" field.
func (msc *MachineSpecificationCreate) SetEngine(s string) *MachineSpecificationCreate {
	msc.mutation.SetEngine(s)
	return msc
}

// SetHydraulicControl sets the "hydraulicControl" field.
func (msc *MachineSpecificationCreate) SetHydraulicControl(s string) *MachineSpecificationCreate {
	msc.mutation.SetHydraulicControl(s)
	return msc
}

// SetDriveControl sets the "driveControl" field.
func (msc *MachineSpecificationCreate) SetDriveControl(s string) *MachineSpecificationCreate {
	msc.mutation.SetDriveControl(s)
	return msc
}

// SetCabin sets the "cabin" field.
func (msc *MachineSpecificationCreate) SetCabin(s string) *MachineSpecificationCreate {
	msc.mutation.SetCabin(s)
	return msc
}

// SetLights sets the "lights" field.
func (msc *MachineSpecificationCreate) SetLights(s string) *MachineSpecificationCreate {
	msc.mutation.SetLights(s)
	return msc
}

// SetBattery sets the "battery" field.
func (msc *MachineSpecificationCreate) SetBattery(s string) *MachineSpecificationCreate {
	msc.mutation.SetBattery(s)
	return msc
}

// SetNillableBattery sets the "battery" field if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillableBattery(s *string) *MachineSpecificationCreate {
	if s != nil {
		msc.SetBattery(*s)
	}
	return msc
}

// SetMachineID sets the "machine" edge to the Machine entity by ID.
func (msc *MachineSpecificationCreate) SetMachineID(id int) *MachineSpecificationCreate {
	msc.mutation.SetMachineID(id)
	return msc
}

// SetNillableMachineID sets the "machine" edge to the Machine entity by ID if the given value is not nil.
func (msc *MachineSpecificationCreate) SetNillableMachineID(id *int) *MachineSpecificationCreate {
	if id != nil {
		msc = msc.SetMachineID(*id)
	}
	return msc
}

// SetMachine sets the "machine" edge to the Machine entity.
func (msc *MachineSpecificationCreate) SetMachine(m *Machine) *MachineSpecificationCreate {
	return msc.SetMachineID(m.ID)
}

// Mutation returns the MachineSpecificationMutation object of the builder.
func (msc *MachineSpecificationCreate) Mutation() *MachineSpecificationMutation {
	return msc.mutation
}

// Save creates the MachineSpecification in the database.
func (msc *MachineSpecificationCreate) Save(ctx context.Context) (*MachineSpecification, error) {
	var (
		err  error
		node *MachineSpecification
	)
	if len(msc.hooks) == 0 {
		if err = msc.check(); err != nil {
			return nil, err
		}
		node, err = msc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineSpecificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = msc.check(); err != nil {
				return nil, err
			}
			msc.mutation = mutation
			if node, err = msc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(msc.hooks) - 1; i >= 0; i-- {
			if msc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = msc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, msc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MachineSpecification)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MachineSpecificationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (msc *MachineSpecificationCreate) SaveX(ctx context.Context) *MachineSpecification {
	v, err := msc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (msc *MachineSpecificationCreate) Exec(ctx context.Context) error {
	_, err := msc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msc *MachineSpecificationCreate) ExecX(ctx context.Context) {
	if err := msc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msc *MachineSpecificationCreate) check() error {
	if _, ok := msc.mutation.MachineNumber(); !ok {
		return &ValidationError{Name: "machineNumber", err: errors.New(`ent: missing required field "MachineSpecification.machineNumber"`)}
	}
	if _, ok := msc.mutation.MachineType(); !ok {
		return &ValidationError{Name: "machineType", err: errors.New(`ent: missing required field "MachineSpecification.machineType"`)}
	}
	if _, ok := msc.mutation.Drive(); !ok {
		return &ValidationError{Name: "drive", err: errors.New(`ent: missing required field "MachineSpecification.drive"`)}
	}
	if _, ok := msc.mutation.LoadCapacity(); !ok {
		return &ValidationError{Name: "loadCapacity", err: errors.New(`ent: missing required field "MachineSpecification.loadCapacity"`)}
	}
	if _, ok := msc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "MachineSpecification.year"`)}
	}
	if _, ok := msc.mutation.MotorHours(); !ok {
		return &ValidationError{Name: "motorHours", err: errors.New(`ent: missing required field "MachineSpecification.motorHours"`)}
	}
	if _, ok := msc.mutation.LiftHeight(); !ok {
		return &ValidationError{Name: "liftHeight", err: errors.New(`ent: missing required field "MachineSpecification.liftHeight"`)}
	}
	if _, ok := msc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "MachineSpecification.weight"`)}
	}
	if _, ok := msc.mutation.Forks(); !ok {
		return &ValidationError{Name: "forks", err: errors.New(`ent: missing required field "MachineSpecification.forks"`)}
	}
	if _, ok := msc.mutation.Engine(); !ok {
		return &ValidationError{Name: "engine", err: errors.New(`ent: missing required field "MachineSpecification.engine"`)}
	}
	if _, ok := msc.mutation.HydraulicControl(); !ok {
		return &ValidationError{Name: "hydraulicControl", err: errors.New(`ent: missing required field "MachineSpecification.hydraulicControl"`)}
	}
	if _, ok := msc.mutation.DriveControl(); !ok {
		return &ValidationError{Name: "driveControl", err: errors.New(`ent: missing required field "MachineSpecification.driveControl"`)}
	}
	if _, ok := msc.mutation.Cabin(); !ok {
		return &ValidationError{Name: "cabin", err: errors.New(`ent: missing required field "MachineSpecification.cabin"`)}
	}
	if _, ok := msc.mutation.Lights(); !ok {
		return &ValidationError{Name: "lights", err: errors.New(`ent: missing required field "MachineSpecification.lights"`)}
	}
	return nil
}

func (msc *MachineSpecificationCreate) sqlSave(ctx context.Context) (*MachineSpecification, error) {
	_node, _spec := msc.createSpec()
	if err := sqlgraph.CreateNode(ctx, msc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (msc *MachineSpecificationCreate) createSpec() (*MachineSpecification, *sqlgraph.CreateSpec) {
	var (
		_node = &MachineSpecification{config: msc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: machinespecification.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machinespecification.FieldID,
			},
		}
	)
	if value, ok := msc.mutation.MachineNumber(); ok {
		_spec.SetField(machinespecification.FieldMachineNumber, field.TypeInt, value)
		_node.MachineNumber = value
	}
	if value, ok := msc.mutation.MachineType(); ok {
		_spec.SetField(machinespecification.FieldMachineType, field.TypeString, value)
		_node.MachineType = value
	}
	if value, ok := msc.mutation.Drive(); ok {
		_spec.SetField(machinespecification.FieldDrive, field.TypeString, value)
		_node.Drive = value
	}
	if value, ok := msc.mutation.LoadCapacity(); ok {
		_spec.SetField(machinespecification.FieldLoadCapacity, field.TypeUint, value)
		_node.LoadCapacity = value
	}
	if value, ok := msc.mutation.Year(); ok {
		_spec.SetField(machinespecification.FieldYear, field.TypeUint, value)
		_node.Year = value
	}
	if value, ok := msc.mutation.MotorHours(); ok {
		_spec.SetField(machinespecification.FieldMotorHours, field.TypeUint, value)
		_node.MotorHours = value
	}
	if value, ok := msc.mutation.PassingHeight(); ok {
		_spec.SetField(machinespecification.FieldPassingHeight, field.TypeUint, value)
		_node.PassingHeight = value
	}
	if value, ok := msc.mutation.LiftHeight(); ok {
		_spec.SetField(machinespecification.FieldLiftHeight, field.TypeUint, value)
		_node.LiftHeight = value
	}
	if value, ok := msc.mutation.Weight(); ok {
		_spec.SetField(machinespecification.FieldWeight, field.TypeUint, value)
		_node.Weight = value
	}
	if value, ok := msc.mutation.Forks(); ok {
		_spec.SetField(machinespecification.FieldForks, field.TypeString, value)
		_node.Forks = value
	}
	if value, ok := msc.mutation.MastType(); ok {
		_spec.SetField(machinespecification.FieldMastType, field.TypeString, value)
		_node.MastType = value
	}
	if value, ok := msc.mutation.AdditionalAttributes(); ok {
		_spec.SetField(machinespecification.FieldAdditionalAttributes, field.TypeString, value)
		_node.AdditionalAttributes = value
	}
	if value, ok := msc.mutation.Equipment(); ok {
		_spec.SetField(machinespecification.FieldEquipment, field.TypeString, value)
		_node.Equipment = value
	}
	if value, ok := msc.mutation.Engine(); ok {
		_spec.SetField(machinespecification.FieldEngine, field.TypeString, value)
		_node.Engine = value
	}
	if value, ok := msc.mutation.HydraulicControl(); ok {
		_spec.SetField(machinespecification.FieldHydraulicControl, field.TypeString, value)
		_node.HydraulicControl = value
	}
	if value, ok := msc.mutation.DriveControl(); ok {
		_spec.SetField(machinespecification.FieldDriveControl, field.TypeString, value)
		_node.DriveControl = value
	}
	if value, ok := msc.mutation.Cabin(); ok {
		_spec.SetField(machinespecification.FieldCabin, field.TypeString, value)
		_node.Cabin = value
	}
	if value, ok := msc.mutation.Lights(); ok {
		_spec.SetField(machinespecification.FieldLights, field.TypeString, value)
		_node.Lights = value
	}
	if value, ok := msc.mutation.Battery(); ok {
		_spec.SetField(machinespecification.FieldBattery, field.TypeString, value)
		_node.Battery = value
	}
	if nodes := msc.mutation.MachineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   machinespecification.MachineTable,
			Columns: []string{machinespecification.MachineColumn},
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
		_node.machine_machine_specification = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MachineSpecificationCreateBulk is the builder for creating many MachineSpecification entities in bulk.
type MachineSpecificationCreateBulk struct {
	config
	builders []*MachineSpecificationCreate
}

// Save creates the MachineSpecification entities in the database.
func (mscb *MachineSpecificationCreateBulk) Save(ctx context.Context) ([]*MachineSpecification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mscb.builders))
	nodes := make([]*MachineSpecification, len(mscb.builders))
	mutators := make([]Mutator, len(mscb.builders))
	for i := range mscb.builders {
		func(i int, root context.Context) {
			builder := mscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MachineSpecificationMutation)
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
					_, err = mutators[i+1].Mutate(root, mscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mscb *MachineSpecificationCreateBulk) SaveX(ctx context.Context) []*MachineSpecification {
	v, err := mscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mscb *MachineSpecificationCreateBulk) Exec(ctx context.Context) error {
	_, err := mscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mscb *MachineSpecificationCreateBulk) ExecX(ctx context.Context) {
	if err := mscb.Exec(ctx); err != nil {
		panic(err)
	}
}
