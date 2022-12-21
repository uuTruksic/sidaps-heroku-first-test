// Code generated by ent, DO NOT EDIT.

package machinespecification

const (
	// Label holds the string label denoting the machinespecification type in the database.
	Label = "machine_specification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMachineNumber holds the string denoting the machinenumber field in the database.
	FieldMachineNumber = "machine_number"
	// FieldMachineType holds the string denoting the machinetype field in the database.
	FieldMachineType = "machine_type"
	// FieldDrive holds the string denoting the drive field in the database.
	FieldDrive = "drive"
	// FieldLoadCapacity holds the string denoting the loadcapacity field in the database.
	FieldLoadCapacity = "load_capacity"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldMotorHours holds the string denoting the motorhours field in the database.
	FieldMotorHours = "motor_hours"
	// FieldPassingHeight holds the string denoting the passingheight field in the database.
	FieldPassingHeight = "passing_height"
	// FieldLiftHeight holds the string denoting the liftheight field in the database.
	FieldLiftHeight = "lift_height"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldForks holds the string denoting the forks field in the database.
	FieldForks = "forks"
	// FieldMastType holds the string denoting the masttype field in the database.
	FieldMastType = "mast_type"
	// FieldAdditionalAttributes holds the string denoting the additionalattributes field in the database.
	FieldAdditionalAttributes = "additional_attributes"
	// FieldEquipment holds the string denoting the equipment field in the database.
	FieldEquipment = "equipment"
	// FieldEngine holds the string denoting the engine field in the database.
	FieldEngine = "engine"
	// FieldHydraulicControl holds the string denoting the hydrauliccontrol field in the database.
	FieldHydraulicControl = "hydraulic_control"
	// FieldDriveControl holds the string denoting the drivecontrol field in the database.
	FieldDriveControl = "drive_control"
	// FieldCabin holds the string denoting the cabin field in the database.
	FieldCabin = "cabin"
	// FieldLights holds the string denoting the lights field in the database.
	FieldLights = "lights"
	// FieldBattery holds the string denoting the battery field in the database.
	FieldBattery = "battery"
	// EdgeMachine holds the string denoting the machine edge name in mutations.
	EdgeMachine = "machine"
	// Table holds the table name of the machinespecification in the database.
	Table = "machine_specifications"
	// MachineTable is the table that holds the machine relation/edge.
	MachineTable = "machine_specifications"
	// MachineInverseTable is the table name for the Machine entity.
	// It exists in this package in order to avoid circular dependency with the "machine" package.
	MachineInverseTable = "machines"
	// MachineColumn is the table column denoting the machine relation/edge.
	MachineColumn = "machine_machine_specification"
)

// Columns holds all SQL columns for machinespecification fields.
var Columns = []string{
	FieldID,
	FieldMachineNumber,
	FieldMachineType,
	FieldDrive,
	FieldLoadCapacity,
	FieldYear,
	FieldMotorHours,
	FieldPassingHeight,
	FieldLiftHeight,
	FieldWeight,
	FieldForks,
	FieldMastType,
	FieldAdditionalAttributes,
	FieldEquipment,
	FieldEngine,
	FieldHydraulicControl,
	FieldDriveControl,
	FieldCabin,
	FieldLights,
	FieldBattery,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "machine_specifications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"machine_machine_specification",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}