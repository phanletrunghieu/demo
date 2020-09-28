package resolver

import (
	"context"

	"github.com/phanletrunghieu/demo/gqlgen/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (prisma.User, error) {
	author, err := r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).Author().Exec(ctx)
	return *author, err
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Posts(nil).Exec(ctx)
}
