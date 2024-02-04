variable "account_id" {
  description = "The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created."
  type        = string
  validation {
    condition     = length(var.account_id) >= 6 && length(var.account_id) <= 30 && regex("[a-z]([-a-z0-9]*[a-z0-9])", var.account_id) == var.account_id
    error_message = "The account_id must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035."
  }
}

variable "cross_project_iam_memberships" {
  default     = {}
  description = "A map of project IDs with a list of IAM roles with optional conditions to add memberships for."
  type = map(list(object({
    role = string
    conditions = optional(list(object({
      description = string
      expression  = string
      title       = string
    })))
  })))
}

variable "description" {
  description = "A text description of the service account. Must be less than or equal to 256 UTF-8 bytes."
  type        = string
}

variable "display_name" {
  description = "The display name for the service account. Can be updated without creating a new resource."
  type        = string
}

variable "folder_iam_memberships" {
  default     = {}
  description = "A map of folder IDs with a list of IAM roles with optional conditions to add memberships for."
  type = map(list(object({
    role = string
    conditions = optional(list(object({
      description = string
      expression  = string
      title       = string
    })))
  })))
}

variable "project_iam_memberships" {
  default     = []
  description = "A list of IAM roles with optional conditions to add memberships for within the same project."
  type = list(object({
    role = string
    conditions = optional(list(object({
      description = string
      expression  = string
      title       = string
    })))
  }))
}

variable "key_aliases" {
  default     = []
  description = "A JSON key will be created and output for each entry in this list."
  type        = list(string)
}

variable "project_id" {
  description = "The ID of the project that the service account will be created in. Defaults to the provider project configuration."
  type        = string
}
