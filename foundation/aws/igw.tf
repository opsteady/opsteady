resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "igw-${var.foundation_aws_name}"
  }
}
