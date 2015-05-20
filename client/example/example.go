package main

import (
	"encoding/json"
	"log"

	"github.com/tukdesk/tuksearch/client"
	"github.com/tukdesk/tuksearch/context"
)

func main() {
	c, err := client.New("http://127.0.0.1:56380")
	if err != nil {
		log.Fatalln(err)
	}

	doc1 := map[string]interface{}{
		"title": "doc 1 title",
		"tags":  []string{"a", "b", "c", "我们的爱"},
	}

	doc2 := map[string]interface{}{
		"title": "title for doc 2",
		"tags":  []string{"d", "b", "爱情保卫战", "e"},
	}

	indexName := "index_test"

	b1, _ := json.Marshal(doc1)
	b2, _ := json.Marshal(doc2)

	log.Println(c.IndexDoc(indexName, "docId1", b1))
	log.Println(c.IndexDoc(indexName, "docId2", b2))
	log.Println(c.DeleteDoc(indexName, "docId3"))

	queries := []context.QueryArgs{
		context.QueryArgs{
			Keyword:   "b",
			Highlight: true,
		},
		context.QueryArgs{
			Keyword:   "爱",
			Highlight: true,
		},
		context.QueryArgs{
			Keyword:   "保卫",
			Highlight: true,
		},
		context.QueryArgs{
			Keyword:   "ttt",
			Highlight: true,
		},
	}

	for i, query := range queries {
		log.Println(i, query.Keyword)
		res, err := c.QueryResult(indexName, query)
		log.Println(i, "err:", err)
		log.Println(i, "res:", res)
	}
}
