package mutation

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/phanletrunghieu/demo/graphql-go/auth"
	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/resolver/model"
)

type signupQueryArgs struct {
	Email    string
	Password string
	Name     string
}

// Signup .
func (m *Mutation) Signup(ctx context.Context, args signupQueryArgs) (*model.AuthPayloadResolver, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        args.Email,
		Name:         args.Name,
		PasswordHash: string(passwordHash),
	}

	err = m.Service.UserService.Create(user)
	if err != nil {
		return nil, err
	}

	// create token
	tokenString, err := auth.GenerateToken(user)

	payload := &domain.AuthPayload{
		Token: tokenString,
		User:  user,
	}

	return model.NewAuthPayloadResolver(ctx, m.Service, payload), err
}
