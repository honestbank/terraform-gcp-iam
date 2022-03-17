variable "google_credentials" {
  description = "GCP Service Account JSON keyfile contents."
}

variable "google_region" {
  description = "GCP region to use."
}

provider "google" {

  // The google_service_account module uses `project`.
  project     = var.project
  region      = var.google_region
  credentials = var.google_credentials
}
