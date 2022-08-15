output "email" {
  description = "The `email` attribute of the `google_service_account` resource."
  value       = google_service_account.service_account.email
}

output "keys" {
  description = "A map of the service account keys created, with each item in the key_alias as a base64 encoded key.  Returns `ERROR` if the `google_service_account_key.keys` resource cannot be accessed."
  sensitive   = true
  value = {
    for map_key, map_value in google_service_account_key.keys :
    map_key => base64decode(map_value.private_key)
  }
}

output "name" {
  value       = google_service_account.service_account.name
  sensitive   = false
  description = "The name of `google_service_account` resource."
}
