---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cosmo_federated_graph Resource - cosmo"
subcategory: ""
description: |-
  Federated Graph Resource for managing federated graphs
---

# cosmo_federated_graph (Resource)

Federated Graph Resource for managing federated graphs

## Example Usage

```terraform
terraform {
  required_providers {
    cosmo = {
      source  = "terraform.local/wundergraph/cosmo"
      version = "0.0.1"
    }
  }
}

resource "cosmo_federated_graph" "test" {
  name        = var.name
  routing_url = var.routing_url
  namespace   = var.namespace
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the federated graph. This is used to identify the graph and must be unique within the namespace.
- `routing_url` (String) The URL of the service that routes requests to the federated graph.

### Optional

- `admission_webhook_secret` (String, Sensitive) The secret token used to authenticate the admission webhook requests.
- `admission_webhook_url` (String) The URL for the admission webhook that will be triggered during graph operations.
- `label_matchers` (List of String) A list of label matchers used to select the services that will form the federated graph.
- `namespace` (String) The namespace in which the federated graph is located. Defaults to 'default' if not provided.
- `readme` (String) Readme content for the federated graph.

### Read-Only

- `id` (String) The unique identifier of the federated graph resource, automatically generated by the system.
