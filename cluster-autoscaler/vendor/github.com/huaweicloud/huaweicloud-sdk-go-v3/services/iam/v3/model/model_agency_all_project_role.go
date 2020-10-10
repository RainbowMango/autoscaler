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

//
type AgencyAllProjectRole struct {
	// 权限ID。
	Id    string     `json:"id"`
	Links *LinksSelf `json:"links"`
	// 权限名。
	Name string `json:"name"`
}

func (o AgencyAllProjectRole) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AgencyAllProjectRole", string(data)}, " ")
}
