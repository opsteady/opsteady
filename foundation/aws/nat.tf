resource "aws_eip" "nat_a" {
  count = var.foundation_aws_nat_a_enabeld ? 1 : 0
  vpc   = true

  tags = {
    Name = "nat-a-${var.foundation_aws_name}"
  }
}

resource "aws_eip" "nat_b" {
  count = var.foundation_aws_nat_b_enabeld ? 1 : 0

  vpc = true

  tags = {
    Name = "nat-b-${var.foundation_aws_name}"
  }
}

resource "aws_eip" "nat_c" {
  count = var.foundation_aws_nat_c_enabeld ? 1 : 0

  vpc = true

  tags = {
    Name = "nat-c-${var.foundation_aws_name}"
  }
}

resource "aws_nat_gateway" "nat_a" {
  count = var.foundation_aws_nat_a_enabeld ? 1 : 0

  allocation_id = aws_eip.nat_a.0.id
  subnet_id     = aws_subnet.pub_a.id

  tags = {
    Name = "nat-a-${var.foundation_aws_name}"
  }
}

resource "aws_nat_gateway" "nat_b" {
  count = var.foundation_aws_nat_b_enabeld ? 1 : 0

  allocation_id = aws_eip.nat_b.0.id
  subnet_id     = aws_subnet.pub_b.id

  tags = {
    Name = "nat-b-${var.foundation_aws_name}"
  }
}

resource "aws_nat_gateway" "nat_c" {
  count = var.foundation_aws_nat_c_enabeld ? 1 : 0

  allocation_id = aws_eip.nat_c.0.id
  subnet_id     = aws_subnet.pub_c.id

  tags = {
    Name = "nat-c-${var.foundation_aws_name}"
  }
}
