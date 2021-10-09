locals {
  gateway_a = var.foundation_aws_nat_a_enabeld ? aws_nat_gateway.nat_a.0.id : var.foundation_aws_nat_b_enabeld ? aws_nat_gateway.nat_b.0.id : aws_nat_gateway.nat_c.0.id
  gateway_b = var.foundation_aws_nat_b_enabeld ? aws_nat_gateway.nat_b.0.id : var.foundation_aws_nat_a_enabeld ? aws_nat_gateway.nat_a.0.id : aws_nat_gateway.nat_c.0.id
  gateway_c = var.foundation_aws_nat_c_enabeld ? aws_nat_gateway.nat_c.0.id : var.foundation_aws_nat_a_enabeld ? aws_nat_gateway.nat_a.0.id : aws_nat_gateway.nat_b.0.id
}

resource "aws_route_table" "pods_a" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pods-a-${var.foundation_aws_name}"
  }
}

resource "aws_route" "pods_nat_a" {
  route_table_id         = aws_route_table.pods_a.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = local.gateway_a
}

resource "aws_route_table_association" "pods_a" {
  subnet_id      = aws_subnet.pods_a.id
  route_table_id = aws_route_table.pods_a.id
}

resource "aws_route_table" "pods_b" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pods-b-${var.foundation_aws_name}"
  }
}

resource "aws_route" "pods_nat_b" {
  route_table_id         = aws_route_table.pods_b.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = local.gateway_b
}

resource "aws_route_table_association" "pods_b" {
  subnet_id      = aws_subnet.pods_b.id
  route_table_id = aws_route_table.pods_b.id
}

resource "aws_route_table" "pods_c" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pods-c-${var.foundation_aws_name}"
  }
}

resource "aws_route" "pods_nat_c" {
  route_table_id         = aws_route_table.pods_c.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = local.gateway_c
}

resource "aws_route_table_association" "pods_c" {
  subnet_id      = aws_subnet.pods_c.id
  route_table_id = aws_route_table.pods_c.id
}
