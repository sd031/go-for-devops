terraform {
  required_providers {
    example = {
      source = "example.com/example/example"
      version = "0.1.0"
    }
  }
}

provider "example" {
}

resource "example_resource" "my_resource" {
  example_value = "example_name"
}