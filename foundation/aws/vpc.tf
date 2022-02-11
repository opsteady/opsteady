resource "aws_vpc" "platform" {
  cidr_block = var.aws_foundation_vpc_cidr

  tags = {
    Name = var.aws_foundation_name
  }
}
