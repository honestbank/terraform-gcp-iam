# `google_service_account` Terraform Component Module Tests

## Permissions

Since this module gets/sets IAM policies, the basic Editor role (`roles/editor`) is insufficient.

For testing purposes, the `github-terratest@test-terraform-project-01.iam.gserviceaccount.com` Google
Service Account is given the Owner role (`roles/owner`) on the `test-terraform-project-01` GCP project.

## Authentication

The `TERRATEST_GOOGLE_CREDENTIALS` environment variable should be set with the service account JSON keyfile contents.
