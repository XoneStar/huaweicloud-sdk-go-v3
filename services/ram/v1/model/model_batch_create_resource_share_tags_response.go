package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateResourceShareTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateResourceShareTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceShareTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceShareTagsResponse", string(data)}, " ")
}
