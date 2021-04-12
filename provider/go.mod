module github.com/pulumi/pulumi-postgresql/provider/v2

go 1.16

require (
	github.com/cyrilgdn/terraform-provider-postgresql v0.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
)

replace github.com/cyrilgdn/terraform-provider-postgresql => github.com/pulumi/terraform-provider-postgresql v1.12.1-0.20210330162646-5e21079b520e
