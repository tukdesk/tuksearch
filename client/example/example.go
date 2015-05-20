package main

import (
	"log"

	"github.com/tukdesk/tuksearch/client"
	"github.com/tukdesk/tuksearch/context"
)

type Doc struct {
	Title    string
	SubField SubField
}

type SubField struct {
	Content string
	Tags    []string
}

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

	doc3 := Doc{}
	doc3.Title = "this is doc 3"
	doc3.SubField.Content = "我们需要一个靠谱的网络供应商"
	doc3.SubField.Tags = []string{"3", "2", "1", "sdk"}

	indexName := "index_test"

	log.Println(c.IndexDoc(indexName, "docId1", doc1))
	log.Println(c.IndexDoc(indexName, "docId2", doc2))
	log.Println(c.IndexDoc(indexName, "docId3", doc3))
	log.Println(c.DeleteDoc(indexName, "docId4"))

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
			Keyword:   "我们",
			Highlight: true,
		},
		context.QueryArgs{
			Keyword:   "sdk",
			Highlight: true,
		},
		context.QueryArgs{
			Keyword:   "ttt",
			Highlight: true,
		},
	}

	for i, query := range queries {
		log.Println(i, query.Keyword)
		res, err := c.Query(indexName, query)
		log.Println(i, "res:", res)
		if err != nil {
			log.Println(i, "err:", err)
		}
	}
}
