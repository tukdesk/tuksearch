package context

import (
	"reflect"

	"github.com/hprose/hprose-go/hprose"

	"github.com/tukdesk/tuksearch/bleve"
)

func init() {
	hprose.ClassManager.Register(reflect.TypeOf(bleve.SearchResult{}), "SearchResult", "json")
}
