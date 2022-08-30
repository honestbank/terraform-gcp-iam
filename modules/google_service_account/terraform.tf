terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.12"
    }
  }

  experiments = [module_variable_optional_attrs]
}
