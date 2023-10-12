package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiAttachablePluginsRequest Request Object
type ListApiAttachablePluginsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// API编号
	ApiId string `json:"api_id"`

	// 发布的环境编号
	EnvId string `json:"env_id"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型
	PluginType *string `json:"plugin_type,omitempty"`

	// 插件编号
	PluginId *string `json:"plugin_id,omitempty"`

	// 集成应用编号
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 集成应用名称
	RomaAppName *string `json:"roma_app_name,omitempty"`
}

func (o ListApiAttachablePluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiAttachablePluginsRequest struct{}"
	}

	return strings.Join([]string{"ListApiAttachablePluginsRequest", string(data)}, " ")
}
