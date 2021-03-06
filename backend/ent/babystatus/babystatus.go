// Code generated by entc, DO NOT EDIT.

package babystatus

const (
	// Label holds the string label denoting the babystatus type in the database.
	Label = "babystatus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBabystatusName holds the string denoting the babystatus_name field in the database.
	FieldBabystatusName = "babystatus_name"

	// EdgeAntenatals holds the string denoting the antenatals edge name in mutations.
	EdgeAntenatals = "antenatals"

	// Table holds the table name of the babystatus in the database.
	Table = "babystatuses"
	// AntenatalsTable is the table the holds the antenatals relation/edge.
	AntenatalsTable = "antenatals"
	// AntenatalsInverseTable is the table name for the Antenatal entity.
	// It exists in this package in order to avoid circular dependency with the "antenatal" package.
	AntenatalsInverseTable = "antenatals"
	// AntenatalsColumn is the table column denoting the antenatals relation/edge.
	AntenatalsColumn = "babystatus_id"
)

// Columns holds all SQL columns for babystatus fields.
var Columns = []string{
	FieldID,
	FieldBabystatusName,
}

var (
	// BabystatusNameValidator is a validator for the "babystatus_name" field. It is called by the builders before save.
	BabystatusNameValidator func(string) error
)
