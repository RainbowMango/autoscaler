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

// Request Object
type ShowVaultRequest struct {
	VaultId string `json:"vault_id"`
}

func (o ShowVaultRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVaultRequest", string(data)}, " ")
}
