terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

# Configure the AWS provider
provider "aws" {
  region = "eu-west-1"
  skip_credentials_validation = true
  skip_metadata_api_check = true
  skip_requesting_account_id = true
}

resource "aws_instance" "example" {
    ami           = "ami-0c55b159cbfafe1f0"
    instance_type = "t3.micro"
    key_name      = "terraform"
}
