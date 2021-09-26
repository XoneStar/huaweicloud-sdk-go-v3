package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateHostRequestBody `json:"body,omitempty"`
}

func (o CreateHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHostRequest struct{}"
	}

	return strings.Join([]string{"CreateHostRequest", string(data)}, " ")
}