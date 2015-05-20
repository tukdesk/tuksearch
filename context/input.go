package context

import (
	"reflect"

	"github.com/hprose/hprose-go/hprose"
)

func init() {
	hprose.ClassManager.Register(reflect.TypeOf(QueryArgs{}), "QueryArgs", "json")
}

type QueryArgs struct {
	Keyword   string `json:"keyword"`
	Highlight bool   `json:"highlight"`
}
