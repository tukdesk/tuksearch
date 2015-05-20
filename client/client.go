package client

import (
	"github.com/hprose/hprose-go/hprose"
)

type Client struct {
	*stub
}

func New(url string) (*Client, error) {
	client := hprose.NewHttpClient(url)
	s := &stub{}
	client.(*hprose.HttpClient).DebugEnabled = true
	client.UseService(s)
	return &Client{
		stub: s,
	}, nil
}
