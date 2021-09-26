package model

import (
	"encoding/json"

	"strings"
)

// CreateTaskRequestBody
type CreateTaskRequestBody struct {
	// name

	Name string `json:"name"`
	// project_id

	ProjectId int32 `json:"project_id"`
	// temps

	Temps *[]string `json:"temps,omitempty"`
	// operate_mode

	OperateMode *int32 `json:"operate_mode,omitempty"`
	// bench_concurrent

	BenchConcurrent *int32 `json:"bench_concurrent,omitempty"`
}

func (o CreateTaskRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestBody", string(data)}, " ")
}