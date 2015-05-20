package client

import (
	"github.com/tukdesk/tuksearch/bleve"
	"github.com/tukdesk/tuksearch/context"

	"github.com/hprose/hprose-go/hprose"
)

type Client struct {
	stub *stub
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

func (this *Client) IndexDoc(indexName, docId string, doc interface{}) error {
	b, err := context.Marshal(doc)
	if err != nil {
		return err
	}
	return this.stub.IndexDoc(indexName, docId, b)
}

func (this *Client) DeleteDoc(indexName, docId string) error {
	return this.stub.DeleteDoc(indexName, docId)
}

func (this *Client) Query(indexName string, args context.QueryArgs) (*bleve.SearchResult, error) {
	return this.stub.Query(indexName, args)
}
