locals {

  CONST_IAM_TYPE_CROSS_PROJECT = "CROSS_PROJECT"
  CONST_IAM_TYPE_FOLDER        = "FOLDER"
  CONST_IAM_TYPE_IN_PROJECT    = "IN_PROJECT"

  cross_project_iam_role_membership_count = try(
    count(var.cross_project_iam_role_memberships[one(keys(var.cross_project_iam_role_memberships))]),
    0
  )
  cross_project_iam_role_membership_project_id = (var.iam_role_membership_type == local.CONST_IAM_TYPE_CROSS_PROJECT && length(var.cross_project_iam_role_memberships) == 1) ? one(keys(var.cross_project_iam_role_memberships)) : null
  cross_project_iam_role_memberships           = (local.cross_project_iam_role_membership_project_id != null) ? var.cross_project_iam_role_memberships[local.cross_project_iam_role_membership_project_id] : []

  folder_iam_role_membership_count = try(
    count(var.folder_iam_role_memberships[one(keys(var.folder_iam_role_memberships))]),
    0
  )
  folder_iam_role_membership_folder_id = (var.iam_role_membership_type == local.CONST_IAM_TYPE_FOLDER && length(var.folder_iam_role_memberships) == 1) ? one(keys(var.folder_iam_role_memberships)) : null
  folder_iam_role_memberships          = (local.folder_iam_role_membership_folder_id != null) ? var.folder_iam_role_memberships[local.folder_iam_role_membership_folder_id] : []

  in_project_role_memberships = (var.iam_role_membership_type == local.CONST_IAM_TYPE_IN_PROJECT && length(var.in_project_roles) > 0) ? var.in_project_roles : []
}

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
  for_each = toset(local.in_project_role_memberships)

  project = var.project
  role    = each.value
  member  = "serviceAccount:${google_service_account.service_account.email}"
  dynamic "condition" {
    for_each = var.in_project_conditions
    content {
      expression  = condition.value.expression
      description = condition.value.description
      title       = condition.value.title
    }
  }
}

resource "google_project_iam_member" "cross_project_iam_role_memberships" {
  for_each = toset(local.cross_project_iam_role_memberships)

  member  = "serviceAccount:${google_service_account.service_account.email}"
  project = local.cross_project_iam_role_membership_project_id
  role    = each.value

  dynamic "condition" {
    for_each = var.cross_project_conditions
    content {
      expression  = condition.value.expression
      description = condition.value.description
      title       = condition.value.title
    }
  }
}

resource "google_folder_iam_member" "folder_iam_role_memberships" {
  for_each = toset(local.folder_iam_role_memberships)

  folder = local.folder_iam_role_membership_folder_id
  member = "serviceAccount:${google_service_account.service_account.email}"
  role   = each.value

  dynamic "condition" {
    for_each = var.folder_conditions
    content {
      expression  = condition.value.expression
      description = condition.value.description
      title       = condition.value.title
    }
  }
}
