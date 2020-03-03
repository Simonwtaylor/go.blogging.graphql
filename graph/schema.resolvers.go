// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/Simonwtaylor/blogging-gql/entities"
	"github.com/Simonwtaylor/blogging-gql/graph/generated"
	"github.com/Simonwtaylor/blogging-gql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := entities.User{
		Active:   input.Active,
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
	}

	isNew := r.DB.NewRecord(user)

	if !isNew {
		return nil, fmt.Errorf("unable to create user as it already contains an id, %v", user)
	}

	r.DB.Create(&user)

	return &model.User{
		ID:        fmt.Sprintf("%d", user.ID),
		CreatedAt: user.CreatedAt.Nanosecond(),
		Active:    user.Active,
		Email:     user.Email,
		Password:  user.Password,
		Username:  user.Username,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		&model.Todo{
			ID:   "1",
			Done: false,
			Text: "Hello world :D",
			User: &model.User{
				ID: "1",
			},
		},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []entities.User
	r.DB.Find(&users)
	var userModels []*model.User
	for _, user := range users {
		userModels = append(userModels, &model.User{
			Email:     user.Email,
			Username:  user.Username,
			Active:    user.Active,
			CreatedAt: user.CreatedAt.Nanosecond(),
			ID:        fmt.Sprintf("%v", user.ID),
		})
	}

	return userModels, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
