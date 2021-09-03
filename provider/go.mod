module github.com/pulumi/pulumi-postgresql/provider/v3

go 1.16

require (
	github.com/cyrilgdn/terraform-provider-postgresql v0.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)

replace github.com/cyrilgdn/terraform-provider-postgresql => github.com/pulumi/terraform-provider-postgresql v1.12.1-0.20210901153353-a48f7e688aca
