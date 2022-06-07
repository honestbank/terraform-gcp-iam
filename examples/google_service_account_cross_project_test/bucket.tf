locals {
  bucket_name = var.run_id
}

module "test_bucket" {
  source = "../modules/terraform-gcp-gcs/modules/gcp_gcs_bucket"

  location = "${var.google_region}2"
  name     = local.bucket_name

  force_destroy = true
}

resource "google_storage_bucket_object" "readable_file" {
  bucket       = module.test_bucket.name
  name         = "readable.txt"
  content      = "hello world"
  content_type = "text/plain; charset=utf-8"
}

resource "google_storage_bucket_object" "unreadable_file" {
  bucket       = module.test_bucket.name
  name         = "unreadable.txt"
  content      = "goodbye world"
  content_type = "text/plain; charset=utf-8"
}
