package arweave

import (
	"fmt"
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Arweave struct {
	client Client
	server string
	port   string
}

type Option func(*Arweave) *Arweave

func New(opts ...Option) *Arweave {
	arweave := &Arweave{
		client: &http.Client{},
		server: "arweave.net",
	}

	for _, opt := range opts {
		opt(arweave)
	}

	return arweave
}

func (a *Arweave) fqdn() string {
	if a.port == "" {
		return fmt.Sprintf("https://%s", a.server)
	}
	return fmt.Sprintf("https://%s:%s", a.server, a.port)
}
