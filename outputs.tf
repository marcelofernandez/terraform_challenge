output "interview_challenge_public_ip" {
  value = aws_instance.interview_challenge.public_ip
}

output "interview_challenge_private_ip" {
  value = aws_instance.interview_challenge.private_ip
}

output "vpc_public_subnets" {
  value = module.vpc.public_subnets
}

# TODO: Export Tags to be asserted by terratest