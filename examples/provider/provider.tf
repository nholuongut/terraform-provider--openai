terraform {
  required_providers {
    openai = {
      source  = "nholuong/openai"
      version = "2.3.2"
    }
  }
}

provider "openai" {
  # Configuration options
}