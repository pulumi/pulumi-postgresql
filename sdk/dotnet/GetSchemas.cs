// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    public static class GetSchemas
    {
        /// <summary>
        /// The ``postgresql.getSchemas`` data source retrieves a list of schema names from a specified PostgreSQL database.
        /// 
        /// 
        /// ## Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using PostgreSql = Pulumi.PostgreSql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySchemas = PostgreSql.GetSchemas.Invoke(new()
        ///     {
        ///         Database = "my_database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSchemasResult> InvokeAsync(GetSchemasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSchemasResult>("postgresql:index/getSchemas:getSchemas", args ?? new GetSchemasArgs(), options.WithDefaults());

        /// <summary>
        /// The ``postgresql.getSchemas`` data source retrieves a list of schema names from a specified PostgreSQL database.
        /// 
        /// 
        /// ## Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using PostgreSql = Pulumi.PostgreSql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySchemas = PostgreSql.GetSchemas.Invoke(new()
        ///     {
        ///         Database = "my_database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSchemasResult> Invoke(GetSchemasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemasResult>("postgresql:index/getSchemas:getSchemas", args ?? new GetSchemasInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``postgresql.getSchemas`` data source retrieves a list of schema names from a specified PostgreSQL database.
        /// 
        /// 
        /// ## Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using PostgreSql = Pulumi.PostgreSql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySchemas = PostgreSql.GetSchemas.Invoke(new()
        ///     {
        ///         Database = "my_database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSchemasResult> Invoke(GetSchemasInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemasResult>("postgresql:index/getSchemas:getSchemas", args ?? new GetSchemasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PostgreSQL database which will be queried for schema names.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
        /// </summary>
        [Input("includeSystemSchemas")]
        public bool? IncludeSystemSchemas { get; set; }

        [Input("likeAllPatterns")]
        private List<string>? _likeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
        /// </summary>
        public List<string> LikeAllPatterns
        {
            get => _likeAllPatterns ?? (_likeAllPatterns = new List<string>());
            set => _likeAllPatterns = value;
        }

        [Input("likeAnyPatterns")]
        private List<string>? _likeAnyPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
        /// </summary>
        public List<string> LikeAnyPatterns
        {
            get => _likeAnyPatterns ?? (_likeAnyPatterns = new List<string>());
            set => _likeAnyPatterns = value;
        }

        [Input("notLikeAllPatterns")]
        private List<string>? _notLikeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
        /// </summary>
        public List<string> NotLikeAllPatterns
        {
            get => _notLikeAllPatterns ?? (_notLikeAllPatterns = new List<string>());
            set => _notLikeAllPatterns = value;
        }

        /// <summary>
        /// Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
        /// 
        /// Note that all optional arguments can be used in conjunction.
        /// </summary>
        [Input("regexPattern")]
        public string? RegexPattern { get; set; }

        public GetSchemasArgs()
        {
        }
        public static new GetSchemasArgs Empty => new GetSchemasArgs();
    }

    public sealed class GetSchemasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PostgreSQL database which will be queried for schema names.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
        /// </summary>
        [Input("includeSystemSchemas")]
        public Input<bool>? IncludeSystemSchemas { get; set; }

        [Input("likeAllPatterns")]
        private InputList<string>? _likeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
        /// </summary>
        public InputList<string> LikeAllPatterns
        {
            get => _likeAllPatterns ?? (_likeAllPatterns = new InputList<string>());
            set => _likeAllPatterns = value;
        }

        [Input("likeAnyPatterns")]
        private InputList<string>? _likeAnyPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
        /// </summary>
        public InputList<string> LikeAnyPatterns
        {
            get => _likeAnyPatterns ?? (_likeAnyPatterns = new InputList<string>());
            set => _likeAnyPatterns = value;
        }

        [Input("notLikeAllPatterns")]
        private InputList<string>? _notLikeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
        /// </summary>
        public InputList<string> NotLikeAllPatterns
        {
            get => _notLikeAllPatterns ?? (_notLikeAllPatterns = new InputList<string>());
            set => _notLikeAllPatterns = value;
        }

        /// <summary>
        /// Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
        /// 
        /// Note that all optional arguments can be used in conjunction.
        /// </summary>
        [Input("regexPattern")]
        public Input<string>? RegexPattern { get; set; }

        public GetSchemasInvokeArgs()
        {
        }
        public static new GetSchemasInvokeArgs Empty => new GetSchemasInvokeArgs();
    }


    [OutputType]
    public sealed class GetSchemasResult
    {
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeSystemSchemas;
        public readonly ImmutableArray<string> LikeAllPatterns;
        public readonly ImmutableArray<string> LikeAnyPatterns;
        public readonly ImmutableArray<string> NotLikeAllPatterns;
        public readonly string? RegexPattern;
        /// <summary>
        /// A list of full names of found schemas.
        /// </summary>
        public readonly ImmutableArray<string> Schemas;

        [OutputConstructor]
        private GetSchemasResult(
            string database,

            string id,

            bool? includeSystemSchemas,

            ImmutableArray<string> likeAllPatterns,

            ImmutableArray<string> likeAnyPatterns,

            ImmutableArray<string> notLikeAllPatterns,

            string? regexPattern,

            ImmutableArray<string> schemas)
        {
            Database = database;
            Id = id;
            IncludeSystemSchemas = includeSystemSchemas;
            LikeAllPatterns = likeAllPatterns;
            LikeAnyPatterns = likeAnyPatterns;
            NotLikeAllPatterns = notLikeAllPatterns;
            RegexPattern = regexPattern;
            Schemas = schemas;
        }
    }
}
