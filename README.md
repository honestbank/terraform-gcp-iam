# Google Service Account Terraform Component Module

Use this module (found in [modules/google_service_account](./modules/google_service_account)) to create a Google Service
Account. Inputs and outputs below are for the wrapper - for the actual module to be used see [here](./modules/google_service_account).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_random"></a> [random](#provider\_random) | 3.3.2 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_google_service_account_instance"></a> [google\_service\_account\_instance](#module\_google\_service\_account\_instance) | ./modules/google_service_account | n/a |

## Resources

| Name | Type |
|------|------|
| [random_id.run_id](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/id) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_folder_id"></a> [folder\_id](#input\_folder\_id) | The GCP project folder to grant permissions into. | `string` | n/a | yes |
| <a name="input_google_credentials"></a> [google\_credentials](#input\_google\_credentials) | GCP Service Account JSON keyfile contents. | `string` | n/a | yes |
| <a name="input_google_project"></a> [google\_project](#input\_google\_project) | The GCP project to use when initializing the google Terraform provider. | `string` | n/a | yes |
| <a name="input_google_region"></a> [google\_region](#input\_google\_region) | The GCP region to use when initializing the google Terraform provider. | `string` | n/a | yes |
| <a name="input_other_project_id"></a> [other\_project\_id](#input\_other\_project\_id) | The GCP project ID of another project to grant permissions into. | `string` | n/a | yes |
| <a name="input_service_account_host_project"></a> [service\_account\_host\_project](#input\_service\_account\_host\_project) | The GCP project in which to create the service account. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_account_email"></a> [account\_email](#output\_account\_email) | The `email` attribute of the `google_service_account` resource. |
| <a name="output_service_account_key_private_keys"></a> [service\_account\_key\_private\_keys](#output\_service\_account\_key\_private\_keys) | n/a |
<!-- END_TF_DOCS -->
