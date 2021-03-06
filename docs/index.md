---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bitbucket Provider"
subcategory: ""
description: |-
  
---

# bitbucket Provider



## Example Usage

```terraform
provider "bitbucket" {
  username = var.username # optionally use BITBUCKET_USERNAME env var
  password = var.password # optionally use BITBUCKET_PASSWORD env var
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **client_id** (String, Sensitive) Client ID for the app accessing the API. Can be specified with the `BITBUCKET_CLIENT_ID` environment variable.
- **client_secret** (String, Sensitive) Client Secret for the app accessing the API. Can be specified with the `BITBUCKET_CLIENT_SECRET` environment variable.
- **password** (String, Sensitive) Password for the user accessing the API. Can be specified with the `BITBUCKET_PASSWORD` environment variable.
- **username** (String) Local user name for the bitbucket API. Can be specified with the `BITBUCKET_USERNAME` environment variable.
