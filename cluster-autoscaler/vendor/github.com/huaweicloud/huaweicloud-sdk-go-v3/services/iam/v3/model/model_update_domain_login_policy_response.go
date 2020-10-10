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
type UpdateDomainLoginPolicyResponse struct {
	LoginPolicy *LoginPolicyResult `json:"login_policy,omitempty"`
}

func (o UpdateDomainLoginPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDomainLoginPolicyResponse", string(data)}, " ")
}
