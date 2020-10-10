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

//
type AvailabilityZone struct {
	// 可用区code
	Code string `json:"code"`
	// az状态.取值：  ACTIVE
	State string `json:"state"`
}

func (o AvailabilityZone) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
