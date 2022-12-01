package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBatchJobRequest struct {

	// 批量处理作业类型
	JobType *string `json:"job_type,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果排序，如按照创建时间降序排序为created_at:desc，升序排序为created_at:asc
	Sort *string `json:"sort,omitempty"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ListBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchJobRequest struct{}"
	}

	return strings.Join([]string{"ListBatchJobRequest", string(data)}, " ")
}
