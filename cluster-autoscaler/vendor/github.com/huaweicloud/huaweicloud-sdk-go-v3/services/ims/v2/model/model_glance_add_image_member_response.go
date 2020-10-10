/*
 * IMS
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
type GlanceAddImageMemberResponse struct {
	// 共享状态
	Status *string `json:"status,omitempty"`
	// 共享时间，格式为UTC时间
	CreatedAt *string `json:"created_at,omitempty"`
	// 更新时间，格式为UTC时间
	UpdatedAt *string `json:"updated_at,omitempty"`
	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`
	// 成员ID
	MemberId *string `json:"member_id,omitempty"`
	// 共享视图
	Schema *string `json:"schema,omitempty"`
}

func (o GlanceAddImageMemberResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GlanceAddImageMemberResponse", string(data)}, " ")
}
