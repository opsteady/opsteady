resource "aws_route_table" "pub_a" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pub-a-${var.aws_foundation_name}"
  }
}

resource "aws_route" "pub_igw_a" {
  route_table_id         = aws_route_table.pub_a.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.igw.id
}

resource "aws_route_table_association" "pub_a" {
  subnet_id      = aws_subnet.pub_a.id
  route_table_id = aws_route_table.pub_a.id
}

resource "aws_route_table" "pub_b" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pub-b-${var.aws_foundation_name}"
  }
}

resource "aws_route" "pub_igw_b" {
  route_table_id         = aws_route_table.pub_b.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.igw.id
}

resource "aws_route_table_association" "pub_b" {
  subnet_id      = aws_subnet.pub_b.id
  route_table_id = aws_route_table.pub_b.id
}


resource "aws_route_table" "pub_c" {
  vpc_id = aws_vpc.platform.id

  tags = {
    Name = "pub-c-${var.aws_foundation_name}"
  }
}

resource "aws_route" "pub_igw_c" {
  route_table_id         = aws_route_table.pub_c.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.igw.id
}
resource "aws_route_table_association" "pub_c" {
  subnet_id      = aws_subnet.pub_c.id
  route_table_id = aws_route_table.pub_c.id
}
