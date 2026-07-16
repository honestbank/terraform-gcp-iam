variable "account_id" {
  type        = string
  default     = "something-"
  description = "The ID of the GCP service account."
}

variable "project_id" {
  type        = string
  description = "The ID of the GCP project."
}
