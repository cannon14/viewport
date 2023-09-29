package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	"math/rand"
)

var Auth0 Authenticator

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {

	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+revel.Config.StringDefault("auth0.domain", "")+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     revel.Config.StringDefault("auth0.client.id", ""),
		ClientSecret: revel.Config.StringDefault("auth0.client.secret", ""),
		RedirectURL:  revel.Config.StringDefault("auth0.callback.url", ""),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "scopes"},
	}

	Auth0 = Authenticator{
		Provider: provider,
		Config:   conf,
	}

	return &Auth0, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func GenerateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
