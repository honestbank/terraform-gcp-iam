variable "google_credentials" {
  description = "GCP Service Account JSON keyfile contents."
}

variable "google_region" {
  description = "GCP region to use."
}

provider "google" {
  // The google_service_account module uses `project`.
  project     = var.project_id // Defined in examples
  region      = var.google_region
  credentials = var.google_credentials
}
