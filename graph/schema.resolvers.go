package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/francismarcus/entgql/ent"
	"github.com/francismarcus/entgql/graph/generated"
	"github.com/francismarcus/entgql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	u, err := r.Client.User.
		Create().
		SetUsername("francismarcus").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (r *mutationResolver) CreateProgram(ctx context.Context, input model.CreateProgramInput) (*ent.Program, error) {
	p, err := r.Client.Program.Create().
		SetName(input.Name).
		SetCreatorID(input.Creator).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating program: %v", err)
	}

	return p, nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	node, err := r.Client.Noder(ctx, id)
	if err == nil {
		return node, nil
	}
	var e *ent.NotFoundError
	if errors.As(err, &e) {
		err = nil
	}
	return nil, err
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return r.Client.User.Query().
		WithPrograms().
		Paginate(ctx, after, first, before, last)
}

func (r *userResolver) Programs(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.ProgramConnection, error) {
	return obj.QueryPrograms().Paginate(ctx, after, first, before, last)

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
