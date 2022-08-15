resource "google_service_account" "service_account" {
  account_id   = var.account_id
  description  = var.description
  display_name = var.display_name
  project      = var.project_id
}

resource "google_service_account_key" "keys" {
  for_each = toset(var.key_aliases)

  # https://github.com/hashicorp/terraform-provider-google/issues/9617
  # project is implicitly passed from google_service_account ID as projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}
  service_account_id = google_service_account.service_account.id
}
