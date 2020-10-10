/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSubNetworkInterfaceResponse struct {
	// 请求ID
	RequestId           *string              `json:"request_id,omitempty"`
	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
}

func (o ShowSubNetworkInterfaceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSubNetworkInterfaceResponse", string(data)}, " ")
}
