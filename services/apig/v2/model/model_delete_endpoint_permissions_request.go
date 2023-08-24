package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointPermissionsRequest Request Object
type DeleteEndpointPermissionsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *EndpointPermissionList `json:"body,omitempty"`
}

func (o DeleteEndpointPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointPermissionsRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointPermissionsRequest", string(data)}, " ")
}
