package main

type SecurityGroupInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		Arn         string `json:"arn"`
		Description string `json:"description"`
		Egress      []struct {
			CidrBlocks     []string `json:"cidr_blocks"`
			Description    string   `json:"description"`
			FromPort       int      `json:"from_port"`
			Ipv6CidrBlocks []any    `json:"ipv6_cidr_blocks"`
			PrefixListIds  []any    `json:"prefix_list_ids"`
			Protocol       string   `json:"protocol"`
			SecurityGroups []any    `json:"security_groups"`
			Self           bool     `json:"self"`
			ToPort         int      `json:"to_port"`
		} `json:"egress"`
		ID      string `json:"id"`
		Ingress []struct {
			CidrBlocks     []string `json:"cidr_blocks"`
			Description    string   `json:"description"`
			FromPort       int      `json:"from_port"`
			Ipv6CidrBlocks []any    `json:"ipv6_cidr_blocks"`
			PrefixListIds  []any    `json:"prefix_list_ids"`
			Protocol       string   `json:"protocol"`
			SecurityGroups []any    `json:"security_groups"`
			Self           bool     `json:"self"`
			ToPort         int      `json:"to_port"`
		} `json:"ingress"`
		Name                string `json:"name"`
		NamePrefix          string `json:"name_prefix"`
		OwnerID             string `json:"owner_id"`
		RevokeRulesOnDelete bool   `json:"revoke_rules_on_delete"`
		Tags                any    `json:"tags"`
		Timeouts            any    `json:"timeouts"`
		VpcID               string `json:"vpc_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}
