module github.com/pulumi/pulumi-postgresql/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.18.1
	github.com/pulumi/pulumi/sdk/v2 v2.18.0
	github.com/terraform-providers/terraform-provider-postgresql v1.7.1
)

replace github.com/terraform-providers/terraform-provider-postgresql => github.com/cyrilgdn/terraform-provider-postgresql v1.11.0
