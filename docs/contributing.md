# Contributing to this repo

## Testing this repo
In order to test this repo, use Terratest github secrets to be able to run the tests.
* secrets.TERRATEST_GOOGLE_PROJECT
* secrets.TERRATEST_GOOGLE_CREDENTIALS
* secrets.TERRATEST_GCP_SA_EMAIL

These will ensure the right projects are used with the right credentials.
Additionally for this repo you must enable permissions if they are not already enabled:
* enable permission `iam.serviceAccounts.create`
* enable permission `resourcemanager.folders.getIamPolicy`


## A note on service account conditions
Google does not allow you to create conditions on primitive roles. These roles are:
* `roles/viewer`
* `roles/admin`
* `roles/owner`

Instead, roles should be scoped down such as `roles/storage.objectViewer` or custom roles. see [here](https://cloud.google.com/storage/docs/access-control/iam-permissions)
