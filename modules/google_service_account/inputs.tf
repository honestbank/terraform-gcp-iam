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

variable "in_project_conditions" {
  type = list(object({
    title       = string,
    description = string,
    expression  = string,
  }))
  default     = []
  description = <<DESC
     A list of conditions to be applied to in project service account.
     Example:
     ```
     module "google_service_account_instance" {
       source = "./modules/google_service_account"

       account_id   = "terraform-id"
       display_name = "google_service_account_instance"
       description  = "An instance of the google_service_account Terraform component module."

       in_project_roles = ["roles/viewer"]

       in_project_conditions = [{
         title = "User is in the same organization as the Terraform project"
         description = "The user is in the same organization as the Terraform project."
         expression = "request.resource.labels.organization_id == project.project_id"
       }]

       key_aliases = ["primary", "secondary", "another_key"]
       project     = var.service_account_host_project
     }
     ```
  DESC
}

variable "cross_project_conditions" {
  type = list(object({
    title       = string,
    description = string,
    expression  = string,
  }))
  default     = []
  description = <<DESC
     A list of conditions to be applied to in project service account.
     Example:
     ```
     module "google_service_account_instance" {
       source = "./modules/google_service_account"

       account_id   = "terraform-id"
       display_name = "google_service_account_instance"
       description  = "An instance of the google_service_account Terraform component module."

       cross_project_iam_role_memberships = ["somemembership"]

       cross_project_conditions = [{
         title = "User is in the same organization as the Terraform project"
         description = "The user is in the same organization as the Terraform project."
         expression = "request.resource.labels.organization_id == project.project_id"
       }]

       key_aliases = ["primary", "secondary", "another_key"]
       project     = var.service_account_host_project
     }
     ```
  DESC
}

variable "folder_conditions" {
  type = list(object({
    title       = string,
    description = string,
    expression  = string,
  }))
  default     = []
  description = <<DESC
     A list of conditions to be applied to in project service account.
     Example:
     ```
     module "google_service_account_instance" {
       source = "./modules/google_service_account"

       account_id   = "terraform-id"
       display_name = "google_service_account_instance"
       description  = "An instance of the google_service_account Terraform component module."

       folder_iam_role_memberships = ["folder_memberships"]

       folder_conditions = [{
         title = "User is in the same organization as the Terraform project"
         description = "The user is in the same organization as the Terraform project."
         expression = "request.resource.labels.organization_id == project.project_id"
       }]

       key_aliases = ["primary", "secondary", "another_key"]
       project     = var.service_account_host_project
     }
     ```
  DESC
}
