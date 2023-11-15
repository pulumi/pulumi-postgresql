// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “DefaultPrivileges“ resource creates and manages default privileges given to a user for a database schema.
//
// > **Note:** This resource needs Postgresql version 9 or above.
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
//			_, err := postgresql.NewDefaultPrivileges(ctx, "readOnlyTables", &postgresql.DefaultPrivilegesArgs{
//				Database:   pulumi.String("test_db"),
//				ObjectType: pulumi.String("table"),
//				Owner:      pulumi.String("db_owner"),
//				Privileges: pulumi.StringArray{
//					pulumi.String("SELECT"),
//				},
//				Role:   pulumi.String("test_role"),
//				Schema: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Examples
//
// Revoke default privileges for functions for "public" role:
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
//			_, err := postgresql.NewDefaultPrivileges(ctx, "revokePublic", &postgresql.DefaultPrivilegesArgs{
//				Database:   pulumi.Any(postgresql_database.Example_db.Name),
//				Role:       pulumi.String("public"),
//				Owner:      pulumi.String("object_owner"),
//				ObjectType: pulumi.String("function"),
//				Privileges: pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DefaultPrivileges struct {
	pulumi.CustomResourceState

	// The database to grant default privileges for this role.
	Database pulumi.StringOutput `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role pulumi.StringOutput `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrOutput `pulumi:"withGrantOption"`
}

// NewDefaultPrivileges registers a new resource with the given unique name, arguments, and options.
func NewDefaultPrivileges(ctx *pulumi.Context,
	name string, args *DefaultPrivilegesArgs, opts ...pulumi.ResourceOption) (*DefaultPrivileges, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	if args.Privileges == nil {
		return nil, errors.New("invalid value for required argument 'Privileges'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("postgresql:index/defaultPrivileg:DefaultPrivileg"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultPrivileges
	err := ctx.RegisterResource("postgresql:index/defaultPrivileges:DefaultPrivileges", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultPrivileges gets an existing DefaultPrivileges resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPrivileges(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultPrivilegesState, opts ...pulumi.ResourceOption) (*DefaultPrivileges, error) {
	var resource DefaultPrivileges
	err := ctx.ReadResource("postgresql:index/defaultPrivileges:DefaultPrivileges", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultPrivileges resources.
type defaultPrivilegesState struct {
	// The database to grant default privileges for this role.
	Database *string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType *string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner *string `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role *string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema *string `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type DefaultPrivilegesState struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringPtrInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringPtrInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringPtrInput
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringPtrInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrInput
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrInput
}

func (DefaultPrivilegesState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegesState)(nil)).Elem()
}

type defaultPrivilegesArgs struct {
	// The database to grant default privileges for this role.
	Database string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner string `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema *string `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a DefaultPrivileges resource.
type DefaultPrivilegesArgs struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringInput
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrInput
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrInput
}

func (DefaultPrivilegesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegesArgs)(nil)).Elem()
}

type DefaultPrivilegesInput interface {
	pulumi.Input

	ToDefaultPrivilegesOutput() DefaultPrivilegesOutput
	ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput
}

func (*DefaultPrivileges) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrivileges)(nil)).Elem()
}

func (i *DefaultPrivileges) ToDefaultPrivilegesOutput() DefaultPrivilegesOutput {
	return i.ToDefaultPrivilegesOutputWithContext(context.Background())
}

func (i *DefaultPrivileges) ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegesOutput)
}

// DefaultPrivilegesArrayInput is an input type that accepts DefaultPrivilegesArray and DefaultPrivilegesArrayOutput values.
// You can construct a concrete instance of `DefaultPrivilegesArrayInput` via:
//
//	DefaultPrivilegesArray{ DefaultPrivilegesArgs{...} }
type DefaultPrivilegesArrayInput interface {
	pulumi.Input

	ToDefaultPrivilegesArrayOutput() DefaultPrivilegesArrayOutput
	ToDefaultPrivilegesArrayOutputWithContext(context.Context) DefaultPrivilegesArrayOutput
}

type DefaultPrivilegesArray []DefaultPrivilegesInput

func (DefaultPrivilegesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPrivileges)(nil)).Elem()
}

func (i DefaultPrivilegesArray) ToDefaultPrivilegesArrayOutput() DefaultPrivilegesArrayOutput {
	return i.ToDefaultPrivilegesArrayOutputWithContext(context.Background())
}

func (i DefaultPrivilegesArray) ToDefaultPrivilegesArrayOutputWithContext(ctx context.Context) DefaultPrivilegesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegesArrayOutput)
}

// DefaultPrivilegesMapInput is an input type that accepts DefaultPrivilegesMap and DefaultPrivilegesMapOutput values.
// You can construct a concrete instance of `DefaultPrivilegesMapInput` via:
//
//	DefaultPrivilegesMap{ "key": DefaultPrivilegesArgs{...} }
type DefaultPrivilegesMapInput interface {
	pulumi.Input

	ToDefaultPrivilegesMapOutput() DefaultPrivilegesMapOutput
	ToDefaultPrivilegesMapOutputWithContext(context.Context) DefaultPrivilegesMapOutput
}

type DefaultPrivilegesMap map[string]DefaultPrivilegesInput

func (DefaultPrivilegesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPrivileges)(nil)).Elem()
}

func (i DefaultPrivilegesMap) ToDefaultPrivilegesMapOutput() DefaultPrivilegesMapOutput {
	return i.ToDefaultPrivilegesMapOutputWithContext(context.Background())
}

func (i DefaultPrivilegesMap) ToDefaultPrivilegesMapOutputWithContext(ctx context.Context) DefaultPrivilegesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegesMapOutput)
}

type DefaultPrivilegesOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrivileges)(nil)).Elem()
}

func (o DefaultPrivilegesOutput) ToDefaultPrivilegesOutput() DefaultPrivilegesOutput {
	return o
}

func (o DefaultPrivilegesOutput) ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput {
	return o
}

// The database to grant default privileges for this role.
func (o DefaultPrivilegesOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
func (o DefaultPrivilegesOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringOutput { return v.ObjectType }).(pulumi.StringOutput)
}

// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
func (o DefaultPrivilegesOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
func (o DefaultPrivilegesOutput) Privileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringArrayOutput { return v.Privileges }).(pulumi.StringArrayOutput)
}

// The name of the role to which grant default privileges on.
func (o DefaultPrivilegesOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The database schema to set default privileges for this role.
func (o DefaultPrivilegesOutput) Schema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.StringPtrOutput { return v.Schema }).(pulumi.StringPtrOutput)
}

// Permit the grant recipient to grant it to others
func (o DefaultPrivilegesOutput) WithGrantOption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultPrivileges) pulumi.BoolPtrOutput { return v.WithGrantOption }).(pulumi.BoolPtrOutput)
}

type DefaultPrivilegesArrayOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPrivileges)(nil)).Elem()
}

func (o DefaultPrivilegesArrayOutput) ToDefaultPrivilegesArrayOutput() DefaultPrivilegesArrayOutput {
	return o
}

func (o DefaultPrivilegesArrayOutput) ToDefaultPrivilegesArrayOutputWithContext(ctx context.Context) DefaultPrivilegesArrayOutput {
	return o
}

func (o DefaultPrivilegesArrayOutput) Index(i pulumi.IntInput) DefaultPrivilegesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultPrivileges {
		return vs[0].([]*DefaultPrivileges)[vs[1].(int)]
	}).(DefaultPrivilegesOutput)
}

type DefaultPrivilegesMapOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPrivileges)(nil)).Elem()
}

func (o DefaultPrivilegesMapOutput) ToDefaultPrivilegesMapOutput() DefaultPrivilegesMapOutput {
	return o
}

func (o DefaultPrivilegesMapOutput) ToDefaultPrivilegesMapOutputWithContext(ctx context.Context) DefaultPrivilegesMapOutput {
	return o
}

func (o DefaultPrivilegesMapOutput) MapIndex(k pulumi.StringInput) DefaultPrivilegesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultPrivileges {
		return vs[0].(map[string]*DefaultPrivileges)[vs[1].(string)]
	}).(DefaultPrivilegesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegesInput)(nil)).Elem(), &DefaultPrivileges{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegesArrayInput)(nil)).Elem(), DefaultPrivilegesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegesMapInput)(nil)).Elem(), DefaultPrivilegesMap{})
	pulumi.RegisterOutputType(DefaultPrivilegesOutput{})
	pulumi.RegisterOutputType(DefaultPrivilegesArrayOutput{})
	pulumi.RegisterOutputType(DefaultPrivilegesMapOutput{})
}
