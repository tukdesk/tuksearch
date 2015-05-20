package api

import (
	"github.com/tukdesk/tuksearch/context"

	"github.com/tukdesk/tuksearch/bleve"
)

func (this *APIService) Query(indexName string, args context.QueryArgs) (*bleve.SearchResult, error) {
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

	return index.Search(search)
}
