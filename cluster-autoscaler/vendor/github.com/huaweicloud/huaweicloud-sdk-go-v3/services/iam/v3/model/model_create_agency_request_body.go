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
type CreateAgencyRequestBody struct {
	Agency *CreateAgencyOption `json:"agency"`
}

func (o CreateAgencyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAgencyRequestBody", string(data)}, " ")
}
