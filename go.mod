module github.com/cgroschupp/terraform-provider-bitbucket

go 1.16

replace github.com/ktrysmt/go-bitbucket => github.com/cgroschupp/go-bitbucket v0.9.9-0.20210402124947-b2c18ae479db

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.5.0
	github.com/ktrysmt/go-bitbucket v0.0.0-00010101000000-000000000000
)
