variable "new_relic_license_key" {
  description = "New Relic License Key"
  type        = string
  sensitive   = true
  default     = "unintialized"
}

variable "new_relic_user_key" {
  description = "New Relic User Key"
  type        = string
  sensitive   = true
  default     = "unintialized"
}