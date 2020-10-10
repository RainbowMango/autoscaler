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
type KeystoneAssociateGroupWithAllProjectPermissionRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o KeystoneAssociateGroupWithAllProjectPermissionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneAssociateGroupWithAllProjectPermissionRequest", string(data)}, " ")
}
