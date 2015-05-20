package context

import (
	"encoding/json"
)

var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
)
