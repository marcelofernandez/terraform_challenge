# SG attached to the EC2 instance
resource "aws_security_group" "sg_interview_challenge" {
  name        = "allow_ssh"
  description = "Allow SSH traffic"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    # Allow ssh connections from anywhere
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.common_tags
}