package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
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

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var (
		account, password string
	)
	if val, ok := md["account"]; ok {
		account = val[0]
	}
	if val, ok := md["password"]; !ok {
		password = val[0]
	}
	fmt.Println(account, password)
	return nil
}