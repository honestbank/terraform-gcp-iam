variable "account_id" {
  description = "The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created."
  type        = string
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

variable "description" {
  description = "A text description of the service account. Must be less than or equal to 256 UTF-8 bytes."
  type        = string
}

variable "display_name" {
  description = "The display name for the service account. Can be updated without creating a new resource."
  type        = string
}

variable "folder_iam_role_memberships" {
  default     = {}
  description = "A map of GCP folder IDs and an associated list of IAM roles to add a membership to."
  type        = map(list(string))
  validation {
    condition     = length(var.folder_iam_role_memberships) < 2
    error_message = "To maintain a cleaner security model, only one folder is currently supported for folder role memberships."
  }
}

variable "iam_role_membership_type" {
  default     = "IN_PROJECT"
  description = "One of [CROSS_PROJECT, FOLDER, IN_PROJECT]."
  type        = string
  validation {
    condition     = contains(["CROSS_PROJECT", "FOLDER", "IN_PROJECT"], var.iam_role_membership_type)
    error_message = "Must be one of [CROSS_PROJECT, FOLDER, IN_PROJECT]."
  }
}

variable "in_project_roles" {
  default     = []
  description = "Roles to assign service account within its own project."
  type        = list(string)
}

variable "key_aliases" {
  default     = []
  description = "A JSON key will be created and output for each entry in this list."
  type        = list(string)
}

variable "project" {
  description = "The ID of the project that the service account will be created in. Defaults to the provider project configuration."
  type        = string
}

variable "in_project_condition_title" {
  type        = string
  description = "The title of the condition that will be assigned to the service account within its own project. Requires title and expression to be set."
  default     = null
}

variable "in_project_condition_description" {
  type        = string
  description = "The description of the condition that will be assigned to the service account within its own project. Requires title and expression to be set."
  default     = null
}

variable "in_project_condition_expression" {
  type        = string
  description = "The expression of the condition that will be assigned to the service account within its own project (use <<~EOT EOT). Requires title and expression to be set."
  default     = null
}

variable "cross_project_condition_title" {
  type        = string
  description = "The title of the condition that will be assigned to the service account for cross-project. Requires title and expression to be set."
  default     = null
}

variable "cross_project_condition_description" {
  type        = string
  description = "The description of the condition that will be assigned to the service account for cross-project. Requires title and expression to be set."
  default     = null
}

variable "cross_project_condition_expression" {
  type        = string
  description = "The expression of the condition that will be assigned to the service account for cross-project (use <<~EOT EOT). Requires title and expression to be set."
  default     = null
}

variable "folder_condition_title" {
  type        = string
  description = "The title of the condition that will be assigned to the service account within its folder. Requires title and expression to be set."
  default     = null
}

variable "folder_condition_description" {
  type        = string
  description = "The description of the condition that will be assigned to the service account within its folder. Requires title and expression to be set."
  default     = null
}

variable "folder_condition_expression" {
  type        = string
  description = "The expression of the condition that will be assigned to the service account within its folder (use <<~EOT EOT). Requires title and expression to be set."
  default     = null
}
