/*
 * EPS
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
type ShowApiVersionRequest struct {
	ApiVersion string `json:"api_version"`
}

func (o ShowApiVersionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}
