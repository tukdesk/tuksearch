package api

import (
	"github.com/tukdesk/tuksearch/bleve"
)

type APIService struct {
	global Global
}

func New(global Global) *APIService {
	return &APIService{
		global: global,
	}
}

func (this *APIService) openIndex(indexName string) (bleve.Index, error) {
	kvstore := "ledisdb"
	kvconfig := map[string]interface{}{
		"client": this.global.IndexStoreClient,
		"prefix": indexName,
	}

	return bleve.OpenNonFileIndexUsing(this.global.IndexMetaPath, this.global.DefaultIndexMapping, kvstore, kvconfig)
}
