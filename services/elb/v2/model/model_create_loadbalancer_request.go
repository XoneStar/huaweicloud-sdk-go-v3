/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLoadbalancerRequest struct {
	Body *CreateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateLoadbalancerRequest", string(data)}, " ")
}