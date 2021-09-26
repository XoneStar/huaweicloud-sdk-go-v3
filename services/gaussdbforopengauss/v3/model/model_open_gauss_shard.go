package model

import (
	"encoding/json"

	"strings"
)

// DN分片扩容时必填
type OpenGaussShard struct {
	// 新增的DN分片扩容数大小

	Count int32 `json:"count"`
}

func (o OpenGaussShard) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussShard struct{}"
	}

	return strings.Join([]string{"OpenGaussShard", string(data)}, " ")
}