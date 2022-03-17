output "account_email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = module.google_service_account_instance.account_email
}

output "service_account_key_private_keys_b64" {
  sensitive = true
  value     = module.google_service_account_instance.service_account_keys_b64
}
