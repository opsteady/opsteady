resource "aws_subnet" "pods_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pods_a
  availability_zone = var.foundation_aws_zone_pods_a

  tags = {
    Name = "pods-a-${var.foundation_aws_name}"
  }
}

resource "aws_subnet" "pods_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pods_b
  availability_zone = var.foundation_aws_zone_pods_b

  tags = {
    Name = "pods-b-${var.foundation_aws_name}"
  }
}

resource "aws_subnet" "pods_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pods_c
  availability_zone = var.foundation_aws_zone_pods_c

  tags = {
    Name = "pods-c-${var.foundation_aws_name}"
  }
}

resource "aws_subnet" "pub_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pub_a
  availability_zone = var.foundation_aws_zone_pub_a

  tags = {
    Name = "pub-a-${var.foundation_aws_name}"
  }
}

resource "aws_subnet" "pub_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pub_b
  availability_zone = var.foundation_aws_zone_pub_b

  tags = {
    Name = "pub-b-${var.foundation_aws_name}"
  }
}

resource "aws_subnet" "pub_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.foundation_aws_subnet_pub_c
  availability_zone = var.foundation_aws_zone_pub_c

  tags = {
    Name = "pub-c-${var.foundation_aws_name}"
  }
}
