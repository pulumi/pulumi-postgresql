module github.com/pulumi/pulumi-postgresql/provider/v3

go 1.16

require (
	github.com/cyrilgdn/terraform-provider-postgresql v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.22.1
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
)

replace (
	github.com/cyrilgdn/terraform-provider-postgresql => github.com/pulumi/terraform-provider-postgresql v1.12.1-0.20220405212505-5e79e19618e4
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
)
