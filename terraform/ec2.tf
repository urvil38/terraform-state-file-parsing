data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "private-instance" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = "t3.micro"
  subnet_id                   = aws_subnet.private-subnet-1.id
  vpc_security_group_ids      = ["${aws_security_group.ec2-instance-security-group.id}"]
  associate_public_ip_address = false
  key_name                    = "assignment-last9-ssh"
  tags = {
    Name = "private-instance"
  }
}

resource "aws_instance" "public-instance" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = "t3.micro"
  subnet_id                   = aws_subnet.public-subnet-1.id
  vpc_security_group_ids      = ["${aws_security_group.ec2-instance-security-group.id}"]
  associate_public_ip_address = true
  key_name                    = "assignment-last9-ssh"
  tags = {
    Name = "public-instance"
  }
}
