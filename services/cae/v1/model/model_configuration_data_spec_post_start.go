package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationDataSpecPostStart - type为\"lifecycle\"时，配置此参数。 - 参数含义：生命周期管理组件配置启动后处理。
type ConfigurationDataSpecPostStart struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationDataSpecPostStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataSpecPostStart struct{}"
	}

	return strings.Join([]string{"ConfigurationDataSpecPostStart", string(data)}, " ")
}
