package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowConfigurationRequest struct {

	// 配置项键值
	ConfigKey string `json:"config_key"`
}

func (o ShowConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationRequest", string(data)}, " ")
}
