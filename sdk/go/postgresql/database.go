// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections pulumi.BoolPtrOutput `pulumi:"allowConnections"`
	// If `true`, the change of the database
	// `owner` will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to `false` (the default), then the previous database `owner` will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	AlterObjectOwnership pulumi.BoolPtrOutput `pulumi:"alterObjectOwnership"`
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

	opts = internal.PkgResourceDefaultOpts(opts)
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
	// If `true`, the change of the database
	// `owner` will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to `false` (the default), then the previous database `owner` will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	AlterObjectOwnership *bool `pulumi:"alterObjectOwnership"`
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
	// If `true`, the change of the database
	// `owner` will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to `false` (the default), then the previous database `owner` will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	AlterObjectOwnership pulumi.BoolPtrInput
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
	// If `true`, the change of the database
	// `owner` will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to `false` (the default), then the previous database `owner` will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	AlterObjectOwnership *bool `pulumi:"alterObjectOwnership"`
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
	// If `true`, the change of the database
	// `owner` will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to `false` (the default), then the previous database `owner` will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	AlterObjectOwnership pulumi.BoolPtrInput
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

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// If `false` then no one can connect to this
// database. The default is `true`, allowing connections (except as restricted by
// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
func (o DatabaseOutput) AllowConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.AllowConnections }).(pulumi.BoolPtrOutput)
}

// If `true`, the change of the database
// `owner` will also include a reassignment of the ownership of preexisting
// objects like tables or sequences from the previous owner to the new one.
// If set to `false` (the default), then the previous database `owner` will still
// hold the ownership of the objects in that database. To alter existing objects in
// the database, you must be a direct or indirect member of the specified role, or
// the username in the provider must be superuser.
func (o DatabaseOutput) AlterObjectOwnership() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.AlterObjectOwnership }).(pulumi.BoolPtrOutput)
}

// How many concurrent connections can be
// established to this database. `-1` (the default) means no limit.
func (o DatabaseOutput) ConnectionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.ConnectionLimit }).(pulumi.IntPtrOutput)
}

// Character set encoding to use in the new database
func (o DatabaseOutput) Encoding() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Encoding }).(pulumi.StringOutput)
}

// If `true`, then this database can be cloned by any
// user with `CREATEDB` privileges; if `false` (the default), then only
// superusers or the owner of the database can clone it.
func (o DatabaseOutput) IsTemplate() pulumi.BoolOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolOutput { return v.IsTemplate }).(pulumi.BoolOutput)
}

// Collation order (LC_COLLATE) to use in the new database
func (o DatabaseOutput) LcCollate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.LcCollate }).(pulumi.StringOutput)
}

// Character classification (LC_CTYPE) to use in the new database
func (o DatabaseOutput) LcCtype() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.LcCtype }).(pulumi.StringOutput)
}

// The name of the database. Must be unique on the PostgreSQL
// server instance where it is configured.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The role name of the user who will own the database, or
// `DEFAULT` to use the default (namely, the user executing the command). To
// create a database owned by another role or to change the owner of an existing
// database, you must be a direct or indirect member of the specified role, or
// the username in the provider is a superuser.
func (o DatabaseOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The name of the tablespace that will be
// associated with the database, or `DEFAULT` to use the template database's
// tablespace.  This tablespace will be the default tablespace used for objects
// created in this database.
func (o DatabaseOutput) TablespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.TablespaceName }).(pulumi.StringOutput)
}

// The name of the template from which to create the new database
func (o DatabaseOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
