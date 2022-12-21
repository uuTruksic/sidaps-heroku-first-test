// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineAccessoriesCreate is the builder for creating a MachineAccessories entity.
type MachineAccessoriesCreate struct {
	config
	mutation *MachineAccessoriesMutation
	hooks    []Hook
}

// SetName1 sets the "name1" field.
func (mac *MachineAccessoriesCreate) SetName1(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName1(s)
	return mac
}

// SetNillableName1 sets the "name1" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName1(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName1(*s)
	}
	return mac
}

// SetName2 sets the "name2" field.
func (mac *MachineAccessoriesCreate) SetName2(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName2(s)
	return mac
}

// SetNillableName2 sets the "name2" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName2(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName2(*s)
	}
	return mac
}

// SetName3 sets the "name3" field.
func (mac *MachineAccessoriesCreate) SetName3(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName3(s)
	return mac
}

// SetNillableName3 sets the "name3" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName3(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName3(*s)
	}
	return mac
}

// SetName4 sets the "name4" field.
func (mac *MachineAccessoriesCreate) SetName4(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName4(s)
	return mac
}

// SetNillableName4 sets the "name4" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName4(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName4(*s)
	}
	return mac
}

// SetName5 sets the "name5" field.
func (mac *MachineAccessoriesCreate) SetName5(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName5(s)
	return mac
}

// SetNillableName5 sets the "name5" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName5(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName5(*s)
	}
	return mac
}

// SetName6 sets the "name6" field.
func (mac *MachineAccessoriesCreate) SetName6(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName6(s)
	return mac
}

// SetNillableName6 sets the "name6" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName6(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName6(*s)
	}
	return mac
}

// SetName7 sets the "name7" field.
func (mac *MachineAccessoriesCreate) SetName7(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName7(s)
	return mac
}

// SetNillableName7 sets the "name7" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName7(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName7(*s)
	}
	return mac
}

// SetName8 sets the "name8" field.
func (mac *MachineAccessoriesCreate) SetName8(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName8(s)
	return mac
}

// SetNillableName8 sets the "name8" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName8(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName8(*s)
	}
	return mac
}

// SetName9 sets the "name9" field.
func (mac *MachineAccessoriesCreate) SetName9(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName9(s)
	return mac
}

// SetNillableName9 sets the "name9" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName9(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName9(*s)
	}
	return mac
}

// SetName10 sets the "name10" field.
func (mac *MachineAccessoriesCreate) SetName10(s string) *MachineAccessoriesCreate {
	mac.mutation.SetName10(s)
	return mac
}

// SetNillableName10 sets the "name10" field if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableName10(s *string) *MachineAccessoriesCreate {
	if s != nil {
		mac.SetName10(*s)
	}
	return mac
}

// SetMachineID sets the "machine" edge to the Machine entity by ID.
func (mac *MachineAccessoriesCreate) SetMachineID(id int) *MachineAccessoriesCreate {
	mac.mutation.SetMachineID(id)
	return mac
}

// SetNillableMachineID sets the "machine" edge to the Machine entity by ID if the given value is not nil.
func (mac *MachineAccessoriesCreate) SetNillableMachineID(id *int) *MachineAccessoriesCreate {
	if id != nil {
		mac = mac.SetMachineID(*id)
	}
	return mac
}

// SetMachine sets the "machine" edge to the Machine entity.
func (mac *MachineAccessoriesCreate) SetMachine(m *Machine) *MachineAccessoriesCreate {
	return mac.SetMachineID(m.ID)
}

// Mutation returns the MachineAccessoriesMutation object of the builder.
func (mac *MachineAccessoriesCreate) Mutation() *MachineAccessoriesMutation {
	return mac.mutation
}

// Save creates the MachineAccessories in the database.
func (mac *MachineAccessoriesCreate) Save(ctx context.Context) (*MachineAccessories, error) {
	var (
		err  error
		node *MachineAccessories
	)
	if len(mac.hooks) == 0 {
		if err = mac.check(); err != nil {
			return nil, err
		}
		node, err = mac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineAccessoriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mac.check(); err != nil {
				return nil, err
			}
			mac.mutation = mutation
			if node, err = mac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mac.hooks) - 1; i >= 0; i-- {
			if mac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (mac *MachineAccessoriesCreate) SaveX(ctx context.Context) *MachineAccessories {
	v, err := mac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mac *MachineAccessoriesCreate) Exec(ctx context.Context) error {
	_, err := mac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mac *MachineAccessoriesCreate) ExecX(ctx context.Context) {
	if err := mac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mac *MachineAccessoriesCreate) check() error {
	return nil
}

func (mac *MachineAccessoriesCreate) sqlSave(ctx context.Context) (*MachineAccessories, error) {
	_node, _spec := mac.createSpec()
	if err := sqlgraph.CreateNode(ctx, mac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mac *MachineAccessoriesCreate) createSpec() (*MachineAccessories, *sqlgraph.CreateSpec) {
	var (
		_node = &MachineAccessories{config: mac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: machineaccessories.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machineaccessories.FieldID,
			},
		}
	)
	if value, ok := mac.mutation.Name1(); ok {
		_spec.SetField(machineaccessories.FieldName1, field.TypeString, value)
		_node.Name1 = value
	}
	if value, ok := mac.mutation.Name2(); ok {
		_spec.SetField(machineaccessories.FieldName2, field.TypeString, value)
		_node.Name2 = value
	}
	if value, ok := mac.mutation.Name3(); ok {
		_spec.SetField(machineaccessories.FieldName3, field.TypeString, value)
		_node.Name3 = value
	}
	if value, ok := mac.mutation.Name4(); ok {
		_spec.SetField(machineaccessories.FieldName4, field.TypeString, value)
		_node.Name4 = value
	}
	if value, ok := mac.mutation.Name5(); ok {
		_spec.SetField(machineaccessories.FieldName5, field.TypeString, value)
		_node.Name5 = value
	}
	if value, ok := mac.mutation.Name6(); ok {
		_spec.SetField(machineaccessories.FieldName6, field.TypeString, value)
		_node.Name6 = value
	}
	if value, ok := mac.mutation.Name7(); ok {
		_spec.SetField(machineaccessories.FieldName7, field.TypeString, value)
		_node.Name7 = value
	}
	if value, ok := mac.mutation.Name8(); ok {
		_spec.SetField(machineaccessories.FieldName8, field.TypeString, value)
		_node.Name8 = value
	}
	if value, ok := mac.mutation.Name9(); ok {
		_spec.SetField(machineaccessories.FieldName9, field.TypeString, value)
		_node.Name9 = value
	}
	if value, ok := mac.mutation.Name10(); ok {
		_spec.SetField(machineaccessories.FieldName10, field.TypeString, value)
		_node.Name10 = value
	}
	if nodes := mac.mutation.MachineIDs(); len(nodes) > 0 {
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
		_node.machine_machine_accessories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MachineAccessoriesCreateBulk is the builder for creating many MachineAccessories entities in bulk.
type MachineAccessoriesCreateBulk struct {
	config
	builders []*MachineAccessoriesCreate
}

// Save creates the MachineAccessories entities in the database.
func (macb *MachineAccessoriesCreateBulk) Save(ctx context.Context) ([]*MachineAccessories, error) {
	specs := make([]*sqlgraph.CreateSpec, len(macb.builders))
	nodes := make([]*MachineAccessories, len(macb.builders))
	mutators := make([]Mutator, len(macb.builders))
	for i := range macb.builders {
		func(i int, root context.Context) {
			builder := macb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MachineAccessoriesMutation)
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
					_, err = mutators[i+1].Mutate(root, macb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, macb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, macb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (macb *MachineAccessoriesCreateBulk) SaveX(ctx context.Context) []*MachineAccessories {
	v, err := macb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (macb *MachineAccessoriesCreateBulk) Exec(ctx context.Context) error {
	_, err := macb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (macb *MachineAccessoriesCreateBulk) ExecX(ctx context.Context) {
	if err := macb.Exec(ctx); err != nil {
		panic(err)
	}
}
