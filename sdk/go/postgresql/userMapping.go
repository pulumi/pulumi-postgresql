// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “UserMapping“ resource creates and manages a user mapping on a PostgreSQL server.
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
//			extPostgresFdw, err := postgresql.NewExtension(ctx, "ext_postgres_fdw", &postgresql.ExtensionArgs{
//				Name: pulumi.String("postgres_fdw"),
//			})
//			if err != nil {
//				return err
//			}
//			myserverPostgres, err := postgresql.NewServer(ctx, "myserver_postgres", &postgresql.ServerArgs{
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
//			remote, err := postgresql.NewRole(ctx, "remote", &postgresql.RoleArgs{
//				Name: pulumi.String("remote"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = postgresql.NewUserMapping(ctx, "remote", &postgresql.UserMappingArgs{
//				ServerName: myserverPostgres.ServerName,
//				UserName:   remote.Name,
//				Options: pulumi.StringMap{
//					"user":     pulumi.String("admin"),
//					"password": pulumi.String("pass"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type UserMapping struct {
	pulumi.CustomResourceState

	// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// The name of an existing server for which the user mapping is to be created.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUserMapping registers a new resource with the given unique name, arguments, and options.
func NewUserMapping(ctx *pulumi.Context,
	name string, args *UserMappingArgs, opts ...pulumi.ResourceOption) (*UserMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserMapping
	err := ctx.RegisterResource("postgresql:index/userMapping:UserMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserMapping gets an existing UserMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserMappingState, opts ...pulumi.ResourceOption) (*UserMapping, error) {
	var resource UserMapping
	err := ctx.ReadResource("postgresql:index/userMapping:UserMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserMapping resources.
type userMappingState struct {
	// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
	Options map[string]string `pulumi:"options"`
	// The name of an existing server for which the user mapping is to be created.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	ServerName *string `pulumi:"serverName"`
	// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	UserName *string `pulumi:"userName"`
}

type UserMappingState struct {
	// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
	Options pulumi.StringMapInput
	// The name of an existing server for which the user mapping is to be created.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	ServerName pulumi.StringPtrInput
	// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	UserName pulumi.StringPtrInput
}

func (UserMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*userMappingState)(nil)).Elem()
}

type userMappingArgs struct {
	// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
	Options map[string]string `pulumi:"options"`
	// The name of an existing server for which the user mapping is to be created.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	ServerName string `pulumi:"serverName"`
	// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a UserMapping resource.
type UserMappingArgs struct {
	// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
	Options pulumi.StringMapInput
	// The name of an existing server for which the user mapping is to be created.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	ServerName pulumi.StringInput
	// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
	// Changing this value
	// will force the creation of a new resource as this value can only be set
	// when the user mapping is created.
	UserName pulumi.StringInput
}

func (UserMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userMappingArgs)(nil)).Elem()
}

type UserMappingInput interface {
	pulumi.Input

	ToUserMappingOutput() UserMappingOutput
	ToUserMappingOutputWithContext(ctx context.Context) UserMappingOutput
}

func (*UserMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**UserMapping)(nil)).Elem()
}

func (i *UserMapping) ToUserMappingOutput() UserMappingOutput {
	return i.ToUserMappingOutputWithContext(context.Background())
}

func (i *UserMapping) ToUserMappingOutputWithContext(ctx context.Context) UserMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMappingOutput)
}

// UserMappingArrayInput is an input type that accepts UserMappingArray and UserMappingArrayOutput values.
// You can construct a concrete instance of `UserMappingArrayInput` via:
//
//	UserMappingArray{ UserMappingArgs{...} }
type UserMappingArrayInput interface {
	pulumi.Input

	ToUserMappingArrayOutput() UserMappingArrayOutput
	ToUserMappingArrayOutputWithContext(context.Context) UserMappingArrayOutput
}

type UserMappingArray []UserMappingInput

func (UserMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserMapping)(nil)).Elem()
}

func (i UserMappingArray) ToUserMappingArrayOutput() UserMappingArrayOutput {
	return i.ToUserMappingArrayOutputWithContext(context.Background())
}

func (i UserMappingArray) ToUserMappingArrayOutputWithContext(ctx context.Context) UserMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMappingArrayOutput)
}

// UserMappingMapInput is an input type that accepts UserMappingMap and UserMappingMapOutput values.
// You can construct a concrete instance of `UserMappingMapInput` via:
//
//	UserMappingMap{ "key": UserMappingArgs{...} }
type UserMappingMapInput interface {
	pulumi.Input

	ToUserMappingMapOutput() UserMappingMapOutput
	ToUserMappingMapOutputWithContext(context.Context) UserMappingMapOutput
}

type UserMappingMap map[string]UserMappingInput

func (UserMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserMapping)(nil)).Elem()
}

func (i UserMappingMap) ToUserMappingMapOutput() UserMappingMapOutput {
	return i.ToUserMappingMapOutputWithContext(context.Background())
}

func (i UserMappingMap) ToUserMappingMapOutputWithContext(ctx context.Context) UserMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMappingMapOutput)
}

type UserMappingOutput struct{ *pulumi.OutputState }

func (UserMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserMapping)(nil)).Elem()
}

func (o UserMappingOutput) ToUserMappingOutput() UserMappingOutput {
	return o
}

func (o UserMappingOutput) ToUserMappingOutputWithContext(ctx context.Context) UserMappingOutput {
	return o
}

// This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
func (o UserMappingOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UserMapping) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// The name of an existing server for which the user mapping is to be created.
// Changing this value
// will force the creation of a new resource as this value can only be set
// when the user mapping is created.
func (o UserMappingOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserMapping) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
// Changing this value
// will force the creation of a new resource as this value can only be set
// when the user mapping is created.
func (o UserMappingOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserMapping) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type UserMappingArrayOutput struct{ *pulumi.OutputState }

func (UserMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserMapping)(nil)).Elem()
}

func (o UserMappingArrayOutput) ToUserMappingArrayOutput() UserMappingArrayOutput {
	return o
}

func (o UserMappingArrayOutput) ToUserMappingArrayOutputWithContext(ctx context.Context) UserMappingArrayOutput {
	return o
}

func (o UserMappingArrayOutput) Index(i pulumi.IntInput) UserMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserMapping {
		return vs[0].([]*UserMapping)[vs[1].(int)]
	}).(UserMappingOutput)
}

type UserMappingMapOutput struct{ *pulumi.OutputState }

func (UserMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserMapping)(nil)).Elem()
}

func (o UserMappingMapOutput) ToUserMappingMapOutput() UserMappingMapOutput {
	return o
}

func (o UserMappingMapOutput) ToUserMappingMapOutputWithContext(ctx context.Context) UserMappingMapOutput {
	return o
}

func (o UserMappingMapOutput) MapIndex(k pulumi.StringInput) UserMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserMapping {
		return vs[0].(map[string]*UserMapping)[vs[1].(string)]
	}).(UserMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserMappingInput)(nil)).Elem(), &UserMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMappingArrayInput)(nil)).Elem(), UserMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMappingMapInput)(nil)).Elem(), UserMappingMap{})
	pulumi.RegisterOutputType(UserMappingOutput{})
	pulumi.RegisterOutputType(UserMappingArrayOutput{})
	pulumi.RegisterOutputType(UserMappingMapOutput{})
}
