terraform {
  required_providers {
    julien = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/julien"
    }
  }
}

provider "julien" {
}

data "julien_color" "three_colors" {
  number = 3
}

output "foo_colors" {
  value = data.julien_color.three_colors.colors
}
