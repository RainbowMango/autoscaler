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
type KeystoneUpdateGroupResponse struct {
	Group *KeystoneGroupResultWithLinksSelf `json:"group,omitempty"`
}

func (o KeystoneUpdateGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateGroupResponse", string(data)}, " ")
}
