package verniy

import (
	"net/http"
	"time"

	"github.com/rl404/verniy/internal/limiter"
)

// Client is anilist client.
type Client struct {
	host    string
	http    http.Client
	limiter limiter.Limiter
}

// New to create new anilist client.
func New() *Client {
	return &Client{
		host: "https://graphql.anilist.co",
		http: http.Client{
			Timeout: 10 * time.Second,
		},
		limiter: limiter.New(90, time.Minute),
	}
}
