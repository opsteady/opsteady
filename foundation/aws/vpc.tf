resource "aws_vpc" "platform" {
  cidr_block = var.foundation_aws_vpc_cidr

  tags = {
    Name = var.foundation_aws_name
  }
}
