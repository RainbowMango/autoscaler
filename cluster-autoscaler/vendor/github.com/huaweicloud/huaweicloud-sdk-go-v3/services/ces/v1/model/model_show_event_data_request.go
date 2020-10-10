/*
 * CES
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
type ShowEventDataRequest struct {
	ContentType string  `json:"Content-Type"`
	Dim0        string  `json:"dim.0"`
	Dim1        *string `json:"dim.1,omitempty"`
	Dim2        *string `json:"dim.2,omitempty"`
	From        int64   `json:"from"`
	Namespace   string  `json:"namespace"`
	To          int64   `json:"to"`
	Type        string  `json:"type"`
}

func (o ShowEventDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEventDataRequest", string(data)}, " ")
}
