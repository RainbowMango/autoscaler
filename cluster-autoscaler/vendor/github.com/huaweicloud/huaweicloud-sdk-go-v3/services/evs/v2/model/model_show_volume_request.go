/*
 * EVS
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
type ShowVolumeRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVolumeRequest", string(data)}, " ")
}
