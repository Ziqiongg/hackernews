package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ziqiongg/hackernews/graph/generated"
	"github.com/ziqiongg/hackernews/graph/model"
	"github.com/ziqiongg/hackernews/internal/links"
)

// CreateLink is the resolver for the createLink field.
//have 2 structs for Link in our project, one is use for our graphql server and one is for our database
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Address = input.Address
	link.Title = input.Title
	linkId := link.SaveLink(r.DB)

	return &model.Link{
		ID: strconv.FormatInt(linkId, 10),
		Title: link.Title,
		Address: link.Address,
	},nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user model.NewUser
	user.Password = input.Password
	user.Username = input.Username
	return user.Password,nil

}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshTokenInput is the resolver for the RefreshTokenInput field.
func (r *mutationResolver) RefreshTokenInput(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshTokenInput - RefreshTokenInput"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// var links []*model.Link
	// dummyLink := model.Link{
	// 	Title:   "hello",
	// 	Address: "https://address.org",
	// 	User:    &model.User{Name: "admin"},
	// }
	// links = append(links, &dummyLink)
	// return links, nil

	// var resultLinks []*model.Link
	// var dbLinks []links.Link
	// dbLinks = links.GetAll()
	// for _, link := range dbLinks{
	// 	resultLinks = append(resultLinks, &model.Link{ID:link.ID, Title:link.Title, Address:link.Address})
	// }
	// return resultLinks, nil
	return nil,nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
