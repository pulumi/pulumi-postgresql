module github.com/pulumi/pulumi-postgresql/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.19.0
	github.com/pulumi/pulumi/sdk/v2 v2.20.1-0.20210212181059-f4b0fa86fedc
	github.com/terraform-providers/terraform-provider-postgresql v1.7.1
)

replace github.com/terraform-providers/terraform-provider-postgresql => github.com/cyrilgdn/terraform-provider-postgresql v1.11.2
