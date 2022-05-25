package types

type SubnetInstance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		Arn                         string `json:"arn"`
		AssignIpv6AddressOnCreation bool   `json:"assign_ipv6_address_on_creation"`
		AvailabilityZone            string `json:"availability_zone"`
		AvailabilityZoneID          string `json:"availability_zone_id"`
		CidrBlock                   string `json:"cidr_block"`
		CustomerOwnedIpv4Pool       string `json:"customer_owned_ipv4_pool"`
		ID                          string `json:"id"`
		Ipv6CidrBlock               string `json:"ipv6_cidr_block"`
		Ipv6CidrBlockAssociationID  string `json:"ipv6_cidr_block_association_id"`
		MapCustomerOwnedIPOnLaunch  bool   `json:"map_customer_owned_ip_on_launch"`
		MapPublicIPOnLaunch         bool   `json:"map_public_ip_on_launch"`
		OutpostArn                  string `json:"outpost_arn"`
		OwnerID                     string `json:"owner_id"`
		Tags                        struct {
			Name string `json:"Name"`
		} `json:"tags"`
		TagsAll struct {
			Name string `json:"Name"`
		} `json:"tags_all"`
		Timeouts any    `json:"timeouts"`
		VpcID    string `json:"vpc_id"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}
