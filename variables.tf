variable "aws_region" {
  description = "The AWS region to deploy into"
  type        = string
  default     = "us-east-1"
}

variable "res_name" {
  description = "The Name tag to set for the resources created"
  type        = string
  default     = "Flugel"
}

variable "res_owner" {
  description = "The owner of the resources created"
  type = string
  default = "InfraTeam"
}

# Dinamically search the latest Ubuntu 20.04 AMIs
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }

  filter {
    name   = "image-type"
    values = ["machine"]
  }

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}