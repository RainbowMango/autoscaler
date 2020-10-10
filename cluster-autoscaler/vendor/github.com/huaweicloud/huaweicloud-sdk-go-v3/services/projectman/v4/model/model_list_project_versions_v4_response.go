/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectVersionsV4Response struct {
	// 迭代总数
	Total *int32 `json:"total,omitempty"`
	// 迭代信息
	Iterations *[]ListProjectVersionsV4ResponseBodyIterations `json:"iterations,omitempty"`
}

func (o ListProjectVersionsV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectVersionsV4Response", string(data)}, " ")
}
