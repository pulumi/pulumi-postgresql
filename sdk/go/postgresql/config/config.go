// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Use rdsIam instead of password authentication (see:
// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
func GetAwsRdsIamAuth(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "postgresql:awsRdsIamAuth")
}

// AWS profile to use for IAM auth
func GetAwsRdsIamProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:awsRdsIamProfile")
}

// AWS IAM role to assume for IAM auth
func GetAwsRdsIamProviderRoleArn(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:awsRdsIamProviderRoleArn")
}

// AWS region to use for IAM auth
func GetAwsRdsIamRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:awsRdsIamRegion")
}

// Use MS Azure identity OAuth token (see:
// https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
func GetAzureIdentityAuth(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "postgresql:azureIdentityAuth")
}
func GetAzureTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:azureTenantId")
}

// SSL client certificate if required by the database.
func GetClientcert(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:clientcert")
}

// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
func GetConnectTimeout(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "postgresql:connectTimeout")
	if err == nil {
		return v
	}
	var value int
	if d := internal.GetEnvOrDefault(180, internal.ParseEnvInt, "PGCONNECT_TIMEOUT"); d != nil {
		value = d.(int)
	}
	return value
}

// The name of the database to connect to in order to conenct to (defaults to `postgres`).
func GetDatabase(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:database")
}

// Database username associated to the connected user (for user name maps)
func GetDatabaseUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:databaseUsername")
}

// Specify the expected version of PostgreSQL.
func GetExpectedVersion(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:expectedVersion")
}

// Service account to impersonate when using GCP IAM authentication.
func GetGcpIamImpersonateServiceAccount(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:gcpIamImpersonateServiceAccount")
}

// Name of PostgreSQL server address to connect to
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:host")
}

// Maximum number of connections to establish to the database. Zero means unlimited.
func GetMaxConnections(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "postgresql:maxConnections")
}

// Password to be used if the PostgreSQL server demands password authentication
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:password")
}

// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
func GetPort(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "postgresql:port")
}
func GetScheme(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:scheme")
}

// Deprecated: Rename PostgreSQL provider `sslMode` attribute to `sslmode`
func GetSslMode(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:sslMode")
}

// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
// PostgreSQL server
func GetSslmode(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "postgresql:sslmode")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "PGSSLMODE"); d != nil {
		value = d.(string)
	}
	return value
}

// The SSL server root certificate file path. The file must contain PEM encoded data.
func GetSslrootcert(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:sslrootcert")
}

// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
// Refreshing state password from Postgres)
func GetSuperuser(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "postgresql:superuser")
}

// PostgreSQL user name to connect as
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:username")
}
