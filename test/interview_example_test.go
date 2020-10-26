package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestInterviewChallenge(t *testing.T) {
	// Run this test in parallel
	t.Parallel()

	// TODO: Pick a random AWS region to test in. This helps ensure your code works in all regions.
	// awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// *** Adjust this settings related to variables.tf ***
	awsRegion := "us-east-1"
	expectedNameTag := "Flugel"
	expectedOwnerTag := "InfraTeam"
	flugelBucket := "challenge.interview.flugel.it"
	// ***

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// TODO: Execute this test in a random AWS region
		// Environment variables to set when running Terraform
		// EnvVars: map[string]string{
		// 	"AWS_DEFAULT_REGION": awsRegion,
		// },
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// EC2 instance test
	instanceID := terraform.Output(t, terraformOptions, "interview_challgenge_instance_id")

	// Get all the tags from the created instance
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// Verify our expected tags Name, Owner are equal to the expected ones
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedNameTag, nameTag)

	ownerTag, containsOwnerTag := instanceTags["Owner"]
	assert.True(t, containsOwnerTag)
	assert.Equal(t, expectedOwnerTag, ownerTag)

	// S3 Bucket tests
	bucketFound := aws.FindS3BucketWithTag(t, awsRegion, "Name", expectedNameTag)
	assert.Equal(t, flugelBucket, bucketFound)

	bucketFound = aws.FindS3BucketWithTag(t, awsRegion, "Owner", expectedOwnerTag)
	assert.Equal(t, flugelBucket, bucketFound)
}
