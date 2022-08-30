output "account_email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = module.google_service_account_instance.email
}

output "service_account_key_private_keys" {
  sensitive = true
  value     = module.google_service_account_instance.keys
}
