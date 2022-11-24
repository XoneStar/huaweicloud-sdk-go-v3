package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeletePtrRecordsRequest struct {
	Body *DeletePtrRecordReq `json:"body,omitempty"`
}

func (o BatchDeletePtrRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePtrRecordsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePtrRecordsRequest", string(data)}, " ")
}