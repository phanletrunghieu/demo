package resolver

import (
	"context"

	"github.com/phanletrunghieu/demo/gqlgen/prisma-client"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Feed(ctx context.Context) ([]prisma.Post, error) {
	published := true
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{Published: &published},
	}).Exec(ctx)
}
func (r *queryResolver) FilterPosts(ctx context.Context, searchString *string) ([]prisma.Post, error) {
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{
			Or: []prisma.PostWhereInput{
				prisma.PostWhereInput{
					TitleContains: searchString,
				},
				prisma.PostWhereInput{
					TitleContains: searchString,
				},
			},
		},
	}).Exec(ctx)
}
func (r *queryResolver) Post(ctx context.Context, id string) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &id}).Exec(ctx)
}
