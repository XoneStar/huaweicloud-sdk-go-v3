package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetGaussMySqlQuotasRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetQuotasRequestBody `json:"body,omitempty"`
}

func (o SetGaussMySqlQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetGaussMySqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"SetGaussMySqlQuotasRequest", string(data)}, " ")
}