package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTemplateRequest struct {
	// 需要查询的模板信息的id

	Id string `json:"id"`
}

func (o ShowTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}