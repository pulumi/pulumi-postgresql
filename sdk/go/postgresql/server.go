// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The “Server“ resource creates and manages a foreign server on a PostgreSQL server.
//
// ## Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			extPostgresFdw, err := postgresql.NewExtension(ctx, "extPostgresFdw", nil)
//			if err != nil {
//				return err
//			}
//			_, err = postgresql.NewServer(ctx, "myserverPostgres", &postgresql.ServerArgs{
//				ServerName: pulumi.String("myserver_postgres"),
//				FdwName:    pulumi.String("postgres_fdw"),
//				Options: pulumi.StringMap{
//					"host":   pulumi.String("foo"),
//					"dbname": pulumi.String("foodb"),
//					"port":   pulumi.String("5432"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				extPostgresFdw,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			extFileFdw, err := postgresql.NewExtension(ctx, "extFileFdw", nil)
//			if err != nil {
//				return err
//			}
//			_, err = postgresql.NewServer(ctx, "myserverFile", &postgresql.ServerArgs{
//				ServerName: pulumi.String("myserver_file"),
//				FdwName:    pulumi.String("file_fdw"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				extFileFdw,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Server struct {
	pulumi.CustomResourceState

	// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
	DropCascade pulumi.BoolPtrOutput `pulumi:"dropCascade"`
	// The name of the foreign-data wrapper that manages the server.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	FdwName pulumi.StringOutput `pulumi:"fdwName"`
	// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// The name of the foreign server to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
	ServerOwner pulumi.StringOutput `pulumi:"serverOwner"`
	// Optional server type, potentially useful to foreign-data wrappers.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	ServerType pulumi.StringPtrOutput `pulumi:"serverType"`
	// Optional server version, potentially useful to foreign-data wrappers.
	ServerVersion pulumi.StringPtrOutput `pulumi:"serverVersion"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FdwName == nil {
		return nil, errors.New("invalid value for required argument 'FdwName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("postgresql:index/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("postgresql:index/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
	DropCascade *bool `pulumi:"dropCascade"`
	// The name of the foreign-data wrapper that manages the server.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	FdwName *string `pulumi:"fdwName"`
	// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
	Options map[string]string `pulumi:"options"`
	// The name of the foreign server to be created.
	ServerName *string `pulumi:"serverName"`
	// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
	ServerOwner *string `pulumi:"serverOwner"`
	// Optional server type, potentially useful to foreign-data wrappers.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	ServerType *string `pulumi:"serverType"`
	// Optional server version, potentially useful to foreign-data wrappers.
	ServerVersion *string `pulumi:"serverVersion"`
}

type ServerState struct {
	// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
	DropCascade pulumi.BoolPtrInput
	// The name of the foreign-data wrapper that manages the server.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	FdwName pulumi.StringPtrInput
	// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
	Options pulumi.StringMapInput
	// The name of the foreign server to be created.
	ServerName pulumi.StringPtrInput
	// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
	ServerOwner pulumi.StringPtrInput
	// Optional server type, potentially useful to foreign-data wrappers.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	ServerType pulumi.StringPtrInput
	// Optional server version, potentially useful to foreign-data wrappers.
	ServerVersion pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
	DropCascade *bool `pulumi:"dropCascade"`
	// The name of the foreign-data wrapper that manages the server.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	FdwName string `pulumi:"fdwName"`
	// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
	Options map[string]string `pulumi:"options"`
	// The name of the foreign server to be created.
	ServerName string `pulumi:"serverName"`
	// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
	ServerOwner *string `pulumi:"serverOwner"`
	// Optional server type, potentially useful to foreign-data wrappers.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	ServerType *string `pulumi:"serverType"`
	// Optional server version, potentially useful to foreign-data wrappers.
	ServerVersion *string `pulumi:"serverVersion"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
	DropCascade pulumi.BoolPtrInput
	// The name of the foreign-data wrapper that manages the server.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	FdwName pulumi.StringInput
	// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
	Options pulumi.StringMapInput
	// The name of the foreign server to be created.
	ServerName pulumi.StringInput
	// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
	ServerOwner pulumi.StringPtrInput
	// Optional server type, potentially useful to foreign-data wrappers.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the foreign server is created.
	ServerType pulumi.StringPtrInput
	// Optional server version, potentially useful to foreign-data wrappers.
	ServerVersion pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

func (i *Server) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: i.ToServerOutputWithContext(ctx).OutputState,
	}
}

// ServerArrayInput is an input type that accepts ServerArray and ServerArrayOutput values.
// You can construct a concrete instance of `ServerArrayInput` via:
//
//	ServerArray{ ServerArgs{...} }
type ServerArrayInput interface {
	pulumi.Input

	ToServerArrayOutput() ServerArrayOutput
	ToServerArrayOutputWithContext(context.Context) ServerArrayOutput
}

type ServerArray []ServerInput

func (ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (i ServerArray) ToServerArrayOutput() ServerArrayOutput {
	return i.ToServerArrayOutputWithContext(context.Background())
}

func (i ServerArray) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerArrayOutput)
}

func (i ServerArray) ToOutput(ctx context.Context) pulumix.Output[[]*Server] {
	return pulumix.Output[[]*Server]{
		OutputState: i.ToServerArrayOutputWithContext(ctx).OutputState,
	}
}

// ServerMapInput is an input type that accepts ServerMap and ServerMapOutput values.
// You can construct a concrete instance of `ServerMapInput` via:
//
//	ServerMap{ "key": ServerArgs{...} }
type ServerMapInput interface {
	pulumi.Input

	ToServerMapOutput() ServerMapOutput
	ToServerMapOutputWithContext(context.Context) ServerMapOutput
}

type ServerMap map[string]ServerInput

func (ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (i ServerMap) ToServerMapOutput() ServerMapOutput {
	return i.ToServerMapOutputWithContext(context.Background())
}

func (i ServerMap) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerMapOutput)
}

func (i ServerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Server] {
	return pulumix.Output[map[string]*Server]{
		OutputState: i.ToServerMapOutputWithContext(ctx).OutputState,
	}
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func (o ServerOutput) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: o.OutputState,
	}
}

// When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
func (o ServerOutput) DropCascade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.BoolPtrOutput { return v.DropCascade }).(pulumi.BoolPtrOutput)
}

// The name of the foreign-data wrapper that manages the server.
// Changing this value
// will force the creation of a new resource as this value can only be set
// when the foreign server is created.
func (o ServerOutput) FdwName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FdwName }).(pulumi.StringOutput)
}

// This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
func (o ServerOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// The name of the foreign server to be created.
func (o ServerOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
func (o ServerOutput) ServerOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ServerOwner }).(pulumi.StringOutput)
}

// Optional server type, potentially useful to foreign-data wrappers.
// Changing this value
// will force the creation of a new resource as this value can only be set
// when the foreign server is created.
func (o ServerOutput) ServerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ServerType }).(pulumi.StringPtrOutput)
}

// Optional server version, potentially useful to foreign-data wrappers.
func (o ServerOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ServerArrayOutput struct{ *pulumi.OutputState }

func (ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (o ServerArrayOutput) ToServerArrayOutput() ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Server] {
	return pulumix.Output[[]*Server]{
		OutputState: o.OutputState,
	}
}

func (o ServerArrayOutput) Index(i pulumi.IntInput) ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Server {
		return vs[0].([]*Server)[vs[1].(int)]
	}).(ServerOutput)
}

type ServerMapOutput struct{ *pulumi.OutputState }

func (ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (o ServerMapOutput) ToServerMapOutput() ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Server] {
	return pulumix.Output[map[string]*Server]{
		OutputState: o.OutputState,
	}
}

func (o ServerMapOutput) MapIndex(k pulumi.StringInput) ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Server {
		return vs[0].(map[string]*Server)[vs[1].(string)]
	}).(ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerArrayInput)(nil)).Elem(), ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerMapInput)(nil)).Elem(), ServerMap{})
	pulumi.RegisterOutputType(ServerOutput{})
	pulumi.RegisterOutputType(ServerArrayOutput{})
	pulumi.RegisterOutputType(ServerMapOutput{})
}
