resource "google_service_account" "service_account" {
  account_id   = var.account_id
  description  = var.description
  display_name = var.display_name
  project      = var.project
}

resource "google_service_account_key" "keys" {
  for_each = toset(var.key_aliases)

  service_account_id = google_service_account.service_account.account_id
}

resource "google_project_iam_member" "in_project_roles" {
  for_each = toset(var.in_project_roles)

  project = var.project
  role    = each.value
  member  = "serviceAccount:${google_service_account.service_account.email}"
}
