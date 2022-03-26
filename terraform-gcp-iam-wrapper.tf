terraform {
  required_version = "~> 1.0"
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
  display_name = "google_service_account_instance"
  description  = "An instance of the google_service_account Terraform component module."

  in_project_roles = ["roles/viewer"]

  key_aliases = ["primary", "secondary", "another_key"]
  project     = var.service_account_host_project
}
