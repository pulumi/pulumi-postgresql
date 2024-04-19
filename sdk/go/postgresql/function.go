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

// The “Function“ resource creates and manages a function on a PostgreSQL
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
//			_, err := postgresql.NewFunction(ctx, "increment", &postgresql.FunctionArgs{
//				Name: pulumi.String("increment"),
//				Args: postgresql.FunctionArgArray{
//					&postgresql.FunctionArgArgs{
//						Name: pulumi.String("i"),
//						Type: pulumi.String("integer"),
//					},
//				},
//				Returns:  pulumi.String("integer"),
//				Language: pulumi.String("plpgsql"),
//				Body:     pulumi.String("BEGIN\n    RETURN i + 1;\nEND;\n"),
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
// ## Import
//
// It is possible to import a `postgresql_function` resource with the following
// command:
//
// ```sh
// $ pulumi import postgresql:index/function:Function function_foo "my_database.my_schema.my_function_name(arguments)"
// ```
// Where `my_database` is the name of the database containing the schema,
// `my_schema` is the name of the schema in the PostgreSQL database, `my_function_name` is the function name to be imported, `arguments` is the argument signature of the function including all non OUT types and
// `postgresql_schema.function_foo` is the name of the resource whose state will be
// populated as a result of the command.
type Function struct {
	pulumi.CustomResourceState

	// List of arguments for the function.
	Args FunctionArgArrayOutput `pulumi:"args"`
	// Function body.
	// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
	Body pulumi.StringOutput `pulumi:"body"`
	// The database where the function is located.
	// If not specified, the function is created in the current database.
	Database pulumi.StringOutput `pulumi:"database"`
	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	DropCascade pulumi.BoolPtrOutput `pulumi:"dropCascade"`
	// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// The name of the argument.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
	Parallel pulumi.StringPtrOutput `pulumi:"parallel"`
	// Type that the function returns. It can be computed from the OUT arguments. Default is void.
	Returns pulumi.StringOutput `pulumi:"returns"`
	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
	SecurityDefiner pulumi.BoolPtrOutput `pulumi:"securityDefiner"`
	// If the function should always return NULL when any of the inputs is NULL. Default is false.
	Strict pulumi.BoolPtrOutput `pulumi:"strict"`
	// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
	Volatility pulumi.StringPtrOutput `pulumi:"volatility"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Function
	err := ctx.RegisterResource("postgresql:index/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("postgresql:index/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// List of arguments for the function.
	Args []FunctionArg `pulumi:"args"`
	// Function body.
	// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
	Body *string `pulumi:"body"`
	// The database where the function is located.
	// If not specified, the function is created in the current database.
	Database *string `pulumi:"database"`
	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	DropCascade *bool `pulumi:"dropCascade"`
	// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
	Language *string `pulumi:"language"`
	// The name of the argument.
	Name *string `pulumi:"name"`
	// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
	Parallel *string `pulumi:"parallel"`
	// Type that the function returns. It can be computed from the OUT arguments. Default is void.
	Returns *string `pulumi:"returns"`
	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	Schema *string `pulumi:"schema"`
	// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
	SecurityDefiner *bool `pulumi:"securityDefiner"`
	// If the function should always return NULL when any of the inputs is NULL. Default is false.
	Strict *bool `pulumi:"strict"`
	// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
	Volatility *string `pulumi:"volatility"`
}

type FunctionState struct {
	// List of arguments for the function.
	Args FunctionArgArrayInput
	// Function body.
	// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
	Body pulumi.StringPtrInput
	// The database where the function is located.
	// If not specified, the function is created in the current database.
	Database pulumi.StringPtrInput
	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	DropCascade pulumi.BoolPtrInput
	// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
	Language pulumi.StringPtrInput
	// The name of the argument.
	Name pulumi.StringPtrInput
	// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
	Parallel pulumi.StringPtrInput
	// Type that the function returns. It can be computed from the OUT arguments. Default is void.
	Returns pulumi.StringPtrInput
	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	Schema pulumi.StringPtrInput
	// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
	SecurityDefiner pulumi.BoolPtrInput
	// If the function should always return NULL when any of the inputs is NULL. Default is false.
	Strict pulumi.BoolPtrInput
	// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
	Volatility pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// List of arguments for the function.
	Args []FunctionArg `pulumi:"args"`
	// Function body.
	// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
	Body string `pulumi:"body"`
	// The database where the function is located.
	// If not specified, the function is created in the current database.
	Database *string `pulumi:"database"`
	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	DropCascade *bool `pulumi:"dropCascade"`
	// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
	Language *string `pulumi:"language"`
	// The name of the argument.
	Name *string `pulumi:"name"`
	// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
	Parallel *string `pulumi:"parallel"`
	// Type that the function returns. It can be computed from the OUT arguments. Default is void.
	Returns *string `pulumi:"returns"`
	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	Schema *string `pulumi:"schema"`
	// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
	SecurityDefiner *bool `pulumi:"securityDefiner"`
	// If the function should always return NULL when any of the inputs is NULL. Default is false.
	Strict *bool `pulumi:"strict"`
	// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
	Volatility *string `pulumi:"volatility"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// List of arguments for the function.
	Args FunctionArgArrayInput
	// Function body.
	// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
	Body pulumi.StringInput
	// The database where the function is located.
	// If not specified, the function is created in the current database.
	Database pulumi.StringPtrInput
	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	DropCascade pulumi.BoolPtrInput
	// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
	Language pulumi.StringPtrInput
	// The name of the argument.
	Name pulumi.StringPtrInput
	// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
	Parallel pulumi.StringPtrInput
	// Type that the function returns. It can be computed from the OUT arguments. Default is void.
	Returns pulumi.StringPtrInput
	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	Schema pulumi.StringPtrInput
	// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
	SecurityDefiner pulumi.BoolPtrInput
	// If the function should always return NULL when any of the inputs is NULL. Default is false.
	Strict pulumi.BoolPtrInput
	// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
	Volatility pulumi.StringPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//	FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//	FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

// List of arguments for the function.
func (o FunctionOutput) Args() FunctionArgArrayOutput {
	return o.ApplyT(func(v *Function) FunctionArgArrayOutput { return v.Args }).(FunctionArgArrayOutput)
}

// Function body.
// This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
func (o FunctionOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

// The database where the function is located.
// If not specified, the function is created in the current database.
func (o FunctionOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// True to automatically drop objects that depend on the function (such as
// operators or triggers), and in turn all objects that depend on those objects. Default is false.
func (o FunctionOutput) DropCascade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.BoolPtrOutput { return v.DropCascade }).(pulumi.BoolPtrOutput)
}

// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
func (o FunctionOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Language }).(pulumi.StringPtrOutput)
}

// The name of the argument.
func (o FunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
func (o FunctionOutput) Parallel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Parallel }).(pulumi.StringPtrOutput)
}

// Type that the function returns. It can be computed from the OUT arguments. Default is void.
func (o FunctionOutput) Returns() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Returns }).(pulumi.StringOutput)
}

// The schema where the function is located.
// If not specified, the function is created in the current schema.
func (o FunctionOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
func (o FunctionOutput) SecurityDefiner() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.BoolPtrOutput { return v.SecurityDefiner }).(pulumi.BoolPtrOutput)
}

// If the function should always return NULL when any of the inputs is NULL. Default is false.
func (o FunctionOutput) Strict() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.BoolPtrOutput { return v.Strict }).(pulumi.BoolPtrOutput)
}

// Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
func (o FunctionOutput) Volatility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Volatility }).(pulumi.StringPtrOutput)
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Function {
		return vs[0].([]*Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Function {
		return vs[0].(map[string]*Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionArrayInput)(nil)).Elem(), FunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionMapInput)(nil)).Elem(), FunctionMap{})
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}
