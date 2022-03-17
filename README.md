# Terraform Layer Module Template Repository

Use this repository as a starting point for building a [Terraform Component Module](https://www.notion.so/honestbank/WIP-How-to-structure-a-Terraform-module-31374a1594f84ef7b185ef4e06b36619).

The recommended usage is to make this a public trunk-based development repo that automatically releases using SemVer on
merge to trunk (typically called `main`). This module is then embedded and instantiated by Layer Modules to manage live
infrastructure.

## Customizations

### Pre-commit

This template contains a [.pre-commit-config.yaml file](./.pre-commit-config.yaml). To use this, please [install pre-commit](https://pre-commit.com/#install)
and run `pre-commit install` to install hooks. The default set of hooks should work for most Terraform modules/repos - please
customize as needed.

### Releases

This template contains a [semantic-release](https://github.com/semantic-release/semantic-release) [configuration file](./release.config.js)
that is configured to produce releases on merge to `main`.

### GitHub Actions

This template contains [a 'terraform' action/workflow](./.github/workflows/terraform.yml) that is configured to run on
PRs and pushes to the `main` branch.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_random"></a> [random](#provider\_random) | 3.1.0 |

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
| <a name="input_google_credentials"></a> [google\_credentials](#input\_google\_credentials) | GCP Service Account JSON keyfile contents. | `string` | n/a | yes |
| <a name="input_google_project"></a> [google\_project](#input\_google\_project) | The GCP project to use when initializing the google Terraform provider. | `string` | n/a | yes |
| <a name="input_google_region"></a> [google\_region](#input\_google\_region) | The GCP region to use when initializing the google Terraform provider. | `string` | n/a | yes |
| <a name="input_service_account_host_project"></a> [service\_account\_host\_project](#input\_service\_account\_host\_project) | The GCP project in which to create the service account. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_account_email"></a> [account\_email](#output\_account\_email) | The `email` attribute of the `google_service_account` resource. |
| <a name="output_service_account_key_private_keys_b64"></a> [service\_account\_key\_private\_keys\_b64](#output\_service\_account\_key\_private\_keys\_b64) | n/a |
<!-- END_TF_DOCS -->
