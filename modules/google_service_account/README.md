# Google Service Account Terraform Component Module

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_google"></a> [google](#requirement\_google) | >= 4.12 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | 4.29.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_folder_iam_member.folder_iam_role_memberships](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/folder_iam_member) | resource |
| [google_project_iam_member.cross_project_iam_role_memberships](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_project_iam_member.in_project_roles](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_service_account.service_account](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account) | resource |
| [google_service_account_key.keys](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account_key) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_account_id"></a> [account\_id](#input\_account\_id) | The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created. | `string` | n/a | yes |
| <a name="input_conditions"></a> [conditions](#input\_conditions) | A list of conditions to be applied to in project service account.<br>     Example:<pre>module "google_service_account_instance" {<br>       source = "./modules/google_service_account"<br><br>       account_id   = "terraform-id"<br>       display_name = "google_service_account_instance"<br>       description  = "An instance of the google_service_account Terraform component module."<br><br>       in_project_roles = ["roles/viewer"]<br><br>       conditions = [{<br>         title = "User is in the same organization as the Terraform project"<br>         description = "The user is in the same organization as the Terraform project."<br>         expression = "request.resource.labels.organization_id == project.project_id"<br>       }]<br><br>       key_aliases = ["primary", "secondary", "another_key"]<br>       project     = var.service_account_host_project<br>     }</pre> | <pre>list(object({<br>    title       = string,<br>    description = string,<br>    expression  = string,<br>  }))</pre> | `[]` | no |
| <a name="input_cross_project_iam_role_memberships"></a> [cross\_project\_iam\_role\_memberships](#input\_cross\_project\_iam\_role\_memberships) | A map of GCP project IDs and an associated list of IAM roles to add a membership to. | `map(list(string))` | `{}` | no |
| <a name="input_description"></a> [description](#input\_description) | A text description of the service account. Must be less than or equal to 256 UTF-8 bytes. | `string` | n/a | yes |
| <a name="input_display_name"></a> [display\_name](#input\_display\_name) | The display name for the service account. Can be updated without creating a new resource. | `string` | n/a | yes |
| <a name="input_folder_iam_role_memberships"></a> [folder\_iam\_role\_memberships](#input\_folder\_iam\_role\_memberships) | A map of GCP folder IDs and an associated list of IAM roles to add a membership to. | `map(list(string))` | `{}` | no |
| <a name="input_iam_role_membership_type"></a> [iam\_role\_membership\_type](#input\_iam\_role\_membership\_type) | One of [CROSS\_PROJECT, FOLDER, IN\_PROJECT]. | `string` | `"IN_PROJECT"` | no |
| <a name="input_in_project_roles"></a> [in\_project\_roles](#input\_in\_project\_roles) | Roles to assign service account within its own project. | `list(string)` | `[]` | no |
| <a name="input_key_aliases"></a> [key\_aliases](#input\_key\_aliases) | A JSON key will be created and output for each entry in this list. | `list(string)` | `[]` | no |
| <a name="input_project"></a> [project](#input\_project) | The ID of the project that the service account will be created in. Defaults to the provider project configuration. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_account_email"></a> [account\_email](#output\_account\_email) | The `email` attribute of the `google_service_account` resource. |
| <a name="output_service_account_keys"></a> [service\_account\_keys](#output\_service\_account\_keys) | A map of the service account keys created, with each item in the key\_alias as a key. Returns `ERROR` if the `google_service_account_key.keys` resource cannot be accessed. |
| <a name="output_service_account_name"></a> [service\_account\_name](#output\_service\_account\_name) | Theh name of `google_service_account` resource. |
<!-- END_TF_DOCS -->
