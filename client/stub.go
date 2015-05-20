package client

import (
	"github.com/tukdesk/tuksearch/context"
)

type stub struct {
	IndexDoc  func(indexName, docId string, data []byte) error
	DeleteDoc func(indexName, docId string) error
	QueryRaw  func(indexName string, args context.QueryArgs) ([]byte, error) `name:"query"`
}
