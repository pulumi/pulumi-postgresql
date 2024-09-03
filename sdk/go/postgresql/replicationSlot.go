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

// layout: "postgresql"
// page_title: "PostgreSQL: ReplicationSlot"
// sidebar_current: "docs-postgresql-resource-postgresql_replication_slot"
// description: |-
// Creates and manages a replication slot on a PostgreSQL server.
// <!-- yaml: line 6: could not find expected ':' -->
//
// # postgresql\_replication\_slot
//
// The `ReplicationSlot` resource creates and manages a replication slot on a PostgreSQL
// server.
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
//			_, err := postgresql.NewReplicationSlot(ctx, "my_slot", &postgresql.ReplicationSlotArgs{
//				Name:   pulumi.String("my_slot"),
//				Plugin: pulumi.String("test_decoding"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ReplicationSlot struct {
	pulumi.CustomResourceState

	// Which database to create the replication slot on. Defaults to provider database.
	Database pulumi.StringOutput `pulumi:"database"`
	// The name of the replication slot.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the output plugin.
	Plugin pulumi.StringOutput `pulumi:"plugin"`
}

// NewReplicationSlot registers a new resource with the given unique name, arguments, and options.
func NewReplicationSlot(ctx *pulumi.Context,
	name string, args *ReplicationSlotArgs, opts ...pulumi.ResourceOption) (*ReplicationSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plugin == nil {
		return nil, errors.New("invalid value for required argument 'Plugin'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicationSlot
	err := ctx.RegisterResource("postgresql:index/replicationSlot:ReplicationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSlot gets an existing ReplicationSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSlotState, opts ...pulumi.ResourceOption) (*ReplicationSlot, error) {
	var resource ReplicationSlot
	err := ctx.ReadResource("postgresql:index/replicationSlot:ReplicationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSlot resources.
type replicationSlotState struct {
	// Which database to create the replication slot on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the replication slot.
	Name *string `pulumi:"name"`
	// Sets the output plugin.
	Plugin *string `pulumi:"plugin"`
}

type ReplicationSlotState struct {
	// Which database to create the replication slot on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the replication slot.
	Name pulumi.StringPtrInput
	// Sets the output plugin.
	Plugin pulumi.StringPtrInput
}

func (ReplicationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSlotState)(nil)).Elem()
}

type replicationSlotArgs struct {
	// Which database to create the replication slot on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the replication slot.
	Name *string `pulumi:"name"`
	// Sets the output plugin.
	Plugin string `pulumi:"plugin"`
}

// The set of arguments for constructing a ReplicationSlot resource.
type ReplicationSlotArgs struct {
	// Which database to create the replication slot on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the replication slot.
	Name pulumi.StringPtrInput
	// Sets the output plugin.
	Plugin pulumi.StringInput
}

func (ReplicationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSlotArgs)(nil)).Elem()
}

type ReplicationSlotInput interface {
	pulumi.Input

	ToReplicationSlotOutput() ReplicationSlotOutput
	ToReplicationSlotOutputWithContext(ctx context.Context) ReplicationSlotOutput
}

func (*ReplicationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSlot)(nil)).Elem()
}

func (i *ReplicationSlot) ToReplicationSlotOutput() ReplicationSlotOutput {
	return i.ToReplicationSlotOutputWithContext(context.Background())
}

func (i *ReplicationSlot) ToReplicationSlotOutputWithContext(ctx context.Context) ReplicationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSlotOutput)
}

// ReplicationSlotArrayInput is an input type that accepts ReplicationSlotArray and ReplicationSlotArrayOutput values.
// You can construct a concrete instance of `ReplicationSlotArrayInput` via:
//
//	ReplicationSlotArray{ ReplicationSlotArgs{...} }
type ReplicationSlotArrayInput interface {
	pulumi.Input

	ToReplicationSlotArrayOutput() ReplicationSlotArrayOutput
	ToReplicationSlotArrayOutputWithContext(context.Context) ReplicationSlotArrayOutput
}

type ReplicationSlotArray []ReplicationSlotInput

func (ReplicationSlotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSlot)(nil)).Elem()
}

func (i ReplicationSlotArray) ToReplicationSlotArrayOutput() ReplicationSlotArrayOutput {
	return i.ToReplicationSlotArrayOutputWithContext(context.Background())
}

func (i ReplicationSlotArray) ToReplicationSlotArrayOutputWithContext(ctx context.Context) ReplicationSlotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSlotArrayOutput)
}

// ReplicationSlotMapInput is an input type that accepts ReplicationSlotMap and ReplicationSlotMapOutput values.
// You can construct a concrete instance of `ReplicationSlotMapInput` via:
//
//	ReplicationSlotMap{ "key": ReplicationSlotArgs{...} }
type ReplicationSlotMapInput interface {
	pulumi.Input

	ToReplicationSlotMapOutput() ReplicationSlotMapOutput
	ToReplicationSlotMapOutputWithContext(context.Context) ReplicationSlotMapOutput
}

type ReplicationSlotMap map[string]ReplicationSlotInput

func (ReplicationSlotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSlot)(nil)).Elem()
}

func (i ReplicationSlotMap) ToReplicationSlotMapOutput() ReplicationSlotMapOutput {
	return i.ToReplicationSlotMapOutputWithContext(context.Background())
}

func (i ReplicationSlotMap) ToReplicationSlotMapOutputWithContext(ctx context.Context) ReplicationSlotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSlotMapOutput)
}

type ReplicationSlotOutput struct{ *pulumi.OutputState }

func (ReplicationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSlot)(nil)).Elem()
}

func (o ReplicationSlotOutput) ToReplicationSlotOutput() ReplicationSlotOutput {
	return o
}

func (o ReplicationSlotOutput) ToReplicationSlotOutputWithContext(ctx context.Context) ReplicationSlotOutput {
	return o
}

// Which database to create the replication slot on. Defaults to provider database.
func (o ReplicationSlotOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSlot) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The name of the replication slot.
func (o ReplicationSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sets the output plugin.
func (o ReplicationSlotOutput) Plugin() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSlot) pulumi.StringOutput { return v.Plugin }).(pulumi.StringOutput)
}

type ReplicationSlotArrayOutput struct{ *pulumi.OutputState }

func (ReplicationSlotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSlot)(nil)).Elem()
}

func (o ReplicationSlotArrayOutput) ToReplicationSlotArrayOutput() ReplicationSlotArrayOutput {
	return o
}

func (o ReplicationSlotArrayOutput) ToReplicationSlotArrayOutputWithContext(ctx context.Context) ReplicationSlotArrayOutput {
	return o
}

func (o ReplicationSlotArrayOutput) Index(i pulumi.IntInput) ReplicationSlotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationSlot {
		return vs[0].([]*ReplicationSlot)[vs[1].(int)]
	}).(ReplicationSlotOutput)
}

type ReplicationSlotMapOutput struct{ *pulumi.OutputState }

func (ReplicationSlotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSlot)(nil)).Elem()
}

func (o ReplicationSlotMapOutput) ToReplicationSlotMapOutput() ReplicationSlotMapOutput {
	return o
}

func (o ReplicationSlotMapOutput) ToReplicationSlotMapOutputWithContext(ctx context.Context) ReplicationSlotMapOutput {
	return o
}

func (o ReplicationSlotMapOutput) MapIndex(k pulumi.StringInput) ReplicationSlotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationSlot {
		return vs[0].(map[string]*ReplicationSlot)[vs[1].(string)]
	}).(ReplicationSlotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSlotInput)(nil)).Elem(), &ReplicationSlot{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSlotArrayInput)(nil)).Elem(), ReplicationSlotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSlotMapInput)(nil)).Elem(), ReplicationSlotMap{})
	pulumi.RegisterOutputType(ReplicationSlotOutput{})
	pulumi.RegisterOutputType(ReplicationSlotArrayOutput{})
	pulumi.RegisterOutputType(ReplicationSlotMapOutput{})
}
