package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var regionPtr string
var envPtr string

// initializer
func init() {
	flag.StringVar(&regionPtr, "region", "us-west-1", "which region")
	flag.StringVar(&envPtr, "environment", "", "which environment")
	flag.Parse()
}

// by environment
func getByEnv() {
	svc := ec2.New(session.New(), &aws.Config{
		Region: &regionPtr,
	})

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String(envPtr),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	CheckErr(err)

	fmt.Println("Total number of instances: ", len(resp.Reservations))
	fmt.Println("")

	for idx, res := range resp.Reservations {
		fmt.Println("Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {

			var Name string
			var Env string

			for _, tag := range inst.Tags {
				if *tag.Key == "Role" {
					Name = *tag.Value
				}
				if *tag.Key == "Environment" {
					Env = *tag.Value
				}
			}

			fmt.Println("Name: ", Name)
			fmt.Println("Environment: ", Env)
			fmt.Println("ID: ", *inst.InstanceId)
			fmt.Println("Private IP: ", *inst.PrivateIpAddress)
			fmt.Println("State: ", *inst.State.Name)
			fmt.Println("Launched: ", *inst.LaunchTime)
			fmt.Println("Type: ", *inst.InstanceType)

		}

		fmt.Println("")

	}
}

// get all envs
func getAllEnv() {
	svc := ec2.New(session.New(), &aws.Config{
		Region: &regionPtr,
	})

	resp, err := svc.DescribeInstances(nil)
	CheckErr(err)

	fmt.Println("Total number of instances: ", len(resp.Reservations))
	fmt.Println("")

	for idx, res := range resp.Reservations {
		fmt.Println("Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {

			var Name string
			var Env string

			for _, tag := range inst.Tags {
				if *tag.Key == "Role" {
					Name = *tag.Value
				}
				if *tag.Key == "Environment" {
					Env = *tag.Value
				}
			}

			fmt.Println("Name: ", Name)
			fmt.Println("Environment: ", Env)
			fmt.Println("ID: ", *inst.InstanceId)
			fmt.Println("Private IP: ", *inst.PrivateIpAddress)
			fmt.Println("State: ", *inst.State.Name)
			fmt.Println("Launched: ", *inst.LaunchTime)
			fmt.Println("Type: ", *inst.InstanceType)

		}

		fmt.Println("")

	}
}

// main function
func main() {
	if envPtr == "" {
		getAllEnv()
	} else {
		getByEnv()
	}
}
