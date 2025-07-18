// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “getSequences“ data source retrieves a list of sequence names from a specified PostgreSQL database.
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
//			_, err := postgresql.GetSequences(ctx, &postgresql.GetSequencesArgs{
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
func GetSequences(ctx *pulumi.Context, args *GetSequencesArgs, opts ...pulumi.InvokeOption) (*GetSequencesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSequencesResult
	err := ctx.Invoke("postgresql:index/getSequences:getSequences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSequences.
type GetSequencesArgs struct {
	// The PostgreSQL database which will be queried for sequence names.
	Database string `pulumi:"database"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``LIKE ALL`` operators.
	LikeAllPatterns []string `pulumi:"likeAllPatterns"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``LIKE ANY`` operators.
	LikeAnyPatterns []string `pulumi:"likeAnyPatterns"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
	NotLikeAllPatterns []string `pulumi:"notLikeAllPatterns"`
	// Expression which will be pattern matched against sequence names in the query using the PostgreSQL ``~`` (regular expression match) operator.
	//
	// Note that all optional arguments can be used in conjunction.
	RegexPattern *string `pulumi:"regexPattern"`
	// List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
	Schemas []string `pulumi:"schemas"`
}

// A collection of values returned by getSequences.
type GetSequencesResult struct {
	Database string `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	LikeAllPatterns    []string `pulumi:"likeAllPatterns"`
	LikeAnyPatterns    []string `pulumi:"likeAnyPatterns"`
	NotLikeAllPatterns []string `pulumi:"notLikeAllPatterns"`
	RegexPattern       *string  `pulumi:"regexPattern"`
	Schemas            []string `pulumi:"schemas"`
	// A list of PostgreSQL sequences retrieved by this data source. Each sequence consists of the fields documented below.
	// ***
	Sequences []GetSequencesSequence `pulumi:"sequences"`
}

func GetSequencesOutput(ctx *pulumi.Context, args GetSequencesOutputArgs, opts ...pulumi.InvokeOption) GetSequencesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSequencesResultOutput, error) {
			args := v.(GetSequencesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("postgresql:index/getSequences:getSequences", args, GetSequencesResultOutput{}, options).(GetSequencesResultOutput), nil
		}).(GetSequencesResultOutput)
}

// A collection of arguments for invoking getSequences.
type GetSequencesOutputArgs struct {
	// The PostgreSQL database which will be queried for sequence names.
	Database pulumi.StringInput `pulumi:"database"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``LIKE ALL`` operators.
	LikeAllPatterns pulumi.StringArrayInput `pulumi:"likeAllPatterns"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``LIKE ANY`` operators.
	LikeAnyPatterns pulumi.StringArrayInput `pulumi:"likeAnyPatterns"`
	// List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
	NotLikeAllPatterns pulumi.StringArrayInput `pulumi:"notLikeAllPatterns"`
	// Expression which will be pattern matched against sequence names in the query using the PostgreSQL ``~`` (regular expression match) operator.
	//
	// Note that all optional arguments can be used in conjunction.
	RegexPattern pulumi.StringPtrInput `pulumi:"regexPattern"`
	// List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
	Schemas pulumi.StringArrayInput `pulumi:"schemas"`
}

func (GetSequencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSequencesArgs)(nil)).Elem()
}

// A collection of values returned by getSequences.
type GetSequencesResultOutput struct{ *pulumi.OutputState }

func (GetSequencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSequencesResult)(nil)).Elem()
}

func (o GetSequencesResultOutput) ToGetSequencesResultOutput() GetSequencesResultOutput {
	return o
}

func (o GetSequencesResultOutput) ToGetSequencesResultOutputWithContext(ctx context.Context) GetSequencesResultOutput {
	return o
}

func (o GetSequencesResultOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v GetSequencesResult) string { return v.Database }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSequencesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSequencesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSequencesResultOutput) LikeAllPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSequencesResult) []string { return v.LikeAllPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSequencesResultOutput) LikeAnyPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSequencesResult) []string { return v.LikeAnyPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSequencesResultOutput) NotLikeAllPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSequencesResult) []string { return v.NotLikeAllPatterns }).(pulumi.StringArrayOutput)
}

func (o GetSequencesResultOutput) RegexPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSequencesResult) *string { return v.RegexPattern }).(pulumi.StringPtrOutput)
}

func (o GetSequencesResultOutput) Schemas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSequencesResult) []string { return v.Schemas }).(pulumi.StringArrayOutput)
}

// A list of PostgreSQL sequences retrieved by this data source. Each sequence consists of the fields documented below.
// ***
func (o GetSequencesResultOutput) Sequences() GetSequencesSequenceArrayOutput {
	return o.ApplyT(func(v GetSequencesResult) []GetSequencesSequence { return v.Sequences }).(GetSequencesSequenceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSequencesResultOutput{})
}
