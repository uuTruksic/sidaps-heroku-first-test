// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/schema"
	"api/ent/session"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	machineFields := schema.Machine{}.Fields()
	_ = machineFields
	// machineDescDescription is the schema descriptor for description field.
	machineDescDescription := machineFields[9].Descriptor()
	// machine.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	machine.DescriptionValidator = machineDescDescription.Validators[0].(func(string) error)
	// machineDescHidden is the schema descriptor for hidden field.
	machineDescHidden := machineFields[10].Descriptor()
	// machine.DefaultHidden holds the default value on creation for the hidden field.
	machine.DefaultHidden = machineDescHidden.Default.(bool)
	// machineDescDeleted is the schema descriptor for deleted field.
	machineDescDeleted := machineFields[11].Descriptor()
	// machine.DefaultDeleted holds the default value on creation for the deleted field.
	machine.DefaultDeleted = machineDescDeleted.Default.(bool)
	// machineDescInternalDescription is the schema descriptor for internalDescription field.
	machineDescInternalDescription := machineFields[13].Descriptor()
	// machine.InternalDescriptionValidator is a validator for the "internalDescription" field. It is called by the builders before save.
	machine.InternalDescriptionValidator = machineDescInternalDescription.Validators[0].(func(string) error)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[4].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[5].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
}
