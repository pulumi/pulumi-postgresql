// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Publication` resource creates and manages a publication on a PostgreSQL
// server.
//
// ## Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := postgresql.NewPublication(ctx, "publication", &postgresql.PublicationArgs{
//				Tables: pulumi.StringArray{
//					pulumi.String("public.test"),
//					pulumi.String("another_schema.test"),
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
// <!--End PulumiCodeChooser -->
type Publication struct {
	pulumi.CustomResourceState

	// Should be ALL TABLES added to the publication. Defaults to 'false'
	AllTables pulumi.BoolOutput `pulumi:"allTables"`
	// Which database to create the publication on. Defaults to provider database.
	Database pulumi.StringOutput `pulumi:"database"`
	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	DropCascade pulumi.BoolPtrOutput `pulumi:"dropCascade"`
	// The name of the publication.
	Name pulumi.StringOutput `pulumi:"name"`
	// Who owns the publication. Defaults to provider user.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	PublishParams pulumi.StringArrayOutput `pulumi:"publishParams"`
	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	PublishViaPartitionRootParam pulumi.BoolPtrOutput `pulumi:"publishViaPartitionRootParam"`
	// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	Tables pulumi.StringArrayOutput `pulumi:"tables"`
}

// NewPublication registers a new resource with the given unique name, arguments, and options.
func NewPublication(ctx *pulumi.Context,
	name string, args *PublicationArgs, opts ...pulumi.ResourceOption) (*Publication, error) {
	if args == nil {
		args = &PublicationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Publication
	err := ctx.RegisterResource("postgresql:index/publication:Publication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublication gets an existing Publication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicationState, opts ...pulumi.ResourceOption) (*Publication, error) {
	var resource Publication
	err := ctx.ReadResource("postgresql:index/publication:Publication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Publication resources.
type publicationState struct {
	// Should be ALL TABLES added to the publication. Defaults to 'false'
	AllTables *bool `pulumi:"allTables"`
	// Which database to create the publication on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	DropCascade *bool `pulumi:"dropCascade"`
	// The name of the publication.
	Name *string `pulumi:"name"`
	// Who owns the publication. Defaults to provider user.
	Owner *string `pulumi:"owner"`
	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	PublishParams []string `pulumi:"publishParams"`
	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	PublishViaPartitionRootParam *bool `pulumi:"publishViaPartitionRootParam"`
	// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	Tables []string `pulumi:"tables"`
}

type PublicationState struct {
	// Should be ALL TABLES added to the publication. Defaults to 'false'
	AllTables pulumi.BoolPtrInput
	// Which database to create the publication on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	DropCascade pulumi.BoolPtrInput
	// The name of the publication.
	Name pulumi.StringPtrInput
	// Who owns the publication. Defaults to provider user.
	Owner pulumi.StringPtrInput
	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	PublishParams pulumi.StringArrayInput
	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	PublishViaPartitionRootParam pulumi.BoolPtrInput
	// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	Tables pulumi.StringArrayInput
}

func (PublicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicationState)(nil)).Elem()
}

type publicationArgs struct {
	// Should be ALL TABLES added to the publication. Defaults to 'false'
	AllTables *bool `pulumi:"allTables"`
	// Which database to create the publication on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	DropCascade *bool `pulumi:"dropCascade"`
	// The name of the publication.
	Name *string `pulumi:"name"`
	// Who owns the publication. Defaults to provider user.
	Owner *string `pulumi:"owner"`
	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	PublishParams []string `pulumi:"publishParams"`
	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	PublishViaPartitionRootParam *bool `pulumi:"publishViaPartitionRootParam"`
	// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	Tables []string `pulumi:"tables"`
}

// The set of arguments for constructing a Publication resource.
type PublicationArgs struct {
	// Should be ALL TABLES added to the publication. Defaults to 'false'
	AllTables pulumi.BoolPtrInput
	// Which database to create the publication on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// Should all subsequent resources of the publication be dropped. Defaults to 'false'
	DropCascade pulumi.BoolPtrInput
	// The name of the publication.
	Name pulumi.StringPtrInput
	// Who owns the publication. Defaults to provider user.
	Owner pulumi.StringPtrInput
	// Which 'publish' options should be turned on. Default to 'insert','update','delete'
	PublishParams pulumi.StringArrayInput
	// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
	PublishViaPartitionRootParam pulumi.BoolPtrInput
	// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
	Tables pulumi.StringArrayInput
}

func (PublicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicationArgs)(nil)).Elem()
}

type PublicationInput interface {
	pulumi.Input

	ToPublicationOutput() PublicationOutput
	ToPublicationOutputWithContext(ctx context.Context) PublicationOutput
}

func (*Publication) ElementType() reflect.Type {
	return reflect.TypeOf((**Publication)(nil)).Elem()
}

func (i *Publication) ToPublicationOutput() PublicationOutput {
	return i.ToPublicationOutputWithContext(context.Background())
}

func (i *Publication) ToPublicationOutputWithContext(ctx context.Context) PublicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationOutput)
}

// PublicationArrayInput is an input type that accepts PublicationArray and PublicationArrayOutput values.
// You can construct a concrete instance of `PublicationArrayInput` via:
//
//	PublicationArray{ PublicationArgs{...} }
type PublicationArrayInput interface {
	pulumi.Input

	ToPublicationArrayOutput() PublicationArrayOutput
	ToPublicationArrayOutputWithContext(context.Context) PublicationArrayOutput
}

type PublicationArray []PublicationInput

func (PublicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Publication)(nil)).Elem()
}

func (i PublicationArray) ToPublicationArrayOutput() PublicationArrayOutput {
	return i.ToPublicationArrayOutputWithContext(context.Background())
}

func (i PublicationArray) ToPublicationArrayOutputWithContext(ctx context.Context) PublicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationArrayOutput)
}

// PublicationMapInput is an input type that accepts PublicationMap and PublicationMapOutput values.
// You can construct a concrete instance of `PublicationMapInput` via:
//
//	PublicationMap{ "key": PublicationArgs{...} }
type PublicationMapInput interface {
	pulumi.Input

	ToPublicationMapOutput() PublicationMapOutput
	ToPublicationMapOutputWithContext(context.Context) PublicationMapOutput
}

type PublicationMap map[string]PublicationInput

func (PublicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Publication)(nil)).Elem()
}

func (i PublicationMap) ToPublicationMapOutput() PublicationMapOutput {
	return i.ToPublicationMapOutputWithContext(context.Background())
}

func (i PublicationMap) ToPublicationMapOutputWithContext(ctx context.Context) PublicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationMapOutput)
}

type PublicationOutput struct{ *pulumi.OutputState }

func (PublicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Publication)(nil)).Elem()
}

func (o PublicationOutput) ToPublicationOutput() PublicationOutput {
	return o
}

func (o PublicationOutput) ToPublicationOutputWithContext(ctx context.Context) PublicationOutput {
	return o
}

// Should be ALL TABLES added to the publication. Defaults to 'false'
func (o PublicationOutput) AllTables() pulumi.BoolOutput {
	return o.ApplyT(func(v *Publication) pulumi.BoolOutput { return v.AllTables }).(pulumi.BoolOutput)
}

// Which database to create the publication on. Defaults to provider database.
func (o PublicationOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// Should all subsequent resources of the publication be dropped. Defaults to 'false'
func (o PublicationOutput) DropCascade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Publication) pulumi.BoolPtrOutput { return v.DropCascade }).(pulumi.BoolPtrOutput)
}

// The name of the publication.
func (o PublicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Who owns the publication. Defaults to provider user.
func (o PublicationOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Which 'publish' options should be turned on. Default to 'insert','update','delete'
func (o PublicationOutput) PublishParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringArrayOutput { return v.PublishParams }).(pulumi.StringArrayOutput)
}

// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
func (o PublicationOutput) PublishViaPartitionRootParam() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Publication) pulumi.BoolPtrOutput { return v.PublishViaPartitionRootParam }).(pulumi.BoolPtrOutput)
}

// Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
func (o PublicationOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringArrayOutput { return v.Tables }).(pulumi.StringArrayOutput)
}

type PublicationArrayOutput struct{ *pulumi.OutputState }

func (PublicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Publication)(nil)).Elem()
}

func (o PublicationArrayOutput) ToPublicationArrayOutput() PublicationArrayOutput {
	return o
}

func (o PublicationArrayOutput) ToPublicationArrayOutputWithContext(ctx context.Context) PublicationArrayOutput {
	return o
}

func (o PublicationArrayOutput) Index(i pulumi.IntInput) PublicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Publication {
		return vs[0].([]*Publication)[vs[1].(int)]
	}).(PublicationOutput)
}

type PublicationMapOutput struct{ *pulumi.OutputState }

func (PublicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Publication)(nil)).Elem()
}

func (o PublicationMapOutput) ToPublicationMapOutput() PublicationMapOutput {
	return o
}

func (o PublicationMapOutput) ToPublicationMapOutputWithContext(ctx context.Context) PublicationMapOutput {
	return o
}

func (o PublicationMapOutput) MapIndex(k pulumi.StringInput) PublicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Publication {
		return vs[0].(map[string]*Publication)[vs[1].(string)]
	}).(PublicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationInput)(nil)).Elem(), &Publication{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationArrayInput)(nil)).Elem(), PublicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationMapInput)(nil)).Elem(), PublicationMap{})
	pulumi.RegisterOutputType(PublicationOutput{})
	pulumi.RegisterOutputType(PublicationArrayOutput{})
	pulumi.RegisterOutputType(PublicationMapOutput{})
}
