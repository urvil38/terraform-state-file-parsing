package main

import (
	"encoding/json"
)

type SecurityGroupInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		Arn         string `json:"arn"`
		Description string `json:"description"`
		Egress      []struct {
			CidrBlocks     []string      `json:"cidr_blocks"`
			Description    string        `json:"description"`
			FromPort       int           `json:"from_port"`
			Ipv6CidrBlocks []interface{} `json:"ipv6_cidr_blocks"`
			PrefixListIds  []interface{} `json:"prefix_list_ids"`
			Protocol       string        `json:"protocol"`
			SecurityGroups []interface{} `json:"security_groups"`
			Self           bool          `json:"self"`
			ToPort         int           `json:"to_port"`
		} `json:"egress"`
		ID      string `json:"id"`
		Ingress []struct {
			CidrBlocks     []string      `json:"cidr_blocks"`
			Description    string        `json:"description"`
			FromPort       int           `json:"from_port"`
			Ipv6CidrBlocks []interface{} `json:"ipv6_cidr_blocks"`
			PrefixListIds  []interface{} `json:"prefix_list_ids"`
			Protocol       string        `json:"protocol"`
			SecurityGroups []interface{} `json:"security_groups"`
			Self           bool          `json:"self"`
			ToPort         int           `json:"to_port"`
		} `json:"ingress"`
		Name                string      `json:"name"`
		NamePrefix          string      `json:"name_prefix"`
		OwnerID             string      `json:"owner_id"`
		RevokeRulesOnDelete bool        `json:"revoke_rules_on_delete"`
		Tags                interface{} `json:"tags"`
		Timeouts            interface{} `json:"timeouts"`
		VpcID               string      `json:"vpc_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}

func parseSecurityGroup(sg []interface{}) []SecurityGroupInstance {
	var sgi []SecurityGroupInstance
	v, _ := json.Marshal(sg)
	json.Unmarshal(v, &sgi)
	return sgi
}
