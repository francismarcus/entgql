package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/francismarcus/entgql/ent"
	"github.com/francismarcus/entgql/ent/user"
	"github.com/francismarcus/entgql/graph/generated"
	"github.com/francismarcus/entgql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	u, err := r.client.User.
		Create().
		SetUsername(input.Username).
		SetEmail(input.Email).
		SetPassword(input.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (r *mutationResolver) CreateProgram(ctx context.Context, input model.CreateProgramInput) (*ent.Program, error) {
	p, err := r.client.Program.Create().
		SetName(input.Name).
		SetCreatorID(input.Creator).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating program: %v", err)
	}

	return p, nil
}

func (r *mutationResolver) FollowUser(ctx context.Context, input model.FollowUserInput) (*ent.User, error) {
	f := r.client.User.UpdateOneID(input.FollowID).AddFollowerIDs(input.UserID).AddFollowersCount(1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowingCount(1).SaveX(ctx)

	return f, nil
}

func (r *mutationResolver) UnFollowUser(ctx context.Context, input model.UnFollowUserInput) (*ent.User, error) {
	f := r.client.User.UpdateOneID(input.FollowID).RemoveFollowerIDs(input.UserID).AddFollowersCount(-1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowingCount(-1).SaveX(ctx)

	return f, nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	node, err := r.client.Noder(ctx, id)
	if err == nil {
		return node, nil
	}
	var e *ent.NotFoundError
	if errors.As(err, &e) {
		err = nil
	}
	return nil, err
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *queryResolver) UsernameAvailable(ctx context.Context, input model.UsernameAvailableInput) (*bool, error) {
	var b bool
	u, err := r.client.User.Query().Where(user.Username(input.Username)).All(ctx)

	if err != nil {
		b = true
	}

	if u != nil {
		b = false
	}

	return &b, nil
}

func (r *queryResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SignupUser(ctx context.Context, input model.SignupUserInput) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return r.client.User.Query().
		WithPrograms().
		Paginate(ctx, after, first, before, last)
}

func (r *tweetResolver) User(ctx context.Context, obj *ent.Tweet) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Programs(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.ProgramConnection, error) {
	return obj.QueryPrograms().Paginate(ctx, after, first, before, last)
}

func (r *userResolver) Followers(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return obj.QueryFollowers().Paginate(ctx, after, first, before, last)
}

func (r *userResolver) Follows(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return obj.QueryFollowing().Paginate(ctx, after, first, before, last)
}

func (r *userResolver) Tweets(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.TweetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tweet returns generated.TweetResolver implementation.
func (r *Resolver) Tweet() generated.TweetResolver { return &tweetResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tweetResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
