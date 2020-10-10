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
type ShowDomainLoginPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainLoginPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowDomainLoginPolicyRequest", string(data)}, " ")
}
