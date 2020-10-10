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

// Response Object
type ListPoolsResponse struct {
	// 请求ID。 注：自动生成 。
	RequestId *string   `json:"request_id,omitempty"`
	PageInfo  *PageInfo `json:"page_info,omitempty"`
	// 后端服务器组列表。
	Pools *[]Pool `json:"pools,omitempty"`
}

func (o ListPoolsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}
