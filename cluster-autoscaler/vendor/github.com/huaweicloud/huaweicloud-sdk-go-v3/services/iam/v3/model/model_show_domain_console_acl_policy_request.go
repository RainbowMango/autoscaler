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

// Request Object
type ShowDomainConsoleAclPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainConsoleAclPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowDomainConsoleAclPolicyRequest", string(data)}, " ")
}
