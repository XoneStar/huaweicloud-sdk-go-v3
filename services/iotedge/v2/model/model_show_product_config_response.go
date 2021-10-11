package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProductConfigResponse struct {
	// 服务id

	ServiceId *string `json:"service_id,omitempty"`
	// 服务类型

	ServiceType *string `json:"service_type,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 属性

	Properties     *[]interface{} `json:"properties,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowProductConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProductConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowProductConfigResponse", string(data)}, " ")
}