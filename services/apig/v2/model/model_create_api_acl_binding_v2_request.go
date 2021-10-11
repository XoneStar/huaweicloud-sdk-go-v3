package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateApiAclBindingV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *AclApiBindingCreate `json:"body,omitempty"`
}

func (o CreateApiAclBindingV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiAclBindingV2Request", string(data)}, " ")
}