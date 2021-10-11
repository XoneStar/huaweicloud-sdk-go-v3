package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApiAclBindingV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 绑定关系编号

	AclBindingsId string `json:"acl_bindings_id"`
}

func (o DeleteApiAclBindingV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiAclBindingV2Request", string(data)}, " ")
}