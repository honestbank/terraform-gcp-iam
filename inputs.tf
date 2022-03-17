variable "google_credentials" {
  type        = string
  description = "GCP Service Account JSON keyfile contents."
}

variable "google_project" {
  type        = string
  description = "The GCP project to use when initializing the google Terraform provider."
}

variable "google_region" {
  type        = string
  description = "The GCP region to use when initializing the google Terraform provider."
}

variable "service_account_host_project" {
  type        = string
  description = "The GCP project in which to create the service account."
}
