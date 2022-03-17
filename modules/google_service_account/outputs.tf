output "account_email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = google_service_account.service_account.email
}

output "service_account_keys_b64" {
  description = "A map of the service account keys created, with each item in the key_alias as a key. Values are base64 encoded. Returns `ERROR` if the `google_service_account_key.keys` resource cannot be accessed."
  sensitive   = true
  value       = try(google_service_account_key.keys, "ERROR")
}
