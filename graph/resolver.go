package graph

import (
	"github.com/francismarcus/entgql/ent"
	// drivers for pg
	_ "github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver has client from ent
type Resolver struct {
	Client *ent.Client
}