locals {
  cross_project_iam_memberships = flatten([
    for project_id, memberships in var.cross_project_iam_memberships : [
      for membership in memberships : {
        project_id = project_id
        role       = membership.role
        conditions = membership.conditions
      }
    ]
  ])

  folder_iam_memberships = flatten([
    for folder_id, memberships in var.folder_iam_memberships : [
      for membership in memberships : {
        folder_id  = folder_id
        role       = membership.role
        conditions = membership.conditions
      }
    ]
  ])
}

resource "google_project_iam_member" "project_iam_memberships" {
  for_each = {
    for membership in var.project_iam_memberships :
    membership.role => membership
  }

  member = "serviceAccount:${google_service_account.service_account.email}"

  project = var.project_id
  role    = each.value.role

  dynamic "condition" {
    for_each = each.value.conditions != null ? each.value.conditions : []
    content {
      description = condition.value.description
      expression  = condition.value.expression
      title       = condition.value.title
    }
  }
}

resource "google_project_iam_member" "cross_project_iam_memberships" {
  for_each = {
    for membership in local.cross_project_iam_memberships :
    "${membership.project_id} - ${membership.role}" => membership
  }

  member = "serviceAccount:${google_service_account.service_account.email}"

  project = each.value.project_id
  role    = each.value.role

  dynamic "condition" {
    for_each = each.value.conditions != null ? each.value.conditions : []
    content {
      description = condition.value.description
      expression  = condition.value.expression
      title       = condition.value.title
    }
  }
}

resource "google_folder_iam_member" "folder_iam_memberships" {
  for_each = {
    for membership in local.folder_iam_memberships :
    "${membership.folder_id} - ${membership.role}" => membership
  }

  member = "serviceAccount:${google_service_account.service_account.email}"

  folder = each.value.folder_id
  role   = each.value.role

  dynamic "condition" {
    for_each = each.value.conditions != null ? each.value.conditions : []
    content {
      description = condition.value.description
      expression  = condition.value.expression
      title       = condition.value.title
    }
  }
}
