package api

import (
	"fmt"
)

var (
	errIndexNameRequired = fmt.Errorf("index name required")
	errDocIdRequired     = fmt.Errorf("doc id required")
)
