// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the postgresql package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The name of the database to connect to in order to conenct to (defaults to `postgres`).
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// Database username associated to the connected user (for user name maps)
	DatabaseUsername pulumi.StringPtrOutput `pulumi:"databaseUsername"`
	// Specify the expected version of PostgreSQL.
	ExpectedVersion pulumi.StringPtrOutput `pulumi:"expectedVersion"`
	// Name of PostgreSQL server address to connect to
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// Password to be used if the PostgreSQL server demands password authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	Scheme   pulumi.StringPtrOutput `pulumi:"scheme"`
	// Deprecated: Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
	SslMode pulumi.StringPtrOutput `pulumi:"sslMode"`
	// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
	// PostgreSQL server
	Sslmode pulumi.StringPtrOutput `pulumi:"sslmode"`
	// The SSL server root certificate file path. The file must contain PEM encoded data.
	Sslrootcert pulumi.StringPtrOutput `pulumi:"sslrootcert"`
	// PostgreSQL user name to connect as
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.ConnectTimeout) {
		args.ConnectTimeout = pulumi.IntPtr(getEnvOrDefault(180, parseEnvInt, "PGCONNECT_TIMEOUT").(int))
	}
	if isZero(args.Sslmode) {
		args.Sslmode = pulumi.StringPtr(getEnvOrDefault("", nil, "PGSSLMODE").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:postgresql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// SSL client certificate if required by the database.
	Clientcert *ProviderClientcert `pulumi:"clientcert"`
	// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// The name of the database to connect to in order to conenct to (defaults to `postgres`).
	Database *string `pulumi:"database"`
	// Database username associated to the connected user (for user name maps)
	DatabaseUsername *string `pulumi:"databaseUsername"`
	// Specify the expected version of PostgreSQL.
	ExpectedVersion *string `pulumi:"expectedVersion"`
	// Name of PostgreSQL server address to connect to
	Host *string `pulumi:"host"`
	// Maximum number of connections to establish to the database. Zero means unlimited.
	MaxConnections *int `pulumi:"maxConnections"`
	// Password to be used if the PostgreSQL server demands password authentication
	Password *string `pulumi:"password"`
	// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
	Port   *int    `pulumi:"port"`
	Scheme *string `pulumi:"scheme"`
	// Deprecated: Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
	SslMode *string `pulumi:"sslMode"`
	// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
	// PostgreSQL server
	Sslmode *string `pulumi:"sslmode"`
	// The SSL server root certificate file path. The file must contain PEM encoded data.
	Sslrootcert *string `pulumi:"sslrootcert"`
	// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
	// Refreshing state password from Postgres)
	Superuser *bool `pulumi:"superuser"`
	// PostgreSQL user name to connect as
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// SSL client certificate if required by the database.
	Clientcert ProviderClientcertPtrInput
	// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
	ConnectTimeout pulumi.IntPtrInput
	// The name of the database to connect to in order to conenct to (defaults to `postgres`).
	Database pulumi.StringPtrInput
	// Database username associated to the connected user (for user name maps)
	DatabaseUsername pulumi.StringPtrInput
	// Specify the expected version of PostgreSQL.
	ExpectedVersion pulumi.StringPtrInput
	// Name of PostgreSQL server address to connect to
	Host pulumi.StringPtrInput
	// Maximum number of connections to establish to the database. Zero means unlimited.
	MaxConnections pulumi.IntPtrInput
	// Password to be used if the PostgreSQL server demands password authentication
	Password pulumi.StringPtrInput
	// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
	Port   pulumi.IntPtrInput
	Scheme pulumi.StringPtrInput
	// Deprecated: Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
	SslMode pulumi.StringPtrInput
	// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
	// PostgreSQL server
	Sslmode pulumi.StringPtrInput
	// The SSL server root certificate file path. The file must contain PEM encoded data.
	Sslrootcert pulumi.StringPtrInput
	// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
	// Refreshing state password from Postgres)
	Superuser pulumi.BoolPtrInput
	// PostgreSQL user name to connect as
	Username pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
