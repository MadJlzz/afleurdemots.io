package context

import (
	"context"
	"github.com/madjlzz/madlens/models"
)

const (
	userKey privateKey = "user"
)

type privateKey string

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	if u, ok := ctx.Value(userKey).(*models.User); ok {
		return u
	}
	return nil
}
