terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
    google-beta = {
      source = "hashicorp/google-beta"
    }
  }
}

provider "google" {
  region = "us-central1"
  scopes = [
    "cloud-platform"
  ]
}
provider "google-beta" {
  region = "us-central1"
}