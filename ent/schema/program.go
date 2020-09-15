package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Program holds the schema definition for the Program entity.
type Program struct {
	ent.Schema
}

// Fields of the Program
func (Program) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Program
func (Program) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("programs").
			Unique(),
	}
}
