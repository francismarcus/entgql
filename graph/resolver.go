package graph

import (
	"github.com/francismarcus/entgql/ent"
	"github.com/francismarcus/entgql/graph/generated"
)

// New adds ent.Client to resolvers
func New(client *ent.Client) generated.Config {
	return generated.Config{
		Resolvers: &Resolver{
			client: client,
		},
	}
}

// Resolver has client from ent
type Resolver struct {
	client *ent.Client
}
