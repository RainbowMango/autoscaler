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
type UpdateIpGroupOption struct {
	// IP地址组的描述信息
	Description *string `json:"description,omitempty"`
	// IP地址组的名称
	Name *string `json:"name,omitempty"`
	// IP地址组中包含的ip列表。
	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`
}

func (o UpdateIpGroupOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIpGroupOption", string(data)}, " ")
}
