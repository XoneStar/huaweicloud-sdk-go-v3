package model

import (
	"encoding/json"

	"strings"
)

type BackendLatencyStats struct {
	// 最大后端延时

	MaxBackendLatency *int32 `json:"max_backend_latency,omitempty"`
	// 平均后端延时

	AvgBackendLatency float32 `json:"avg_backend_latency,omitempty"`
}

func (o BackendLatencyStats) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BackendLatencyStats struct{}"
	}

	return strings.Join([]string{"BackendLatencyStats", string(data)}, " ")
}