resource "aws_subnet" "eks_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_eks_a
  availability_zone = var.aws_foundation_zone_eks_a

  tags = {
    Name = "eks-a-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "eks_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_eks_b
  availability_zone = var.aws_foundation_zone_eks_b

  tags = {
    Name = "eks-b-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "eks_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_eks_c
  availability_zone = var.aws_foundation_zone_eks_c

  tags = {
    Name = "eks-c-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "pods_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pods_a
  availability_zone = var.aws_foundation_zone_pods_a

  tags = {
    Name = "pods-a-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "pods_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pods_b
  availability_zone = var.aws_foundation_zone_pods_b

  tags = {
    Name = "pods-b-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "pods_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pods_c
  availability_zone = var.aws_foundation_zone_pods_c

  tags = {
    Name = "pods-c-${var.aws_foundation_name}"
  }
}

resource "aws_subnet" "pub_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pub_a
  availability_zone = var.aws_foundation_zone_pub_a

  tags = {
    Name = "pub-a-${var.aws_foundation_name}"
    "kubernetes.io/role/elb" : "1"
  }
}

resource "aws_subnet" "pub_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pub_b
  availability_zone = var.aws_foundation_zone_pub_b

  tags = {
    Name = "pub-b-${var.aws_foundation_name}"
    "kubernetes.io/role/elb" : "1"
  }
}

resource "aws_subnet" "pub_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_pub_c
  availability_zone = var.aws_foundation_zone_pub_c

  tags = {
    Name = "pub-c-${var.aws_foundation_name}"
    "kubernetes.io/role/elb" : "1"
  }
}

resource "aws_subnet" "prv_a" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_prv_a
  availability_zone = var.aws_foundation_zone_prv_a

  tags = {
    Name = "prv-a-${var.aws_foundation_name}"
    "kubernetes.io/role/interal-elb" : "1"
  }
}

resource "aws_subnet" "prv_b" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_prv_b
  availability_zone = var.aws_foundation_zone_prv_b

  tags = {
    Name = "prv-b-${var.aws_foundation_name}"
    "kubernetes.io/role/internal-elb" : "1"
  }
}

resource "aws_subnet" "prv_c" {
  vpc_id            = aws_vpc.platform.id
  cidr_block        = var.aws_foundation_subnet_prv_c
  availability_zone = var.aws_foundation_zone_prv_c

  tags = {
    Name = "prv-c-${var.aws_foundation_name}"
    "kubernetes.io/role/internal-elb" : "1"
  }
}
