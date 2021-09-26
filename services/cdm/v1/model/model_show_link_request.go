package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLinkRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 连接名称

	LinkName string `json:"link_name"`
}

func (o ShowLinkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowLinkRequest", string(data)}, " ")
}