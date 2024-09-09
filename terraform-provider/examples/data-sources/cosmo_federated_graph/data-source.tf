terraform {
  required_providers {
    cosmo = {
      source  = "terraform.local/wundergraph/cosmo"
      version = "0.0.1"
    }
  }
}

data "cosmo_federated_graph" "test" {
  name      = var.name
  namespace = var.namespace
}