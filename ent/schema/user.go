package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.String("email").
			Unique(),
		field.String("password"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("follows_count").
			Default(0),
		field.Int("followers_count").
			Default(0),
		field.Int("tweets_count").
			Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
		edge.To("programs", Program.Type),
		edge.To("tweets", Tweet.Type),
	}
}
