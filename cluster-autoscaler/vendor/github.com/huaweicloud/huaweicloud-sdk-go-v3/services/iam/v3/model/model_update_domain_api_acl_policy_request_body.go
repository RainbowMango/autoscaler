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
type UpdateDomainApiAclPolicyRequestBody struct {
	ApiAclPolicy *AclPolicyOption `json:"api_acl_policy"`
}

func (o UpdateDomainApiAclPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDomainApiAclPolicyRequestBody", string(data)}, " ")
}
