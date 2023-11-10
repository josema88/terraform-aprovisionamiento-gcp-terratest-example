resource "google_compute_instance" "default" {
  project      = var.project
  name         = "${var.instance_name}-${var.sandbox_id}"
  machine_type = var.machine_type
  zone         = var.zone

  tags = var.tags

  boot_disk {
    initialize_params {
      image = var.boot_image
    }
  }

  network_interface {
    network = var.network_interface
    access_config {
    }
  }
  metadata_startup_script = file("${path.module}/provisioning.sh")
}
resource "google_compute_firewall" "default" {
  project = var.project
  name    = "${var.firewall_rule_name}-${var.sandbox_id}"
  network = var.network_interface

  allow {
    protocol = var.allow_protocol
    ports    = var.allow_ports
  }

  target_tags = var.tags
  source_tags = []
}