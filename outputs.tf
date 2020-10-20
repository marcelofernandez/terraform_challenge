output "interview_example_public_ip" {
  value = aws_instance.interview_example.public_ip
}

output "interview_example_private_ip" {
  value = aws_instance.interview_example.private_ip
}

# TODO: Export Tags to ve asserted by terratest