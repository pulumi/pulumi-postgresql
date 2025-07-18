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
    /// The ``postgresql.PhysicalReplicationSlot`` resource creates and manages a physical replication slot on a PostgreSQL
    /// server. This is useful to setup a cross datacenter replication, with Patroni for example, or permit
    /// any stand-by cluster to replicate physically data.
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
    ///     var mySlot = new PostgreSql.PhysicalReplicationSlot("my_slot", new()
    ///     {
    ///         Name = "my_slot",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot")]
    public partial class PhysicalReplicationSlot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PhysicalReplicationSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PhysicalReplicationSlot(string name, PhysicalReplicationSlotArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot", name, args ?? new PhysicalReplicationSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PhysicalReplicationSlot(string name, Input<string> id, PhysicalReplicationSlotState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PhysicalReplicationSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PhysicalReplicationSlot Get(string name, Input<string> id, PhysicalReplicationSlotState? state = null, CustomResourceOptions? options = null)
        {
            return new PhysicalReplicationSlot(name, id, state, options);
        }
    }

    public sealed class PhysicalReplicationSlotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PhysicalReplicationSlotArgs()
        {
        }
        public static new PhysicalReplicationSlotArgs Empty => new PhysicalReplicationSlotArgs();
    }

    public sealed class PhysicalReplicationSlotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PhysicalReplicationSlotState()
        {
        }
        public static new PhysicalReplicationSlotState Empty => new PhysicalReplicationSlotState();
    }
}
