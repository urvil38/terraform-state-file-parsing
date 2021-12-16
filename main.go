package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type tfState struct {
	Version          int        `json:"version"`
	TerraformVersion string     `json:"terraform_version"`
	Serial           int        `json:"serial"`
	Lineage          string     `json:"lineage"`
	Resources        []Resource `json:"resources"`
}

type Resource struct {
	Mode      string        `json:"mode"`
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	Provider  string        `json:"provider"`
	Instances []interface{} `json:"instances"`
}

type tfInstance map[string]interface{}

type tf map[string]tfInstance

type Parser struct {
	state tf
}

func main() {
	filePath := flag.String("file", "terraform.tfstate.backup", "path of terraform state file")

	flag.Parse()

	tfFile := filepath.Base(*filePath)
	if !strings.HasSuffix(tfFile, "tfstate.backup") && !strings.HasSuffix(tfFile, "tfstate") {
		fmt.Println("I don't know about this file format :(")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	var tfParser Parser
	err = tfParser.buildState(b)
	if err != nil {
		log.Fatal(err)
	}

	ec2s := tfParser.listPublicInstances()
	if len(ec2s) > 0 {
		fmt.Printf("I found following EC2 instances which can be access from the Internet !!!!!!!!\n\n\n")
		v, _ := json.MarshalIndent(ec2s, "", "   ")
		fmt.Println(string(v))
	} else {
		fmt.Println("You are safe from the Internet :)")
	}
}

func (p *Parser) buildState(b []byte) error {
	var tfs tfState

	err := json.Unmarshal(b, &tfs)
	if err != nil {
		return err
	}

	tfMap := make(tf)

	for _, r := range tfs.Resources {
		var inter interface{}
		switch r.Type {
		case "aws_subnet":
			inter = parseSubnet(r.Instances)
		case "aws_route":
			inter = parseRoute(r.Instances)
		case "aws_route_table_association":
			inter = parseRouteTableAssoc(r.Instances)
		case "aws_security_group":
			inter = parseSecurityGroup(r.Instances)
		case "aws_instance":
			inter = parseEc2Instance(r.Instances)
		default:
			inter = r.Instances
		}

		_, ok := tfMap[r.Type]
		if !ok {
			tfI := make(tfInstance)
			tfI[r.Name] = inter
			tfMap[r.Type] = tfI
		} else {
			tfMap[r.Type][r.Name] = inter
		}
	}

	p.state = tfMap

	return nil
}

func (p Parser) isPublicSubnet(awsSubId string) bool {
	ass := p.state["aws_route_table_association"]

	routeTables := make(map[string]bool)

	for _, v := range ass {
		assocInstance := v.([]RouteTableAssocInstance)[0]
		if assocInstance.Attributes.SubnetID == awsSubId {
			routeTables[assocInstance.Attributes.RouteTableID] = true
		}
	}

	if len(routeTables) == 0 {
		return false
	}

	tfRoutes := p.state["aws_route"]
	for _, v := range tfRoutes {
		assocInstance := v.([]RouteInstance)[0]
		routeTableId := assocInstance.Attributes.RouteTableID
		gateWayId := assocInstance.Attributes.GatewayID
		destinationAddr := assocInstance.Attributes.DestinationCidrBlock

		if gateWayId != "" && destinationAddr == "0.0.0.0/0" && routeTables[routeTableId] {
			return true
		}
	}

	return false
}

func (p Parser) listPublicInstances() []Ec2Instance {
	tfec2s := p.state["aws_instance"]

	var publicEc2s []Ec2Instance

	for _, v := range tfec2s {
		ec2 := v.([]Ec2Instance)[0]
		subnetId := ec2.Attributes.SubnetID
		publicIP := ec2.Attributes.PublicIP
		var sgs []string

		for _, d := range ec2.Dependencies {
			if strings.HasPrefix(d, "aws_security_group") {
				sgs = append(sgs, strings.TrimPrefix(d, "aws_security_group."))
			}
		}

		if publicIP != "" && p.isPublicSubnet(subnetId) && p.canAccess(sgs, "0.0.0.0", "22") {
			publicEc2s = append(publicEc2s, ec2)
		}
	}

	return publicEc2s
}

func (p Parser) canAccess(tfSgs []string, inter, port string) bool {
	for _, tfSg := range tfSgs {
		tsg := p.state["aws_security_group"][tfSg]
		tsgInstance := tsg.([]SecurityGroupInstance)[0]

		for _, ing := range tsgInstance.Attributes.Ingress {
			if strconv.Itoa(ing.ToPort) == port && existsCIDR(ing.CidrBlocks, inter) {
				return true
			}
		}
	}
	return false
}

func existsCIDR(cidrs []string, cidr string) bool {
	for _, v := range cidrs {
		if strings.Contains(v, cidr) {
			return true
		}
	}
	return false
}
