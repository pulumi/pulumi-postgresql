// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.
//
// > **Note:** This resource needs Postgresql version 9 or above.
//
// ## Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-postgresql/sdk/v2/go/postgresql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := postgresql.NewDefaultPrivileges(ctx, "readOnlyTables", &postgresql.DefaultPrivilegesArgs{
// 			Database:   pulumi.String("test_db"),
// 			ObjectType: pulumi.String("table"),
// 			Owner:      pulumi.String("db_owner"),
// 			Privileges: pulumi.StringArray{
// 				pulumi.String("SELECT"),
// 			},
// 			Role:   pulumi.String("test_role"),
// 			Schema: pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DefaultPrivileges struct {
	pulumi.CustomResourceState

	// The database to grant default privileges for this role.
	Database pulumi.StringOutput `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The list of privileges to apply as default privileges.
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role pulumi.StringOutput `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema pulumi.StringOutput `pulumi:"schema"`
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
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("postgresql:index/defaultPrivileg:DefaultPrivileg"),
		},
	})
	opts = append(opts, aliases)
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
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	ObjectType *string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner *string `pulumi:"owner"`
	// The list of privileges to apply as default privileges.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role *string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema *string `pulumi:"schema"`
}

type DefaultPrivilegesState struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringPtrInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	ObjectType pulumi.StringPtrInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringPtrInput
	// The list of privileges to apply as default privileges.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringPtrInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrInput
}

func (DefaultPrivilegesState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegesState)(nil)).Elem()
}

type defaultPrivilegesArgs struct {
	// The database to grant default privileges for this role.
	Database string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	ObjectType string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner string `pulumi:"owner"`
	// The list of privileges to apply as default privileges.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a DefaultPrivileges resource.
type DefaultPrivilegesArgs struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	ObjectType pulumi.StringInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringInput
	// The list of privileges to apply as default privileges.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringInput
}

func (DefaultPrivilegesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegesArgs)(nil)).Elem()
}

type DefaultPrivilegesInput interface {
	pulumi.Input

	ToDefaultPrivilegesOutput() DefaultPrivilegesOutput
	ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput
}

func (DefaultPrivileges) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrivileges)(nil)).Elem()
}

func (i DefaultPrivileges) ToDefaultPrivilegesOutput() DefaultPrivilegesOutput {
	return i.ToDefaultPrivilegesOutputWithContext(context.Background())
}

func (i DefaultPrivileges) ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegesOutput)
}

type DefaultPrivilegesOutput struct {
	*pulumi.OutputState
}

func (DefaultPrivilegesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrivilegesOutput)(nil)).Elem()
}

func (o DefaultPrivilegesOutput) ToDefaultPrivilegesOutput() DefaultPrivilegesOutput {
	return o
}

func (o DefaultPrivilegesOutput) ToDefaultPrivilegesOutputWithContext(ctx context.Context) DefaultPrivilegesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefaultPrivilegesOutput{})
}
