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
    /// The ``postgresql.SecurityLabel`` resource creates and manages security labels.
    /// 
    /// See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-security-label.html)
    /// 
    /// &gt; **Note:** This resource needs Postgresql version 11 or above.
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
    ///     var myRole = new PostgreSql.Role("my_role", new()
    ///     {
    ///         Name = "my_role",
    ///         Login = true,
    ///     });
    /// 
    ///     var workload = new PostgreSql.SecurityLabel("workload", new()
    ///     {
    ///         ObjectType = "role",
    ///         ObjectName = myRole.Name,
    ///         LabelProvider = "pgaadauth",
    ///         Label = "aadauth,oid=00000000-0000-0000-0000-000000000000,type=service",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security label is an attribute that can be added multiple times, so no import is needed, simply apply again.
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/securityLabel:SecurityLabel")]
    public partial class SecurityLabel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The value of the security label.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// The name of the provider with which this label is to be associated.
        /// </summary>
        [Output("labelProvider")]
        public Output<string> LabelProvider { get; private set; } = null!;

        /// <summary>
        /// The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
        /// </summary>
        [Output("objectName")]
        public Output<string> ObjectName { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to apply this security label to.
        /// </summary>
        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityLabel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityLabel(string name, SecurityLabelArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/securityLabel:SecurityLabel", name, args ?? new SecurityLabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityLabel(string name, Input<string> id, SecurityLabelState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/securityLabel:SecurityLabel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityLabel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityLabel Get(string name, Input<string> id, SecurityLabelState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityLabel(name, id, state, options);
        }
    }

    public sealed class SecurityLabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of the security label.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The name of the provider with which this label is to be associated.
        /// </summary>
        [Input("labelProvider", required: true)]
        public Input<string> LabelProvider { get; set; } = null!;

        /// <summary>
        /// The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
        /// </summary>
        [Input("objectName", required: true)]
        public Input<string> ObjectName { get; set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to apply this security label to.
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        public SecurityLabelArgs()
        {
        }
        public static new SecurityLabelArgs Empty => new SecurityLabelArgs();
    }

    public sealed class SecurityLabelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of the security label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The name of the provider with which this label is to be associated.
        /// </summary>
        [Input("labelProvider")]
        public Input<string>? LabelProvider { get; set; }

        /// <summary>
        /// The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
        /// </summary>
        [Input("objectName")]
        public Input<string>? ObjectName { get; set; }

        /// <summary>
        /// The PostgreSQL object type to apply this security label to.
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        public SecurityLabelState()
        {
        }
        public static new SecurityLabelState Empty => new SecurityLabelState();
    }
}