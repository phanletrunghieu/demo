package mutation

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/phanletrunghieu/demo/graphql-go/auth"
	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/resolver/model"
)

type loginQueryArgs struct {
	Email    string
	Password string
}

// Login .
func (m *Mutation) Login(ctx context.Context, args loginQueryArgs) (*model.AuthPayloadResolver, error) {
	user, err := m.Service.UserService.FindByEmail(args.Email)
	if err != nil {
		return nil, err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(args.Password))
	if err != nil {
		return nil, errors.New("Wrong password")
	}

	// create token
	tokenString, err := auth.GenerateToken(user)

	payload := &domain.AuthPayload{
		Token: tokenString,
		User:  user,
	}

	return model.NewAuthPayloadResolver(ctx, m.Service, payload), nil
}
