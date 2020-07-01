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
    /// The ``postgresql.Grant`` resource creates and manages privileges given to a user for a database schema.
    /// 
    /// &gt; **Note:** This resource needs Postgresql version 9 or above.
    /// 
    /// ## Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using PostgreSql = Pulumi.PostgreSql;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var readonlyTables = new PostgreSql.Grant("readonlyTables", new PostgreSql.GrantArgs
    ///         {
    ///             Database = "test_db",
    ///             ObjectType = "table",
    ///             Privileges = 
    ///             {
    ///                 "SELECT",
    ///             },
    ///             Role = "test_role",
    ///             Schema = "public",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Grant : Pulumi.CustomResource
    {
        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
        /// </summary>
        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;

        /// <summary>
        /// The list of privileges to grant.
        /// </summary>
        [Output("privileges")]
        public Output<ImmutableArray<string>> Privileges { get; private set; } = null!;

        /// <summary>
        /// The name of the role to grant privileges on.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The database schema to grant privileges on for this role.
        /// </summary>
        [Output("schema")]
        public Output<string?> Schema { get; private set; } = null!;

        /// <summary>
        /// Permit the grant recipient to grant it to others
        /// </summary>
        [Output("withGrantOption")]
        public Output<bool?> WithGrantOption { get; private set; } = null!;


        /// <summary>
        /// Create a Grant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Grant(string name, GrantArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/grant:Grant", name, args ?? new GrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Grant(string name, Input<string> id, GrantState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/grant:Grant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Grant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Grant Get(string name, Input<string> id, GrantState? state = null, CustomResourceOptions? options = null)
        {
            return new Grant(name, id, state, options);
        }
    }

    public sealed class GrantArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        [Input("privileges", required: true)]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to grant.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to grant privileges on.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The database schema to grant privileges on for this role.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Permit the grant recipient to grant it to others
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public GrantArgs()
        {
        }
    }

    public sealed class GrantState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        [Input("privileges")]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to grant.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to grant privileges on.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The database schema to grant privileges on for this role.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Permit the grant recipient to grant it to others
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public GrantState()
        {
        }
    }
}
