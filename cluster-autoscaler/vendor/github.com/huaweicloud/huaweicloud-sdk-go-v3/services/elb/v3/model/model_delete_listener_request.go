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
type DeleteListenerRequest struct {
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}
