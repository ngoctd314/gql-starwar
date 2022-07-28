package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ngoctd314/gql-starwar/graph/generated"
	"github.com/ngoctd314/gql-starwar/models"
)

// ID is the resolver for the id field.
func (r *droidResolver) ID(ctx context.Context, obj *models.Droid) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Name is the resolver for the name field.
func (r *droidResolver) Name(ctx context.Context, obj *models.Droid) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Friends is the resolver for the friends field.
func (r *droidResolver) Friends(ctx context.Context, obj *models.Droid) ([]models.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// FriendsConnection is the resolver for the friendsConnection field.
func (r *droidResolver) FriendsConnection(ctx context.Context, obj *models.Droid, first *int, after *string) (*models.FriendsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// AppearsIn is the resolver for the appearsIn field.
func (r *droidResolver) AppearsIn(ctx context.Context, obj *models.Droid) ([]models.Episode, error) {
	panic(fmt.Errorf("not implemented"))
}

// PrimaryFunction is the resolver for the primaryFunction field.
func (r *droidResolver) PrimaryFunction(ctx context.Context, obj *models.Droid) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TotalCount is the resolver for the totalCount field.
func (r *friendsConnectionResolver) TotalCount(ctx context.Context, obj *models.FriendsConnection) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Edges is the resolver for the edges field.
func (r *friendsConnectionResolver) Edges(ctx context.Context, obj *models.FriendsConnection) ([]*models.FriendsEdge, error) {
	panic(fmt.Errorf("not implemented"))
}

// Friends is the resolver for the friends field.
func (r *friendsConnectionResolver) Friends(ctx context.Context, obj *models.FriendsConnection) ([]models.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// PageInfo is the resolver for the pageInfo field.
func (r *friendsConnectionResolver) PageInfo(ctx context.Context, obj *models.FriendsConnection) (*models.PageInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Height is the resolver for the height field.
func (r *humanResolver) Height(ctx context.Context, obj *models.Human, unit *models.LengthUnit) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Friends is the resolver for the friends field.
func (r *humanResolver) Friends(ctx context.Context, obj *models.Human) ([]models.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// FriendsConnection is the resolver for the friendsConnection field.
func (r *humanResolver) FriendsConnection(ctx context.Context, obj *models.Human, first *int, after *string) (*models.FriendsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Starships is the resolver for the starships field.
func (r *humanResolver) Starships(ctx context.Context, obj *models.Human) ([]*models.Starship, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, episode models.Episode, review models.ReviewInput) (*models.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

// Hero is the resolver for the hero field.
func (r *queryResolver) Hero(ctx context.Context, episode *models.Episode) (models.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context, episode models.Episode, since *time.Time) ([]*models.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, text string) ([]models.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (models.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// Droid is the resolver for the droid field.
func (r *queryResolver) Droid(ctx context.Context, id string) (*models.Droid, error) {
	panic(fmt.Errorf("not implemented"))
}

// Human is the resolver for the human field.
func (r *queryResolver) Human(ctx context.Context, id string) (*models.Human, error) {
	panic(fmt.Errorf("not implemented"))
}

// Starship is the resolver for the starship field.
func (r *queryResolver) Starship(ctx context.Context, id string) (*models.Starship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Starts is the resolver for the starts field.
func (r *reviewResolver) Starts(ctx context.Context, obj *models.Review) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Commentary is the resolver for the commentary field.
func (r *reviewResolver) Commentary(ctx context.Context, obj *models.Review) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Time is the resolver for the time field.
func (r *reviewResolver) Time(ctx context.Context, obj *models.Review) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Droid returns generated.DroidResolver implementation.
func (r *Resolver) Droid() generated.DroidResolver { return &droidResolver{r} }

// FriendsConnection returns generated.FriendsConnectionResolver implementation.
func (r *Resolver) FriendsConnection() generated.FriendsConnectionResolver {
	return &friendsConnectionResolver{r}
}

// Human returns generated.HumanResolver implementation.
func (r *Resolver) Human() generated.HumanResolver { return &humanResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

type droidResolver struct{ *Resolver }
type friendsConnectionResolver struct{ *Resolver }
type humanResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
