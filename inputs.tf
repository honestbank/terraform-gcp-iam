variable "google_credentials" {
  description = "GCP Service Account JSON keyfile contents."
  type        = string
}

variable "google_project" {
  description = "The GCP project to use when initializing the google Terraform provider."
  type        = string
}

variable "google_region" {
  description = "The GCP region to use when initializing the google Terraform provider."
  type        = string
}

variable "service_account_host_project" {
  description = "The GCP project in which to create the service account."
  type        = string
}

variable "other_project_id" {
  description = "The GCP project ID of another project to grant permissions into."
  type        = string
}

variable "folder_id" {
  description = "The GCP project folder to grant permissions into."
  type        = string
}
