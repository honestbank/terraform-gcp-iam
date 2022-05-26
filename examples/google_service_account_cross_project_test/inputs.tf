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
variable "run_id" {
  type        = string
  description = "The unique ID of the run."
}

variable "cross_project_iam_role_memberships" {
  default     = {}
  description = "A map of GCP project IDs and an associated list of IAM roles to add a membership to."
  type        = map(list(string))
  validation {
    condition     = length(var.cross_project_iam_role_memberships) < 2
    error_message = "To maintain a cleaner security model, only one project is currently supported for cross-project role memberships."
  }
}