package client

import (
	"encoding/json"

	"github.com/tukdesk/tuksearch/context"

	"github.com/hprose/hprose-go/hprose"
	"github.com/tukdesk/tuksearch/bleve"
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

func (this *Client) QueryResult(indexName string, args context.QueryArgs) (*bleve.SearchResult, error) {
	b, err := this.QueryRaw(indexName, args)
	if err != nil {
		return nil, err
	}

	res := &bleve.SearchResult{}
	return res, json.Unmarshal(b, res)
}
