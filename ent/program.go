// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/francismarcus/entgql/ent/program"
	"github.com/francismarcus/entgql/ent/user"
)

// Program is the model entity for the Program schema.
type Program struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProgramQuery when eager-loading is set.
	Edges         ProgramEdges `json:"edges"`
	user_programs *int
}

// ProgramEdges holds the relations/edges for other nodes in the graph.
type ProgramEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProgramEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// The edge creator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Program) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Program) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_programs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Program fields.
func (pr *Program) assignValues(values ...interface{}) error {
	if m, n := len(values), len(program.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pr.Name = value.String
	}
	values = values[1:]
	if len(values) == len(program.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_programs", value)
		} else if value.Valid {
			pr.user_programs = new(int)
			*pr.user_programs = int(value.Int64)
		}
	}
	return nil
}

// QueryCreator queries the creator edge of the Program.
func (pr *Program) QueryCreator() *UserQuery {
	return (&ProgramClient{config: pr.config}).QueryCreator(pr)
}

// Update returns a builder for updating this Program.
// Note that, you need to call Program.Unwrap() before calling this method, if this Program
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Program) Update() *ProgramUpdateOne {
	return (&ProgramClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Program) Unwrap() *Program {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Program is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Program) String() string {
	var builder strings.Builder
	builder.WriteString("Program(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Programs is a parsable slice of Program.
type Programs []*Program

func (pr Programs) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}