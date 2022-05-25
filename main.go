package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/urvil38/terraform-state-file-parsing/types"
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
	Instances []any `json:"instances"`
}

type tfInstance map[string]any

type tf map[string]tfInstance

type Parser struct {
	state tf
}

func main() {

	var verbose bool
	filePath := flag.String("file", "terraform.tfstate.backup", "path of terraform state file")
	flag.BoolVar(&verbose, "verbose", false, "provide detailed output")
	flag.BoolVar(&verbose, "vv", false, "provide detailed output")

	flag.Parse()

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

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
		if verbose {
			v, _ := json.MarshalIndent(ec2s, "", "   ")
			fmt.Println(string(v))
		} else {
			var pis []types.PublicInstance
			for _, ec2 := range ec2s {
				pis = append(pis, types.PublicInstance{
					Id:        ec2.Attributes.ID,
					PrivateIP: ec2.Attributes.PrivateIP,
					PublicIP:  ec2.Attributes.PublicIP,
				})
			}
			v, _ := json.MarshalIndent(pis, "", "   ")
			fmt.Println(string(v))
		}
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
		var inter any
		switch r.Type {
		case "aws_subnet":
			inter = parseInstance[types.SubnetInstance](r.Instances)
		case "aws_route":
			inter = parseInstance[types.RouteInstance](r.Instances)
		case "aws_route_table_association":
			inter = parseInstance[types.RouteTableAssocInstance](r.Instances)
		case "aws_security_group":
			inter = parseInstance[types.SecurityGroupInstance](r.Instances)
		case "aws_instance":
			inter = parseInstance[types.Ec2Instance](r.Instances)
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
		assocInstance := v.(types.RouteTableAssocInstance)
		if assocInstance.Attributes.SubnetID == awsSubId {
			routeTables[assocInstance.Attributes.RouteTableID] = true
		}
	}

	if len(routeTables) == 0 {
		return false
	}

	tfRoutes := p.state["aws_route"]
	for _, v := range tfRoutes {
		assocInstance := v.(types.RouteInstance)
		routeTableId := assocInstance.Attributes.RouteTableID
		gateWayId := assocInstance.Attributes.GatewayID
		destinationAddr := assocInstance.Attributes.DestinationCidrBlock

		if gateWayId != "" && destinationAddr == "0.0.0.0/0" && routeTables[routeTableId] {
			return true
		}
	}

	return false
}

func (p Parser) listPublicInstances() []types.Ec2Instance {
	tfec2s := p.state["aws_instance"]

	var publicEc2s []types.Ec2Instance

	for _, v := range tfec2s {
		ec2 := v.(types.Ec2Instance)
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
		tsgInstance := tsg.(types.SecurityGroupInstance)

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
