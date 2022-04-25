package main

type RouteTableAssocInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		GatewayID    any    `json:"gateway_id"`
		ID           string `json:"id"`
		RouteTableID string `json:"route_table_id"`
		SubnetID     string `json:"subnet_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}
