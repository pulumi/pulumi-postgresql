// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections pulumi.BoolPtrOutput `pulumi:"allowConnections"`
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit pulumi.IntPtrOutput `pulumi:"connectionLimit"`
	// Character set encoding to use in the new database
	Encoding pulumi.StringOutput `pulumi:"encoding"`
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate pulumi.BoolOutput `pulumi:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate pulumi.StringOutput `pulumi:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype pulumi.StringOutput `pulumi:"lcCtype"`
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringOutput `pulumi:"name"`
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName pulumi.StringOutput `pulumi:"tablespaceName"`
	// The name of the template from which to create the new database
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("postgresql:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("postgresql:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections *bool `pulumi:"allowConnections"`
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Character set encoding to use in the new database
	Encoding *string `pulumi:"encoding"`
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate *bool `pulumi:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate *string `pulumi:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype *string `pulumi:"lcCtype"`
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name *string `pulumi:"name"`
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner *string `pulumi:"owner"`
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName *string `pulumi:"tablespaceName"`
	// The name of the template from which to create the new database
	Template *string `pulumi:"template"`
}

type DatabaseState struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections pulumi.BoolPtrInput
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit pulumi.IntPtrInput
	// Character set encoding to use in the new database
	Encoding pulumi.StringPtrInput
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate pulumi.BoolPtrInput
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate pulumi.StringPtrInput
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype pulumi.StringPtrInput
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringPtrInput
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner pulumi.StringPtrInput
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName pulumi.StringPtrInput
	// The name of the template from which to create the new database
	Template pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections *bool `pulumi:"allowConnections"`
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Character set encoding to use in the new database
	Encoding *string `pulumi:"encoding"`
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate *bool `pulumi:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate *string `pulumi:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype *string `pulumi:"lcCtype"`
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name *string `pulumi:"name"`
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner *string `pulumi:"owner"`
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName *string `pulumi:"tablespaceName"`
	// The name of the template from which to create the new database
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections pulumi.BoolPtrInput
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit pulumi.IntPtrInput
	// Character set encoding to use in the new database
	Encoding pulumi.StringPtrInput
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate pulumi.BoolPtrInput
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate pulumi.StringPtrInput
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype pulumi.StringPtrInput
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringPtrInput
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner pulumi.StringPtrInput
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName pulumi.StringPtrInput
	// The name of the template from which to create the new database
	Template pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil)).Elem()
}

func (i Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct {
	*pulumi.OutputState
}

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseOutput)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
