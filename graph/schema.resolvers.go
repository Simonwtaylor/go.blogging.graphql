// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/Simonwtaylor/blogging-gql/entities"
	"github.com/Simonwtaylor/blogging-gql/graph/generated"
	"github.com/Simonwtaylor/blogging-gql/graph/model"
)

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

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := entities.Post{
		Content:    input.Content,
		DatePosted: time.Now(),
	}
	var user entities.User
	r.DB.First(&user)

	post.User = user

	isNew := r.DB.NewRecord(post)

	if !isNew {
		return nil, fmt.Errorf("unable to create user as it already contains an id, %v", user)
	}

	r.DB.Create(&post)

	return &model.Post{
		ID:      fmt.Sprintf("%d", post.ID),
		Content: post.Content,
		User: &model.User{
			ID:        fmt.Sprintf("%d", post.User.ID),
			CreatedAt: post.User.CreatedAt.Nanosecond(),
			Active:    post.User.Active,
			Email:     post.User.Email,
			Password:  post.User.Password,
			Username:  post.User.Username,
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

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []entities.Post
	r.DB.Preload("User").Find(&posts)
	var postModels []*model.Post
	for _, post := range posts {
		postModels = append(postModels, &model.Post{
			Content: post.Content,
			ID:      fmt.Sprintf("%v", post.ID),
			User: &model.User{
				Email:     post.User.Email,
				Username:  post.User.Username,
				Active:    post.User.Active,
				CreatedAt: post.User.CreatedAt.Nanosecond(),
				ID:        fmt.Sprintf("%v", post.User.ID),
			},
		})
	}

	return postModels, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
