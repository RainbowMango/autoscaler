/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListRegionsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 区域信息列表。
	Regions *[]Region `json:"regions,omitempty"`
}

func (o KeystoneListRegionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListRegionsResponse", string(data)}, " ")
}
