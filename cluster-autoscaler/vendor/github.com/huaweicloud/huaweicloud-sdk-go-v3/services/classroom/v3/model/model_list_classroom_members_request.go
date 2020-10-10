/*
 * Classroom
 *
 * devcloud classedge api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClassroomMembersRequest struct {
	ClassroomId string  `json:"classroom_id"`
	Offset      *int32  `json:"offset,omitempty"`
	Limit       *int32  `json:"limit,omitempty"`
	Filter      *string `json:"filter,omitempty"`
}

func (o ListClassroomMembersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListClassroomMembersRequest", string(data)}, " ")
}
