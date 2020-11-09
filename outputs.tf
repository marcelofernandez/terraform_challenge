output "interview_challenge_public_ip" {
  value = aws_instance.interview_challenge.public_ip
}

output "interview_challenge_private_ip" {
  value = aws_instance.interview_challenge.private_ip
}

output "interview_challgenge_instance_id" {
  value = aws_instance.interview_challenge.id
}
