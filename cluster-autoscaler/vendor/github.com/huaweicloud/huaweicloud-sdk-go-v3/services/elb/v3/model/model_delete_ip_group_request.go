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
type DeleteIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`
}

func (o DeleteIpGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteIpGroupRequest", string(data)}, " ")
}
