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
type UpdateUserResponse struct {
	User *UpdateUserResult `json:"user,omitempty"`
}

func (o UpdateUserResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}
