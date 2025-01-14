package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InjectionConfig 集群注入信息
type InjectionConfig struct {

	// 是否重启pod
	RestartPod *bool `json:"restartPod,omitempty"`

	Namespaces *Selector `json:"namespaces,omitempty"`
}

func (o InjectionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InjectionConfig struct{}"
	}

	return strings.Join([]string{"InjectionConfig", string(data)}, " ")
}
