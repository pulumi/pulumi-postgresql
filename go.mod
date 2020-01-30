module github.com/pulumi/pulumi-postgresql

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.5
	github.com/terraform-providers/terraform-provider-postgresql v1.4.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
