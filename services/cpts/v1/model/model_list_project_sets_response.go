package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectSetsResponse struct {
	// 状态码

	Code *string `json:"code,omitempty"`
	// 描述

	Message *string `json:"message,omitempty"`
	// 工程集详细信息

	Projects       *[]ProjectsSet `json:"projects,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListProjectSetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectSetsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectSetsResponse", string(data)}, " ")
}