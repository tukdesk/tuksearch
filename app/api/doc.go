package api

import (
	"encoding/json"
)

func (this *APIService) IndexDoc(indexName, docId string, data []byte) error {
	if indexName == "" {
		return errIndexNameRequired
	}

	if docId == "" {
		return errDocIdRequired
	}

	index, err := this.openIndex(indexName)
	if err != nil {
		return err
	}

	i, err := docFromBytes(data)
	if err != nil {
		return err
	}

	if err := index.Index(docId, i); err != nil {
		return err
	}

	return nil
}

func (this *APIService) DeleteDoc(indexName, docId string) error {
	if indexName == "" {
		return errIndexNameRequired
	}

	if docId == "" {
		return errDocIdRequired
	}

	index, err := this.openIndex(indexName)
	if err != nil {
		return err
	}

	if err := index.Delete(docId); err != nil {
		return err
	}

	return nil

}

func docFromBytes(data []byte) (interface{}, error) {
	var i interface{}
	return i, json.Unmarshal(data, &i)
}
