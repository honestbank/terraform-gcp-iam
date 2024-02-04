terraform {
  required_version = "~> 1.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.14.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
  }
}



provider "google" {
  region      = var.google_region
  project     = var.google_project
  credentials = var.google_credentials
}

resource "random_id" "run_id" {
  byte_length = 4
}

module "google_service_account_instance" {
  source = "./modules/google_service_account"

  account_id   = "terraform-${random_id.run_id.hex}"
  project_id   = var.service_account_host_project
  display_name = "Terraform ${random_id.run_id.hex}"
  description  = "An instance of the google_service_account Terraform component module."

  project_iam_memberships = [
    {
      role = "roles/viewer"
    }
  ]

  cross_project_iam_memberships = {
    (var.other_project_id) = [
      {
        role = "roles/viewer"
      }
    ]
  }

  folder_iam_memberships = {
    (var.folder_id) = [
      {
        role = "roles/viewer"
      }
    ]
  }

  key_aliases = [
    "primary",
    "secondary",
    "another_key"
  ]
}
