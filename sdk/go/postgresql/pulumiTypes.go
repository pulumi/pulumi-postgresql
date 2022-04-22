// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderClientcert struct {
	Cert string `pulumi:"cert"`
	Key  string `pulumi:"key"`
}

// ProviderClientcertInput is an input type that accepts ProviderClientcertArgs and ProviderClientcertOutput values.
// You can construct a concrete instance of `ProviderClientcertInput` via:
//
//          ProviderClientcertArgs{...}
type ProviderClientcertInput interface {
	pulumi.Input

	ToProviderClientcertOutput() ProviderClientcertOutput
	ToProviderClientcertOutputWithContext(context.Context) ProviderClientcertOutput
}

type ProviderClientcertArgs struct {
	Cert pulumi.StringInput `pulumi:"cert"`
	Key  pulumi.StringInput `pulumi:"key"`
}

func (ProviderClientcertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderClientcert)(nil)).Elem()
}

func (i ProviderClientcertArgs) ToProviderClientcertOutput() ProviderClientcertOutput {
	return i.ToProviderClientcertOutputWithContext(context.Background())
}

func (i ProviderClientcertArgs) ToProviderClientcertOutputWithContext(ctx context.Context) ProviderClientcertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderClientcertOutput)
}

func (i ProviderClientcertArgs) ToProviderClientcertPtrOutput() ProviderClientcertPtrOutput {
	return i.ToProviderClientcertPtrOutputWithContext(context.Background())
}

func (i ProviderClientcertArgs) ToProviderClientcertPtrOutputWithContext(ctx context.Context) ProviderClientcertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderClientcertOutput).ToProviderClientcertPtrOutputWithContext(ctx)
}

// ProviderClientcertPtrInput is an input type that accepts ProviderClientcertArgs, ProviderClientcertPtr and ProviderClientcertPtrOutput values.
// You can construct a concrete instance of `ProviderClientcertPtrInput` via:
//
//          ProviderClientcertArgs{...}
//
//  or:
//
//          nil
type ProviderClientcertPtrInput interface {
	pulumi.Input

	ToProviderClientcertPtrOutput() ProviderClientcertPtrOutput
	ToProviderClientcertPtrOutputWithContext(context.Context) ProviderClientcertPtrOutput
}

type providerClientcertPtrType ProviderClientcertArgs

func ProviderClientcertPtr(v *ProviderClientcertArgs) ProviderClientcertPtrInput {
	return (*providerClientcertPtrType)(v)
}

func (*providerClientcertPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderClientcert)(nil)).Elem()
}

func (i *providerClientcertPtrType) ToProviderClientcertPtrOutput() ProviderClientcertPtrOutput {
	return i.ToProviderClientcertPtrOutputWithContext(context.Background())
}

func (i *providerClientcertPtrType) ToProviderClientcertPtrOutputWithContext(ctx context.Context) ProviderClientcertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderClientcertPtrOutput)
}

type ProviderClientcertOutput struct{ *pulumi.OutputState }

func (ProviderClientcertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderClientcert)(nil)).Elem()
}

func (o ProviderClientcertOutput) ToProviderClientcertOutput() ProviderClientcertOutput {
	return o
}

func (o ProviderClientcertOutput) ToProviderClientcertOutputWithContext(ctx context.Context) ProviderClientcertOutput {
	return o
}

func (o ProviderClientcertOutput) ToProviderClientcertPtrOutput() ProviderClientcertPtrOutput {
	return o.ToProviderClientcertPtrOutputWithContext(context.Background())
}

func (o ProviderClientcertOutput) ToProviderClientcertPtrOutputWithContext(ctx context.Context) ProviderClientcertPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderClientcert) *ProviderClientcert {
		return &v
	}).(ProviderClientcertPtrOutput)
}

func (o ProviderClientcertOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderClientcert) string { return v.Cert }).(pulumi.StringOutput)
}

func (o ProviderClientcertOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderClientcert) string { return v.Key }).(pulumi.StringOutput)
}

type ProviderClientcertPtrOutput struct{ *pulumi.OutputState }

func (ProviderClientcertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderClientcert)(nil)).Elem()
}

func (o ProviderClientcertPtrOutput) ToProviderClientcertPtrOutput() ProviderClientcertPtrOutput {
	return o
}

func (o ProviderClientcertPtrOutput) ToProviderClientcertPtrOutputWithContext(ctx context.Context) ProviderClientcertPtrOutput {
	return o
}

func (o ProviderClientcertPtrOutput) Elem() ProviderClientcertOutput {
	return o.ApplyT(func(v *ProviderClientcert) ProviderClientcert {
		if v != nil {
			return *v
		}
		var ret ProviderClientcert
		return ret
	}).(ProviderClientcertOutput)
}

func (o ProviderClientcertPtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderClientcert) *string {
		if v == nil {
			return nil
		}
		return &v.Cert
	}).(pulumi.StringPtrOutput)
}

func (o ProviderClientcertPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderClientcert) *string {
		if v == nil {
			return nil
		}
		return &v.Key
	}).(pulumi.StringPtrOutput)
}

type SchemaPolicy struct {
	// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
	Create *bool `pulumi:"create"`
	// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
	CreateWithGrant *bool `pulumi:"createWithGrant"`
	// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
	Role *string `pulumi:"role"`
	// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
	Usage *bool `pulumi:"usage"`
	// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
	UsageWithGrant *bool `pulumi:"usageWithGrant"`
}

// SchemaPolicyInput is an input type that accepts SchemaPolicyArgs and SchemaPolicyOutput values.
// You can construct a concrete instance of `SchemaPolicyInput` via:
//
//          SchemaPolicyArgs{...}
type SchemaPolicyInput interface {
	pulumi.Input

	ToSchemaPolicyOutput() SchemaPolicyOutput
	ToSchemaPolicyOutputWithContext(context.Context) SchemaPolicyOutput
}

type SchemaPolicyArgs struct {
	// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
	Create pulumi.BoolPtrInput `pulumi:"create"`
	// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
	CreateWithGrant pulumi.BoolPtrInput `pulumi:"createWithGrant"`
	// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
	Role pulumi.StringPtrInput `pulumi:"role"`
	// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
	Usage pulumi.BoolPtrInput `pulumi:"usage"`
	// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
	UsageWithGrant pulumi.BoolPtrInput `pulumi:"usageWithGrant"`
}

func (SchemaPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaPolicy)(nil)).Elem()
}

func (i SchemaPolicyArgs) ToSchemaPolicyOutput() SchemaPolicyOutput {
	return i.ToSchemaPolicyOutputWithContext(context.Background())
}

func (i SchemaPolicyArgs) ToSchemaPolicyOutputWithContext(ctx context.Context) SchemaPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPolicyOutput)
}

// SchemaPolicyArrayInput is an input type that accepts SchemaPolicyArray and SchemaPolicyArrayOutput values.
// You can construct a concrete instance of `SchemaPolicyArrayInput` via:
//
//          SchemaPolicyArray{ SchemaPolicyArgs{...} }
type SchemaPolicyArrayInput interface {
	pulumi.Input

	ToSchemaPolicyArrayOutput() SchemaPolicyArrayOutput
	ToSchemaPolicyArrayOutputWithContext(context.Context) SchemaPolicyArrayOutput
}

type SchemaPolicyArray []SchemaPolicyInput

func (SchemaPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchemaPolicy)(nil)).Elem()
}

func (i SchemaPolicyArray) ToSchemaPolicyArrayOutput() SchemaPolicyArrayOutput {
	return i.ToSchemaPolicyArrayOutputWithContext(context.Background())
}

func (i SchemaPolicyArray) ToSchemaPolicyArrayOutputWithContext(ctx context.Context) SchemaPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPolicyArrayOutput)
}

type SchemaPolicyOutput struct{ *pulumi.OutputState }

func (SchemaPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaPolicy)(nil)).Elem()
}

func (o SchemaPolicyOutput) ToSchemaPolicyOutput() SchemaPolicyOutput {
	return o
}

func (o SchemaPolicyOutput) ToSchemaPolicyOutputWithContext(ctx context.Context) SchemaPolicyOutput {
	return o
}

// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
func (o SchemaPolicyOutput) Create() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchemaPolicy) *bool { return v.Create }).(pulumi.BoolPtrOutput)
}

// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
func (o SchemaPolicyOutput) CreateWithGrant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchemaPolicy) *bool { return v.CreateWithGrant }).(pulumi.BoolPtrOutput)
}

// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
func (o SchemaPolicyOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchemaPolicy) *string { return v.Role }).(pulumi.StringPtrOutput)
}

// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
func (o SchemaPolicyOutput) Usage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchemaPolicy) *bool { return v.Usage }).(pulumi.BoolPtrOutput)
}

// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
func (o SchemaPolicyOutput) UsageWithGrant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchemaPolicy) *bool { return v.UsageWithGrant }).(pulumi.BoolPtrOutput)
}

type SchemaPolicyArrayOutput struct{ *pulumi.OutputState }

func (SchemaPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchemaPolicy)(nil)).Elem()
}

func (o SchemaPolicyArrayOutput) ToSchemaPolicyArrayOutput() SchemaPolicyArrayOutput {
	return o
}

func (o SchemaPolicyArrayOutput) ToSchemaPolicyArrayOutputWithContext(ctx context.Context) SchemaPolicyArrayOutput {
	return o
}

func (o SchemaPolicyArrayOutput) Index(i pulumi.IntInput) SchemaPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SchemaPolicy {
		return vs[0].([]SchemaPolicy)[vs[1].(int)]
	}).(SchemaPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderClientcertInput)(nil)).Elem(), ProviderClientcertArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderClientcertPtrInput)(nil)).Elem(), ProviderClientcertArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPolicyInput)(nil)).Elem(), SchemaPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPolicyArrayInput)(nil)).Elem(), SchemaPolicyArray{})
	pulumi.RegisterOutputType(ProviderClientcertOutput{})
	pulumi.RegisterOutputType(ProviderClientcertPtrOutput{})
	pulumi.RegisterOutputType(SchemaPolicyOutput{})
	pulumi.RegisterOutputType(SchemaPolicyArrayOutput{})
}
