package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageTagsResponse Response Object
type ListBandwidthPackageTagsResponse struct {

	// 资源ID标识符。
	RequestId string `json:"request_id"`

	// 带宽包的所有标签。
	Tags           []Tag `json:"tags"`
	HttpStatusCode int   `json:"-"`
}

func (o ListBandwidthPackageTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageTagsResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageTagsResponse", string(data)}, " ")
}
