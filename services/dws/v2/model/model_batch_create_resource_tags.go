package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签列表信息
type BatchCreateResourceTags struct {

	// 标签列表
	Tags []BatchCreateResourceTag `json:"tags"`
}

func (o BatchCreateResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTags struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTags", string(data)}, " ")
}
