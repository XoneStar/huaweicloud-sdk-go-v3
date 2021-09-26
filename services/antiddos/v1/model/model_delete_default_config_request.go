package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDefaultConfigRequest struct {
}

func (o DeleteDefaultConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteDefaultConfigRequest", string(data)}, " ")
}