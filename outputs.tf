output "account_email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = module.google_service_account_instance.account_email
}

output "service_account_key_private_keys" {
  sensitive = true
  value     = module.google_service_account_instance.service_account_keys
}

output "service_account_name" {
  value       = google_service_account.service_account.name
  sensitive   = false
  description = "Theh name of `google_service_account` resource."
}
