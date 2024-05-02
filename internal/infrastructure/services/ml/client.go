package ml

import "fmt"

type Client struct {
	Host string
}

func NewClient(host string) (*Client, error) {
	if host == "" {
		return nil, fmt.Errorf("host is nil") //nolint:goerr113
	}

	return &Client{
		Host: host,
	}, nil
}
