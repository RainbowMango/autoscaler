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
type DeleteVpcRouteResponse struct {
}

func (o DeleteVpcRouteResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcRouteResponse", string(data)}, " ")
}
