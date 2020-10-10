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

// This is a auto create Body Object
type UpdateL7PolicyRequestBody struct {
	L7policy *UpdateL7PolicyOption `json:"l7policy"`
}

func (o UpdateL7PolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateL7PolicyRequestBody", string(data)}, " ")
}
