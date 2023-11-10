variable "project" {
}
variable "instance_name" {
  default = "my-instance"
}
variable "machine_type" {
  default = "n2-standard-2"
}
variable "zone" {
  default = "us-central1-a"
}
variable "tags" {
  default = ["web-app"]
}
variable "boot_image" {
  default = "ubuntu-2204-jammy-v20231030"
}
variable "network_interface" {
  default = "default"
}
variable "firewall_rule_name" {
  default = "web-app"
}
variable "allow_ports" {
  default = ["80"]
}
variable "allow_protocol" {
  default = "tcp"
}
variable "sandbox_id" {
}