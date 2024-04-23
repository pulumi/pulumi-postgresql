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
    /// The `postgresql.Subscription` resource creates and manages a publication on a PostgreSQL
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
    ///     var subscription = new PostgreSql.Subscription("subscription", new()
    ///     {
    ///         Name = "subscription",
    ///         Conninfo = "host=localhost port=5432 dbname=mydb user=postgres password=postgres",
    ///         Publications = new[]
    ///         {
    ///             "publication",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Postgres documentation
    /// 
    /// - https://www.postgresql.org/docs/current/sql-createsubscription.html
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/subscription:Subscription")]
    public partial class Subscription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
        /// </summary>
        [Output("conninfo")]
        public Output<string> Conninfo { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
        /// </summary>
        [Output("createSlot")]
        public Output<bool?> CreateSlot { get; private set; } = null!;

        /// <summary>
        /// Which database to create the subscription on. Defaults to provider database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Names of the publications on the publisher to subscribe to
        /// </summary>
        [Output("publications")]
        public Output<ImmutableArray<string>> Publications { get; private set; } = null!;

        /// <summary>
        /// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
        /// </summary>
        [Output("slotName")]
        public Output<string?> SlotName { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/subscription:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/subscription:Subscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "conninfo",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, state, options);
        }
    }

    public sealed class SubscriptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("conninfo", required: true)]
        private Input<string>? _conninfo;

        /// <summary>
        /// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
        /// </summary>
        public Input<string>? Conninfo
        {
            get => _conninfo;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _conninfo = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
        /// </summary>
        [Input("createSlot")]
        public Input<bool>? CreateSlot { get; set; }

        /// <summary>
        /// Which database to create the subscription on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("publications", required: true)]
        private InputList<string>? _publications;

        /// <summary>
        /// Names of the publications on the publisher to subscribe to
        /// </summary>
        public InputList<string> Publications
        {
            get => _publications ?? (_publications = new InputList<string>());
            set => _publications = value;
        }

        /// <summary>
        /// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
        /// </summary>
        [Input("slotName")]
        public Input<string>? SlotName { get; set; }

        public SubscriptionArgs()
        {
        }
        public static new SubscriptionArgs Empty => new SubscriptionArgs();
    }

    public sealed class SubscriptionState : global::Pulumi.ResourceArgs
    {
        [Input("conninfo")]
        private Input<string>? _conninfo;

        /// <summary>
        /// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
        /// </summary>
        public Input<string>? Conninfo
        {
            get => _conninfo;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _conninfo = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
        /// </summary>
        [Input("createSlot")]
        public Input<bool>? CreateSlot { get; set; }

        /// <summary>
        /// Which database to create the subscription on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("publications")]
        private InputList<string>? _publications;

        /// <summary>
        /// Names of the publications on the publisher to subscribe to
        /// </summary>
        public InputList<string> Publications
        {
            get => _publications ?? (_publications = new InputList<string>());
            set => _publications = value;
        }

        /// <summary>
        /// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
        /// </summary>
        [Input("slotName")]
        public Input<string>? SlotName { get; set; }

        public SubscriptionState()
        {
        }
        public static new SubscriptionState Empty => new SubscriptionState();
    }
}
