package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelReadonlyClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelReadonlyClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelReadonlyClusterResponse struct{}"
	}

	return strings.Join([]string{"CancelReadonlyClusterResponse", string(data)}, " ")
}
