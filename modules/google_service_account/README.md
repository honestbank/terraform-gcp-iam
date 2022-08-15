# Google Service Account Terraform Component Module

Module for managing Google Service Accounts plus:
* Service Account Keys
* IAM bindings for projects
* IAM bindings for folders

> NOTE: Due to the very large number of possible IAM bindings, only the most basic are included, you will need to
> use the rest of them outside this module. See the [IAM documentation](https://cloud.google.com/iam/docs/overview)
> for more details.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_google"></a> [google](#requirement\_google) | >= 4.12 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | >= 4.12 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_folder_iam_member.folder_iam_memberships](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/folder_iam_member) | resource |
| [google_project_iam_member.cross_project_iam_memberships](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_project_iam_member.project_iam_memberships](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_service_account.service_account](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account) | resource |
| [google_service_account_key.keys](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account_key) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_account_id"></a> [account\_id](#input\_account\_id) | The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created. | `string` | n/a | yes |
| <a name="input_cross_project_iam_memberships"></a> [cross\_project\_iam\_memberships](#input\_cross\_project\_iam\_memberships) | A map of project IDs with a list of IAM roles with optional conditions to add memberships for. | <pre>map(list(object({<br>    role = string<br>    conditions = optional(list(object({<br>      description = string<br>      expression  = string<br>      title       = string<br>    })))<br>  })))</pre> | `{}` | no |
| <a name="input_description"></a> [description](#input\_description) | A text description of the service account. Must be less than or equal to 256 UTF-8 bytes. | `string` | n/a | yes |
| <a name="input_display_name"></a> [display\_name](#input\_display\_name) | The display name for the service account. Can be updated without creating a new resource. | `string` | n/a | yes |
| <a name="input_folder_iam_memberships"></a> [folder\_iam\_memberships](#input\_folder\_iam\_memberships) | A map of folder IDs with a list of IAM roles with optional conditions to add memberships for. | <pre>map(list(object({<br>    role = string<br>    conditions = optional(list(object({<br>      description = string<br>      expression  = string<br>      title       = string<br>    })))<br>  })))</pre> | `{}` | no |
| <a name="input_key_aliases"></a> [key\_aliases](#input\_key\_aliases) | A JSON key will be created and output for each entry in this list. | `list(string)` | `[]` | no |
| <a name="input_project_iam_memberships"></a> [project\_iam\_memberships](#input\_project\_iam\_memberships) | A list of IAM roles with optional conditions to add memberships for within the same project. | <pre>list(object({<br>    role = string<br>    conditions = optional(list(object({<br>      description = string<br>      expression  = string<br>      title       = string<br>    })))<br>  }))</pre> | `[]` | no |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | The ID of the project that the service account will be created in. Defaults to the provider project configuration. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_email"></a> [email](#output\_email) | The `email` attribute of the `google_service_account` resource. |
| <a name="output_keys"></a> [keys](#output\_keys) | A map of the service account keys created, with each item in the key\_alias as a base64 encoded key.  Returns `ERROR` if the `google_service_account_key.keys` resource cannot be accessed. |
| <a name="output_name"></a> [name](#output\_name) | The name of `google_service_account` resource. |
<!-- END_TF_DOCS -->
