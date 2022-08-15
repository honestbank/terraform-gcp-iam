module "service_account" {
  source = "../../modules/google_service_account"

  project_id = var.project_id

  account_id   = var.account_id
  description  = "An instance of the google_service_account Terraform component module."
  display_name = var.account_id

  # Create two keys
  key_aliases = [
    "foo",
    "bar",
  ]

  # Used in tests to verify that the keys can be used to make an authenticated API request.
  project_iam_memberships = [
    {
      role = "roles/browser"
    }
  ]
}
