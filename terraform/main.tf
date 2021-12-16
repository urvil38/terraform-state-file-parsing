terraform {
    required_version = "0.12.29"
}

provider "aws" {
    version = "~> 3.0"
    profile = "default"
    region = "ap-south-1"
}