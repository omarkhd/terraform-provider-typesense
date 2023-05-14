---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "typesense Provider"
subcategory: ""
description: |-
  Manage your Typesense clusters
---

# typesense Provider

Manage your Typesense clusters

## Example Usage

```terraform
provider "typesense" {
  key = "foobarbaz"  # Or use TYPESENSE_MANAGEMENT_KEY envvar
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `key` (String, Sensitive) Cloud Management API Key