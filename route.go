package main

import "encoding/json"

type RouteInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		CarrierGatewayID         string      `json:"carrier_gateway_id"`
		DestinationCidrBlock     string      `json:"destination_cidr_block"`
		DestinationIpv6CidrBlock string      `json:"destination_ipv6_cidr_block"`
		DestinationPrefixListID  string      `json:"destination_prefix_list_id"`
		EgressOnlyGatewayID      string      `json:"egress_only_gateway_id"`
		GatewayID                string      `json:"gateway_id"`
		ID                       string      `json:"id"`
		InstanceID               string      `json:"instance_id"`
		InstanceOwnerID          string      `json:"instance_owner_id"`
		LocalGatewayID           string      `json:"local_gateway_id"`
		NatGatewayID             string      `json:"nat_gateway_id"`
		NetworkInterfaceID       string      `json:"network_interface_id"`
		Origin                   string      `json:"origin"`
		RouteTableID             string      `json:"route_table_id"`
		State                    string      `json:"state"`
		Timeouts                 interface{} `json:"timeouts"`
		TransitGatewayID         string      `json:"transit_gateway_id"`
		VpcEndpointID            string      `json:"vpc_endpoint_id"`
		VpcPeeringConnectionID   string      `json:"vpc_peering_connection_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}

func parseRoute(s []interface{}) RouteInstance {
	var si []RouteInstance
	v, _ := json.Marshal(s)
	json.Unmarshal(v, &si)
	if len(si) > 0 {
		return si[0]
	}
	return RouteInstance{}
}
