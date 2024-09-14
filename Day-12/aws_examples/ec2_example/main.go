package main

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 service client
	svc := ec2.NewFromConfig(cfg)

	// Pull the latest Ubuntu AMI
	amiID, err := getLatestUbuntuAMI(svc)
	if err != nil {
		log.Fatalf("failed to get latest Ubuntu AMI, %v", err)
	}
	fmt.Printf("Latest Ubuntu AMI ID: %s\n", amiID)

	// Call the create instance function
	instanceID, err := createEC2Instance(svc, amiID)
	if err != nil {
		log.Fatalf("failed to create instance, %v", err)
	}
	fmt.Printf("Created instance with ID: %s\n", instanceID)

	// List all EC2 instances
	err = listEC2Instances(svc)
	if err != nil {
		log.Fatalf("failed to list instances, %v", err)
	}

	// Delete the EC2 instance
	err = deleteEC2Instance(svc, instanceID)
	if err != nil {
		log.Fatalf("failed to delete instance, %v", err)
	}
	fmt.Printf("Deleted instance with ID: %s\n", instanceID)
}

// getLatestUbuntuAMI pulls the latest Ubuntu AMI from AWS
func getLatestUbuntuAMI(svc *ec2.Client) (string, error) {
	input := &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("architecture"),
				Values: []string{"x86_64"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
			{
				Name:   aws.String("root-device-type"),
				Values: []string{"ebs"},
			},
			{
				Name:   aws.String("owner-id"),
				Values: []string{"099720109477"}, // Canonical's AWS Account ID
			},
		},
	}

	result, err := svc.DescribeImages(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to describe images, %w", err)
	}

	// Sort images by creation date to find the latest one
	sort.Slice(result.Images, func(i, j int) bool {
		return *result.Images[i].CreationDate > *result.Images[j].CreationDate
	})

	if len(result.Images) == 0 {
		return "", fmt.Errorf("no images found")
	}

	// Return the latest Ubuntu AMI ID
	return *result.Images[0].ImageId, nil
}

// createEC2Instance creates an EC2 instance with a given AMI ID
func createEC2Instance(svc *ec2.Client, amiID string) (string, error) {
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(amiID),
		InstanceType: types.InstanceTypeT2Micro, // Change to desired instance type
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	}

	// Create the instance
	result, err := svc.RunInstances(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create instance, %w", err)
	}

	// Return the ID of the new instance
	return *result.Instances[0].InstanceId, nil
}

// listEC2Instances lists all EC2 instances
func listEC2Instances(svc *ec2.Client) error {
	input := &ec2.DescribeInstancesInput{}

	// Retrieve the list of instances
	result, err := svc.DescribeInstances(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to list instances, %w", err)
	}

	// Print the instance details
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, State: %s\n", *instance.InstanceId, instance.State.Name)
		}
	}

	return nil
}

// deleteEC2Instance terminates an EC2 instance
func deleteEC2Instance(svc *ec2.Client, instanceID string) error {
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
	}

	// Terminate the instance
	_, err := svc.TerminateInstances(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to terminate instance, %w", err)
	}

	return nil
}
