# Google Service Account Terraform Component Module

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_google"></a> [google](#requirement\_google) | ~> 4.12 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | ~> 4.12 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_service_account.service_account](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account) | resource |
| [google_service_account_key.key](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account_key) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_account_id"></a> [account\_id](#input\_account\_id) | The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created. | `string` | n/a | yes |
| <a name="input_create_service_account_key"></a> [create\_service\_account\_key](#input\_create\_service\_account\_key) | n/a | `bool` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | A text description of the service account. Must be less than or equal to 256 UTF-8 bytes. | `string` | n/a | yes |
| <a name="input_display_name"></a> [display\_name](#input\_display\_name) | The display name for the service account. Can be updated without creating a new resource. | `string` | n/a | yes |
| <a name="input_project"></a> [project](#input\_project) | The ID of the project that the service account will be created in. Defaults to the provider project configuration. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_account_email"></a> [account\_email](#output\_account\_email) | The `email` attribute of the `google_service_account` resource. |
| <a name="output_service_account_key_private_key"></a> [service\_account\_key\_private\_key](#output\_service\_account\_key\_private\_key) | n/a |
<!-- END_TF_DOCS -->
