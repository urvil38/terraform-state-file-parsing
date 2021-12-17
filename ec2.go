package main

import "encoding/json"

type Ec2Instance struct {
	SchemaVersion int `json:"schema_version"`
	Attributes    struct {
		Ami                      string `json:"ami"`
		Arn                      string `json:"arn"`
		AssociatePublicIPAddress bool   `json:"associate_public_ip_address"`
		AvailabilityZone         string `json:"availability_zone"`
		CPUCoreCount             int    `json:"cpu_core_count"`
		CPUThreadsPerCore        int    `json:"cpu_threads_per_core"`
		CreditSpecification      []struct {
			CPUCredits string `json:"cpu_credits"`
		} `json:"credit_specification"`
		DisableAPITermination bool          `json:"disable_api_termination"`
		EbsBlockDevice        []interface{} `json:"ebs_block_device"`
		EbsOptimized          bool          `json:"ebs_optimized"`
		EnclaveOptions        []struct {
			Enabled bool `json:"enabled"`
		} `json:"enclave_options"`
		EphemeralBlockDevice              []interface{} `json:"ephemeral_block_device"`
		GetPasswordData                   bool          `json:"get_password_data"`
		Hibernation                       bool          `json:"hibernation"`
		HostID                            interface{}   `json:"host_id"`
		IamInstanceProfile                string        `json:"iam_instance_profile"`
		ID                                string        `json:"id"`
		InstanceInitiatedShutdownBehavior interface{}   `json:"instance_initiated_shutdown_behavior"`
		InstanceState                     string        `json:"instance_state"`
		InstanceType                      string        `json:"instance_type"`
		Ipv6AddressCount                  int           `json:"ipv6_address_count"`
		Ipv6Addresses                     []interface{} `json:"ipv6_addresses"`
		KeyName                           string        `json:"key_name"`
		MetadataOptions                   []struct {
			HTTPEndpoint            string `json:"http_endpoint"`
			HTTPPutResponseHopLimit int    `json:"http_put_response_hop_limit"`
			HTTPTokens              string `json:"http_tokens"`
		} `json:"metadata_options"`
		Monitoring                bool          `json:"monitoring"`
		NetworkInterface          []interface{} `json:"network_interface"`
		OutpostArn                string        `json:"outpost_arn"`
		PasswordData              string        `json:"password_data"`
		PlacementGroup            string        `json:"placement_group"`
		PrimaryNetworkInterfaceID string        `json:"primary_network_interface_id"`
		PrivateDNS                string        `json:"private_dns"`
		PrivateIP                 string        `json:"private_ip"`
		PublicDNS                 string        `json:"public_dns"`
		PublicIP                  string        `json:"public_ip"`
		RootBlockDevice           []struct {
			DeleteOnTermination bool   `json:"delete_on_termination"`
			DeviceName          string `json:"device_name"`
			Encrypted           bool   `json:"encrypted"`
			Iops                int    `json:"iops"`
			KmsKeyID            string `json:"kms_key_id"`
			Tags                struct {
			} `json:"tags"`
			Throughput int    `json:"throughput"`
			VolumeID   string `json:"volume_id"`
			VolumeSize int    `json:"volume_size"`
			VolumeType string `json:"volume_type"`
		} `json:"root_block_device"`
		SecondaryPrivateIps []interface{} `json:"secondary_private_ips"`
		SecurityGroups      []interface{} `json:"security_groups"`
		SourceDestCheck     bool          `json:"source_dest_check"`
		SubnetID            string        `json:"subnet_id"`
		Tags                struct {
			Name string `json:"Name"`
		} `json:"tags"`
		Tenancy             string      `json:"tenancy"`
		Timeouts            interface{} `json:"timeouts"`
		UserData            interface{} `json:"user_data"`
		UserDataBase64      interface{} `json:"user_data_base64"`
		VolumeTags          interface{} `json:"volume_tags"`
		VpcSecurityGroupIds []string    `json:"vpc_security_group_ids"`
	} `json:"attributes"`
	Private      string   `json:"private"`
	Dependencies []string `json:"dependencies"`
}

func parseEc2Instance(s []interface{}) Ec2Instance {
	var si []Ec2Instance
	v, _ := json.Marshal(s)
	json.Unmarshal(v, &si)
	if len(si) > 0 {
		return si[0]
	}
	return Ec2Instance{}
}
