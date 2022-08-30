output "keys" {
  value     = module.service_account.keys
  sensitive = true
}

output "service_account_email" {
  value = module.service_account.email
}
