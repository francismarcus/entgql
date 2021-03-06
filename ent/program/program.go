// Code generated by entc, DO NOT EDIT.

package program

const (
	// Label holds the string label denoting the program type in the database.
	Label = "program"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"

	// Table holds the table name of the program in the database.
	Table = "programs"
	// CreatorTable is the table the holds the creator relation/edge.
	CreatorTable = "programs"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "user_programs"
)

// Columns holds all SQL columns for program fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Program type.
var ForeignKeys = []string{
	"user_programs",
}
