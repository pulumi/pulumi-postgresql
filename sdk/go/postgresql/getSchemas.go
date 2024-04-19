// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “getSchemas“ data source retrieves a list of schema names from a specified PostgreSQL database.
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
//			_, err := postgresql.GetSchemas(ctx, &postgresql.GetSchemasArgs{
//				Database: "my_database",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSchemas(ctx *pulumi.Context, args *GetSchemasArgs, opts ...pulumi.InvokeOption) (*GetSchemasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSchemasResult
	err := ctx.Invoke("postgresql:index/getSchemas:getSchemas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSchemas.
type GetSchemasArgs struct {
	// The PostgreSQL database which will be queried for schema names.
	Database string `pulumi:"database"`
	// Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
	IncludeSystemSchemas *bool `pulumi:"includeSystemSchemas"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
	LikeAllPatterns []string `pulumi:"likeAllPatterns"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
	LikeAnyPatterns []string `pulumi:"likeAnyPatterns"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
	NotLikeAllPatterns []string `pulumi:"notLikeAllPatterns"`
	// Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
	//
	// Note that all optional arguments can be used in conjunction.
	RegexPattern *string `pulumi:"regexPattern"`
}

// A collection of values returned by getSchemas.
type GetSchemasResult struct {
	Database string `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string   `pulumi:"id"`
	IncludeSystemSchemas *bool    `pulumi:"includeSystemSchemas"`
	LikeAllPatterns      []string `pulumi:"likeAllPatterns"`
	LikeAnyPatterns      []string `pulumi:"likeAnyPatterns"`
	NotLikeAllPatterns   []string `pulumi:"notLikeAllPatterns"`
	RegexPattern         *string  `pulumi:"regexPattern"`
	// A list of full names of found schemas.
	Schemas []string `pulumi:"schemas"`
}

func GetSchemasOutput(ctx *pulumi.Context, args GetSchemasOutputArgs, opts ...pulumi.InvokeOption) GetSchemasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSchemasResult, error) {
			args := v.(GetSchemasArgs)
			r, err := GetSchemas(ctx, &args, opts...)
			var s GetSchemasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSchemasResultOutput)
}

// A collection of arguments for invoking getSchemas.
type GetSchemasOutputArgs struct {
	// The PostgreSQL database which will be queried for schema names.
	Database pulumi.StringInput `pulumi:"database"`
	// Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
	IncludeSystemSchemas pulumi.BoolPtrInput `pulumi:"includeSystemSchemas"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
	LikeAllPatterns pulumi.StringArrayInput `pulumi:"likeAllPatterns"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
	LikeAnyPatterns pulumi.StringArrayInput `pulumi:"likeAnyPatterns"`
	// List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
	NotLikeAllPatterns pulumi.StringArrayInput `pulumi:"notLikeAllPatterns"`
	// Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
	//
	// Note that all optional arguments can be used in conjunction.
	RegexPattern pulumi.StringPtrInput `pulumi:"regexPattern"`
}

func (GetSchemasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchemasArgs)(nil)).Elem()
}

// A collection of values returned by getSchemas.
type GetSchemasResultOutput struct{ *pulumi.OutputState }

func (GetSchemasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchemasResult)(nil)).Elem()
}

func (o GetSchemasResultOutput) ToGetSchemasResultOutput() GetSchemasResultOutput {
	return o
}

func (o GetSchemasResultOutput) ToGetSchemasResultOutputWithContext(ctx context.Context) GetSchemasResultOutput {
	return o
}

func (o GetSchemasResultOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v GetSchemasResult) string { return v.Database }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSchemasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSchemasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSchemasResultOutput) IncludeSystemSchemas() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSchemasResult) *bool { return v.IncludeSystemSchemas }).(pulumi.BoolPtrOutput)
}

func (o GetSchemasResultOutput) LikeAllPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSchemasResult) []string { return v.LikeAllPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSchemasResultOutput) LikeAnyPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSchemasResult) []string { return v.LikeAnyPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSchemasResultOutput) NotLikeAllPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSchemasResult) []string { return v.NotLikeAllPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSchemasResultOutput) RegexPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSchemasResult) *string { return v.RegexPattern }).(pulumi.StringPtrOutput)
}

// A list of full names of found schemas.
func (o GetSchemasResultOutput) Schemas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSchemasResult) []string { return v.Schemas }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSchemasResultOutput{})
}
