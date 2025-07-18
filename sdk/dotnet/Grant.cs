// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-grant.html)
    /// 
    /// &gt; **Note:** This resource needs Postgresql version 9 or above.
    /// **Note:** Using column &amp; table grants on the _same_ table with the _same_ privileges can lead to unexpected behaviours.
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
    ///     // Grant SELECT privileges on 2 tables
    ///     var readonlyTables = new PostgreSql.Grant("readonly_tables", new()
    ///     {
    ///         Database = "test_db",
    ///         Role = "test_role",
    ///         Schema = "public",
    ///         ObjectType = "table",
    ///         Objects = new[]
    ///         {
    ///             "table1",
    ///             "table2",
    ///         },
    ///         Privileges = new[]
    ///         {
    ///             "SELECT",
    ///         },
    ///     });
    /// 
    ///     // Grant SELECT &amp; INSERT privileges on 2 columns in 1 table
    ///     var readInsertColumn = new PostgreSql.Grant("read_insert_column", new()
    ///     {
    ///         Database = "test_db",
    ///         Role = "test_role",
    ///         Schema = "public",
    ///         ObjectType = "column",
    ///         Objects = new[]
    ///         {
    ///             "table1",
    ///         },
    ///         Columns = new[]
    ///         {
    ///             "col1",
    ///             "col2",
    ///         },
    ///         Privileges = new[]
    ///         {
    ///             "UPDATE",
    ///             "INSERT",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Examples
    /// 
    /// Revoke default accesses for public schema:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using PostgreSql = Pulumi.PostgreSql;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var revokePublic = new PostgreSql.Grant("revoke_public", new()
    ///     {
    ///         Database = "test_db",
    ///         Role = "public",
    ///         Schema = "public",
    ///         ObjectType = "schema",
    ///         Privileges = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/grant:Grant")]
    public partial class Grant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
        /// </summary>
        [Output("columns")]
        public Output<ImmutableArray<string>> Columns { get; private set; } = null!;

        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
        /// </summary>
        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;

        /// <summary>
        /// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
        /// </summary>
        [Output("objects")]
        public Output<ImmutableArray<string>> Objects { get; private set; } = null!;

        /// <summary>
        /// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
        /// </summary>
        [Output("privileges")]
        public Output<ImmutableArray<string>> Privileges { get; private set; } = null!;

        /// <summary>
        /// The name of the role to grant privileges on, Set it to "public" for all roles.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The database schema to grant privileges on for this role (Required except if object_type is "database")
        /// </summary>
        [Output("schema")]
        public Output<string?> Schema { get; private set; } = null!;

        /// <summary>
        /// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
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

    public sealed class GrantArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<string>? _columns;

        /// <summary>
        /// The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
        /// </summary>
        public InputList<string> Columns
        {
            get => _columns ?? (_columns = new InputList<string>());
            set => _columns = value;
        }

        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        [Input("objects")]
        private InputList<string>? _objects;

        /// <summary>
        /// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
        /// </summary>
        public InputList<string> Objects
        {
            get => _objects ?? (_objects = new InputList<string>());
            set => _objects = value;
        }

        [Input("privileges", required: true)]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to grant privileges on, Set it to "public" for all roles.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The database schema to grant privileges on for this role (Required except if object_type is "database")
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public GrantArgs()
        {
        }
        public static new GrantArgs Empty => new GrantArgs();
    }

    public sealed class GrantState : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<string>? _columns;

        /// <summary>
        /// The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
        /// </summary>
        public InputList<string> Columns
        {
            get => _columns ?? (_columns = new InputList<string>());
            set => _columns = value;
        }

        /// <summary>
        /// The database to grant privileges on for this role.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        [Input("objects")]
        private InputList<string>? _objects;

        /// <summary>
        /// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
        /// </summary>
        public InputList<string> Objects
        {
            get => _objects ?? (_objects = new InputList<string>());
            set => _objects = value;
        }

        [Input("privileges")]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to grant privileges on, Set it to "public" for all roles.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The database schema to grant privileges on for this role (Required except if object_type is "database")
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public GrantState()
        {
        }
        public static new GrantState Empty => new GrantState();
    }
}
