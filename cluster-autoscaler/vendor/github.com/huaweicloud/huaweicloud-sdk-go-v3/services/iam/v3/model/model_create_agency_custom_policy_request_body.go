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
type CreateAgencyCustomPolicyRequestBody struct {
	Role *AgencyPolicyRoleOption `json:"role"`
}

func (o CreateAgencyCustomPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAgencyCustomPolicyRequestBody", string(data)}, " ")
}
