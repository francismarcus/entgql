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

/*
	TODO: #1 A user shouldn't be able to follow themselfs @francismarcus
	This is also true for unfollowUser
*/
func (r *mutationResolver) FollowUser(ctx context.Context, input model.FollowUserInput) (*ent.User, error) {
	f := r.client.User.UpdateOneID(input.FollowID).AddFollowerIDs(input.UserID).AddFollowersCount(1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowsCount(1).SaveX(ctx)

	return f, nil
}

// TODO: #2 A user who is not following the user should not be able to remove following
func (r *mutationResolver) UnFollowUser(ctx context.Context, input model.UnFollowUserInput) (*ent.User, error) {
	f := r.client.User.UpdateOneID(input.FollowID).RemoveFollowerIDs(input.UserID).AddFollowersCount(-1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowsCount(-1).SaveX(ctx)

	return f, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.AuthPayload, error) {
	u, err := r.client.User.Query().Where(user.Username(input.Username)).Only(ctx)
	token := "asdasdsa"

	if err != nil {
		return nil, fmt.Errorf("failed loginUser: %v", err)
	}

	if u.Password != input.Password {
		return nil, fmt.Errorf("Wrong password: %v", input.Password)
	}

	return &model.AuthPayload{
		User:  u,
		Token: &token,
	}, nil
}

func (r *mutationResolver) SignupUser(ctx context.Context, input model.SignupUserInput) (*model.AuthPayload, error) {
	u, err := r.client.User.Create().SetUsername(input.Username).SetEmail(input.Email).SetPassword(input.Password).Save(ctx)
	token := "asdas"

	if err != nil {
		return nil, fmt.Errorf("failed signupUser: %v", err)
	}

	return &model.AuthPayload{
		User:  u,
		Token: &token,
	}, nil
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
	u, err := r.client.User.Query().Where(user.Username(input.Username)).Only(ctx)

	if err != nil {
		b = true
	}

	if u != nil {

		b = false
		return &b, fmt.Errorf("Username taken: %v", u)
	}

	return &b, nil
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return r.client.User.Query().
		WithPrograms().
		Paginate(ctx, after, first, before, last)
}

func (r *queryResolver) User(ctx context.Context, input model.ByIDInput) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(input.ID)).Only(ctx)
}

func (r *tweetResolver) User(ctx context.Context, obj *ent.Tweet) (*ent.User, error) {
	return obj.QueryCreator().Only(ctx)
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
	return obj.QueryTweets().Paginate(ctx, after, first, before, last)
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
