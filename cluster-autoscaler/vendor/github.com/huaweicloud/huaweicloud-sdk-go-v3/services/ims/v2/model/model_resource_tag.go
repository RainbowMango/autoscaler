/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 标签键值
type ResourceTag struct {
	// 标签的键
	Key string `json:"key"`
	// 标签的值
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
