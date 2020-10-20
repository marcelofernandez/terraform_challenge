terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region     = var.aws_region
}

# Deploy an EC2 instance running last version of Ubuntu 20.04
resource "aws_instance" "interview_example" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.nano"

  root_block_device {
    volume_size = "8"
    volume_type = "gp2"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    aws_security_group.sg_interview_example.id,
  ]

  tags = {
    Name = var.res_name
    Owner = var.res_owner
  }
}
