package client

import (
	"github.com/tukdesk/tuksearch/context"

	"github.com/tukdesk/tuksearch/bleve"
)

type stub struct {
	IndexDoc  func(indexName, docId string, data []byte) error
	DeleteDoc func(indexName, docId string) error
	Query     func(indexName string, args context.QueryArgs) (*bleve.SearchResult, error)
}
