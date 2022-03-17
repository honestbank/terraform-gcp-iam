variable "account_id" {
  type        = string
  description = "The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created."
}

variable "description" {
  type        = string
  description = "A text description of the service account. Must be less than or equal to 256 UTF-8 bytes."
}

variable "display_name" {
  type        = string
  description = "The display name for the service account. Can be updated without creating a new resource."
}

variable "in_project_roles" {
  type        = list(string)
  description = "Roles to assign service account within its own project."
}

variable "key_aliases" {
  type        = list(string)
  description = "A JSON key will be created and output for each entry in this list."
  default     = []
}

variable "project" {
  type        = string
  description = "The ID of the project that the service account will be created in. Defaults to the provider project configuration."
}
