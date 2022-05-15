# `google_service_account` Terraform Component Module Tests

## Permissions

Since this module gets/sets IAM policies, the basic Editor role (`roles/editor`) is insufficient.

### In-Project IAM Role Memberships

* This is specified by setting the `iam_role_membership_type` to `IN_PROJECT`.
* The `github-terratest@test-terraform-project-01.iam.gserviceaccount.com` Google
Service Account is given the Owner role (`roles/owner`) on the `test-terraform-project-01` GCP project.

### Cross-Project IAM Role Memberships

>
> Since the project specified below (`storage-0994`) is already located in the folder specified by the following section,
> permissions have been granted at the folder-level, rather than both project and folder level.
>

* This is specified by setting the `iam_role_membership_type` to `CROSS_PROJECT`.
* The `github-terratest@test-terraform-project-01.iam.gserviceaccount.com` Google
  Service Account is given the Owner role (`roles/owner`) on the `storage-0994` GCP project.

### Folder-Level IAM Role Memberships

* This is specified by setting the `iam_role_membership_type` to `FOLDER`.
* The `github-terratest@test-terraform-project-01.iam.gserviceaccount.com` Google
  Service Account is given the Owner role (`roles/owner`) on the `502911218937` GCP folder (`terraform automated testing`).

## Authentication

The `TERRATEST_GOOGLE_CREDENTIALS` environment variable should be set with the service account JSON keyfile contents.
