provider "google" {
  region      = var.google_region
  project     = var.google_project
  credentials = var.google_credentials
}
