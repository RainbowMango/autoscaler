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

// Request Object
type CreateCertificateRequest struct {
	Body *CreateCertificateRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateCertificateRequest", string(data)}, " ")
}
