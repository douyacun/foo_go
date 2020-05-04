package auth

import (
	"context"
)

type Auth struct {
	Account     string
	Password string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"account": a.Account, "password": a.Password}, nil
}

func (a *Auth) RequireTransportSecurity() bool  {
	return false
}
