package shared

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kks-learning-management-api/shared/failure"
)

func SliceStringToInterfaces(slices []string) []interface{} {
	results := []interface{}{}
	for _, s := range slices {
		results = append(results, s)
	}
	return results
}

const (
	IdKey    = "id"
	RoleKey  = "role"
	TokenKey = "token"
)

// GetIdFromContext retrieves the id from the request context.
func GetIdFromContext(r *http.Request) (string, error) {
	username, ok := r.Context().Value(IdKey).(string)
	if !ok {
		return "", failure.BadRequestFromString(fmt.Sprintf("`%s` not found in context", IdKey))
	}
	return username, nil
}

func GetRoleFromContext(r *http.Request) (string, error) {
	role, ok := r.Context().Value(RoleKey).(string)
	if !ok {
		return "", failure.BadRequestFromString(fmt.Sprintf("`%s` not found in context", RoleKey))
	}
	return role, nil
}

func GetTokenFromContext(r *http.Request) (string, error) {
	token, ok := r.Context().Value(TokenKey).(string)
	if !ok {
		return "", failure.BadRequestFromString(fmt.Sprintf("`%s` not found in context", TokenKey))
	}
	return token, nil
}

// WithId adds the id to the context.
func WithId(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, IdKey, username)
}

func WithRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, RoleKey, role)
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, TokenKey, token)
}
