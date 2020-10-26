# Terraform Architecture - Flugel.it interview challenge

## Interview challenge main description:
 - Creates a new VPC using 10.0.x.0/24 addresses in three different AZs
 - Creates a new EC2 instance inside this VPC
 - Creates a SG attached to the EC2 instance which only allows incoming SSH traffic.
 - Creates an S3 bucket
 - Every resource has Name = 'Flugel' and Owner = 'InfraTeam' tag names.
 - Terratest script included

## TODO:
 - Github Actions flow definition, plus previous linting and testing.
 - Only AWS cloud is supported.

## Main resources created/managed:
- VPC named 'VPC_IC' in us-east-1
- EC2 instance inside us-east-1a AZ
- S3 bucket named challenge.interview.flugel.it

## How to use this example:
- [Install Terraform](https://www.terraform.io/downloads.html). It's just a binary, you leave it anywhere under the PATH variable.
- Configure AWS credentials in ~/.aws/ or export `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.
- Configure settings in the variables.tf file (region, AZs, bucket name, etc.). Check if Terratest test module must be adjusted afterwards.
- Configure TF Cloud in the backend.tf file (optional)

## Commands:
- `terraform init` initializes the aws directory, installing all necessary plugins.
- `terraform plan` syncs `terraform.tfstate` file (or TF Cloud state) against actual cloud state, shows the changes needed to make to the cloud (but doesn't do anything).
- `terraform apply` applies the planned changes to the cloud.
- When you're done, run `terraform destroy`.

## Running automated tests against this module
- Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.
- `cd test`
- `dep ensure`
- `go test -v -run TestInterviewChallenge`