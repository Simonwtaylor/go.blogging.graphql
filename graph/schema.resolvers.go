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
	"github.com/biezhi/gorm-paginator/pagination"
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

func (r *mutationResolver) CreateBike(ctx context.Context, input model.NewBike) (*model.Bike, error) {
	bike := entities.Bike{
		Active: input.Active,
		Make:   input.Make,
		Model:  input.Model,
		Price:  input.Price,
		Reg:    input.Reg,
	}

	isNew := r.DB.NewRecord(bike)

	if !isNew {
		return nil, fmt.Errorf("unable to create user as it already contains an id, %v", bike)
	}

	r.DB.Create(&bike)

	return &model.Bike{
		ID:        fmt.Sprintf("%v", bike.ID),
		Active:    bike.Active,
		Make:      bike.Make,
		Model:     bike.Model,
		Reg:       bike.Reg,
		Price:     bike.Price,
		CreatedAt: bike.CreatedAt.Nanosecond(),
		UpdatedAt: bike.UpdatedAt.Nanosecond(),
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, input model.PaginatedQuery) (*model.UserQuery, error) {
	var users []entities.User
	db := r.DB.Find(&users)
	count := len(users)
	pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  input.Cursor,
		Limit: input.PageSize,
	}, &users)

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

	return &model.UserQuery{
		Data: userModels,
		Meta: &model.Meta{
			Cursor:     input.Cursor,
			TotalItems: count,
		},
	}, nil
}

func (r *queryResolver) Posts(ctx context.Context, input model.PaginatedQuery) (*model.PostQuery, error) {
	var posts []entities.Post
	db := r.DB.Preload("User").Find(&posts)
	count := len(posts)
	pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  input.Cursor,
		Limit: input.PageSize,
	}, &posts)

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

	return &model.PostQuery{
		Data: postModels,
		Meta: &model.Meta{
			Cursor:     input.Cursor,
			TotalItems: count,
		},
	}, nil
}

func (r *queryResolver) Bikes(ctx context.Context, input model.PaginatedQuery) (*model.BikeQuery, error) {
	var bikes []entities.Bike
	db := r.DB.Find(&bikes)
	count := len(bikes)
	pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  input.Cursor,
		Limit: input.PageSize,
	}, &bikes)

	var bikeModels []*model.Bike
	for _, bike := range bikes {
		bikeModels = append(bikeModels, &model.Bike{
			ID:        fmt.Sprintf("%v", bike.ID),
			Active:    bike.Active,
			Make:      bike.Make,
			Model:     bike.Model,
			Reg:       bike.Reg,
			Price:     bike.Price,
			CreatedAt: bike.CreatedAt.Nanosecond(),
			UpdatedAt: bike.UpdatedAt.Nanosecond(),
		})
	}

	return &model.BikeQuery{
		Data: bikeModels,
		Meta: &model.Meta{
			Cursor:     input.Cursor,
			TotalItems: count,
		},
	}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
