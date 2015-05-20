package api

import (
	"encoding/json"

	"github.com/tukdesk/tuksearch/context"

	"github.com/tukdesk/tuksearch/bleve"
)

func (this *APIService) Query(indexName string, args context.QueryArgs) ([]byte, error) {
	if indexName == "" {
		return nil, errIndexNameRequired
	}
	index, err := this.openIndex(indexName)
	if err != nil {
		return nil, err
	}

	query := bleve.NewMatchQuery(args.Keyword)
	search := bleve.NewSearchRequest(query)
	if args.Highlight {
		search.Highlight = bleve.NewHighlight()
	}

	res, err := index.Search(search)
	if err != nil {
		return nil, err
	}

	return json.Marshal(res)
}
