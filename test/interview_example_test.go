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

	expectedNameTag := "Flugel"
	expectedOwnerTag := "InfraTeam"

	awsRegion := "us-east-1"

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

	// Run `terraform output` to get the value of an output variable
	instanceID := terraform.Output(t, terraformOptions, "inteview_challgenge_instance_id")

	// Get all the tags from the created instance
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// Verify our expected tags Name, Owner are equal to the expected ones
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedNameTag, nameTag)

	ownerTag, containsOwnerTag := instanceTags["Owner"]
	assert.True(t, containsOwnerTag)
	assert.Equal(t, expectedOwnerTag, ownerTag)
}
