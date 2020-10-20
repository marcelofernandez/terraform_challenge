# Terraform Architecture - Flugel.it interview test

## Interview Test main description:
 - Creates a new EC2 instance
 - Creates a SG which only allows incoming SSH traffic.
 - Creates an S3 bucket [WIP]
 - Every resource has Name = 'Flugel' and Owner = 'InfraTeam' tag names.
 - Terratest script

## TODO:
 - Create the EC2 instance in a new VPC (ATM this script uses default VPC).
 - Only AWS cloud is supported, GCP and Azure is planned in the future.

## Main resources created/managed:
- EC2 instance named 'interview_example'
- S3 bucket exam.interview.flugel.it

## How to use this example:
- [Install Terraform](https://www.terraform.io/downloads.html). It's just a binary, you leave it anywhere under the PATH variable.
- Configure AWS credentials in ~/.aws/ or export `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.
- Configure settings in the variables.tf file.

### Commands:
- `go test -v -run TestInterviewExample` runs the terratest script
- `terraform init` initializes the aws directory, installing all necessary plugins.
- `terraform plan` syncs `terraform.tfstate` file against actual cloud state, shows the changes needed to make to the cloud (but doesn't do anything).
- `terraform apply` applies the planned changes to the cloud.
