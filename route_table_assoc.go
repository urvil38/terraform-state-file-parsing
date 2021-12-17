package main

import "encoding/json"

type RouteTableAssocInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		GatewayID    interface{} `json:"gateway_id"`
		ID           string      `json:"id"`
		RouteTableID string      `json:"route_table_id"`
		SubnetID     string      `json:"subnet_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}

func parseRouteTableAssoc(s []interface{}) RouteTableAssocInstance {
	var si []RouteTableAssocInstance
	v, _ := json.Marshal(s)
	json.Unmarshal(v, &si)
	if len(si) > 0 {
		return si[0]
	}
	return RouteTableAssocInstance{}
}
