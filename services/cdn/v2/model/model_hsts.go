package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Hsts HSTS：配置HSTS后，将强制客户端（如浏览器）使用 HTTPS 协议访问服务器，提升访问安全性。
type Hsts struct {

	// 状态，on：打开，off：关闭。
	Status string `json:"status"`

	// 过期时间,即：响应头“Strict-Transport-Security”在客户端的缓存时间。单位:秒,取值范围:0~63072000。  > status参数为on时，必传。
	MaxAge *int32 `json:"max_age,omitempty"`

	// 包含子域名，on：包含，off：不包含。   > status参数为on时，必传。
	IncludeSubdomains *string `json:"include_subdomains,omitempty"`
}

func (o Hsts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hsts struct{}"
	}

	return strings.Join([]string{"Hsts", string(data)}, " ")
}
