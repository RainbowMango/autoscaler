/*
 * Cbr
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
type DeleteVaultResponse struct {
}

func (o DeleteVaultResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVaultResponse", string(data)}, " ")
}
