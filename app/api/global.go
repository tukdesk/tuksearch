package api

import (
	"github.com/tukdesk/ledisdbcli"

	"github.com/tukdesk/tuksearch/bleve"
)

type Global struct {
	IndexMetaPath       string
	DefaultIndexMapping *bleve.IndexMapping
	IndexStoreClient    *ledisdbcli.Client
}
