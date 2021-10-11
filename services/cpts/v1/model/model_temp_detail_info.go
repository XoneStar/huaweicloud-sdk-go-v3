package model

import (
	"encoding/json"

	"strings"
)

type TempDetailInfo struct {
	// description

	Description *string `json:"description,omitempty"`
	// id

	Id *int32 `json:"id,omitempty"`
	// is_quoted

	IsQuoted *bool `json:"is_quoted,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// temp_type

	TempType *int32 `json:"temp_type,omitempty"`
	// update_time

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o TempDetailInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TempDetailInfo struct{}"
	}

	return strings.Join([]string{"TempDetailInfo", string(data)}, " ")
}