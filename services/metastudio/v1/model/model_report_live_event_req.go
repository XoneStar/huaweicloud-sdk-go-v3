package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportLiveEventReq struct {

	// 事件条目数。
	Total int32 `json:"total"`

	// 事件内容。
	Events *[]LiveEvent `json:"events,omitempty"`
}

func (o ReportLiveEventReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportLiveEventReq struct{}"
	}

	return strings.Join([]string{"ReportLiveEventReq", string(data)}, " ")
}
