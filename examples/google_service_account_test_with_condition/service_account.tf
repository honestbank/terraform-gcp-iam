module "bucket_service_account" {
  source = "../../modules/google_service_account"

  account_id   = "terratest-${var.run_id}"
  display_name = "terratest-${var.run_id}"
  description  = "An instance of the google_service_account Terraform component module."

  in_project_roles         = ["roles/viewer"]
  iam_role_membership_type = "IN_PROJECT"

  in_project_conditions = [
    {
      title       = "User can read readable file"
      description = "User can read readable file"
      expression  = <<EOF
        resource.service == 'storage.googleapis.com' &&
        resource.name == 'projects/_/buckets/${local.bucket_name}/readable.txt'
      EOF
    }
  ]

  project = var.google_project
}
