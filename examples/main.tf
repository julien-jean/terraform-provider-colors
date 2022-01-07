terraform {
  required_providers {
    colors = {
      version = "0.2.0"
      source  = "hashicorp.com/julien-jean/colors"
    }
  }
}

provider "colors" {
}

data "colors_random_colors" "three_colors" {
  number = 3
}

resource "null_resource" "cluster" {
  for_each = data.colors_random_colors.three_colors
}

output "foo_colors" {
  value = data.colors_random_colors.three_colors
}
