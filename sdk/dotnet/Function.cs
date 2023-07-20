// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    /// <summary>
    /// The ``postgresql.Function`` resource creates and manages a function on a PostgreSQL
    /// server.
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
    ///     var increment = new PostgreSql.Function("increment", new()
    ///     {
    ///         Args = new[]
    ///         {
    ///             new PostgreSql.Inputs.FunctionArgArgs
    ///             {
    ///                 Name = "i",
    ///                 Type = "integer",
    ///             },
    ///         },
    ///         Body = @"    BEGIN
    ///         RETURN i + 1;
    ///     END;
    /// 
    /// ",
    ///         Language = "plpgsql",
    ///         Returns = "integer",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/function:Function")]
    public partial class Function : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of arguments for the function.
        /// </summary>
        [Output("args")]
        public Output<ImmutableArray<Outputs.FunctionArg>> Args { get; private set; } = null!;

        /// <summary>
        /// Function body.
        /// This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// The database where the function is located.
        /// If not specified, the function is created in the current database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// True to automatically drop objects that depend on the function (such as 
        /// operators or triggers), and in turn all objects that depend on those objects. Default is false.
        /// </summary>
        [Output("dropCascade")]
        public Output<bool?> DropCascade { get; private set; } = null!;

        /// <summary>
        /// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        /// </summary>
        [Output("language")]
        public Output<string?> Language { get; private set; } = null!;

        /// <summary>
        /// The name of the argument.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type that the function returns. It can be computed from the OUT arguments. Default is void.
        /// </summary>
        [Output("returns")]
        public Output<string> Returns { get; private set; } = null!;

        /// <summary>
        /// The schema where the function is located.
        /// If not specified, the function is created in the current schema.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/function:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/function:Function", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<Inputs.FunctionArgArgs>? _args;

        /// <summary>
        /// List of arguments for the function.
        /// </summary>
        public InputList<Inputs.FunctionArgArgs> Args
        {
            get => _args ?? (_args = new InputList<Inputs.FunctionArgArgs>());
            set => _args = value;
        }

        /// <summary>
        /// Function body.
        /// This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        /// <summary>
        /// The database where the function is located.
        /// If not specified, the function is created in the current database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// True to automatically drop objects that depend on the function (such as 
        /// operators or triggers), and in turn all objects that depend on those objects. Default is false.
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        /// <summary>
        /// The name of the argument.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type that the function returns. It can be computed from the OUT arguments. Default is void.
        /// </summary>
        [Input("returns")]
        public Input<string>? Returns { get; set; }

        /// <summary>
        /// The schema where the function is located.
        /// If not specified, the function is created in the current schema.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public FunctionArgs()
        {
        }
        public static new FunctionArgs Empty => new FunctionArgs();
    }

    public sealed class FunctionState : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<Inputs.FunctionArgGetArgs>? _args;

        /// <summary>
        /// List of arguments for the function.
        /// </summary>
        public InputList<Inputs.FunctionArgGetArgs> Args
        {
            get => _args ?? (_args = new InputList<Inputs.FunctionArgGetArgs>());
            set => _args = value;
        }

        /// <summary>
        /// Function body.
        /// This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The database where the function is located.
        /// If not specified, the function is created in the current database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// True to automatically drop objects that depend on the function (such as 
        /// operators or triggers), and in turn all objects that depend on those objects. Default is false.
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        /// <summary>
        /// The name of the argument.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type that the function returns. It can be computed from the OUT arguments. Default is void.
        /// </summary>
        [Input("returns")]
        public Input<string>? Returns { get; set; }

        /// <summary>
        /// The schema where the function is located.
        /// If not specified, the function is created in the current schema.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public FunctionState()
        {
        }
        public static new FunctionState Empty => new FunctionState();
    }
}
