package github

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client struct {
	v4Client *githubv4.Client
}

func New(ctx context.Context, token string) *Client {
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	var v4Client *githubv4.Client
	if ghHost := os.Getenv("GH_HOST"); ghHost != "" {
		v4Client = githubv4.NewEnterpriseClient(ghHost, httpClient)
	} else {
		v4Client = githubv4.NewClient(httpClient)
	}

	return &Client{
		v4Client: v4Client,
	}
}
