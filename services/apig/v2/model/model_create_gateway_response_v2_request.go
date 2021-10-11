package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateGatewayResponseV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`

	Body *ResponsesCreate `json:"body,omitempty"`
}

func (o CreateGatewayResponseV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"CreateGatewayResponseV2Request", string(data)}, " ")
}