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

// 恢复请求body
type BackupRestoreReq struct {
	Restore *BackupRestore `json:"restore"`
}

func (o BackupRestoreReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BackupRestoreReq", string(data)}, " ")
}
