output "account_email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = google_service_account.service_account.email
}

output "service_account_keys" {
  sensitive = true
  value     = try(google_service_account_key.keys, "ERROR")
}
