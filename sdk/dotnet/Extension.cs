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
    /// The ``postgresql.Extension`` resource creates and manages an extension on a PostgreSQL
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
    ///     var myExtension = new PostgreSql.Extension("myExtension");
    /// 
    /// });
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/extension:Extension")]
    public partial class Extension : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
        /// </summary>
        [Output("createCascade")]
        public Output<bool?> CreateCascade { get; private set; } = null!;

        /// <summary>
        /// Which database to create the extension on. Defaults to provider database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
        /// </summary>
        [Output("dropCascade")]
        public Output<bool?> DropCascade { get; private set; } = null!;

        /// <summary>
        /// The name of the extension.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets the schema of an extension.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// Sets the version number of the extension.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Extension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Extension(string name, ExtensionArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql:index/extension:Extension", name, args ?? new ExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Extension(string name, Input<string> id, ExtensionState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/extension:Extension", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Extension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Extension Get(string name, Input<string> id, ExtensionState? state = null, CustomResourceOptions? options = null)
        {
            return new Extension(name, id, state, options);
        }
    }

    public sealed class ExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
        /// </summary>
        [Input("createCascade")]
        public Input<bool>? CreateCascade { get; set; }

        /// <summary>
        /// Which database to create the extension on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The name of the extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the schema of an extension.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Sets the version number of the extension.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ExtensionArgs()
        {
        }
        public static new ExtensionArgs Empty => new ExtensionArgs();
    }

    public sealed class ExtensionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
        /// </summary>
        [Input("createCascade")]
        public Input<bool>? CreateCascade { get; set; }

        /// <summary>
        /// Which database to create the extension on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The name of the extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the schema of an extension.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Sets the version number of the extension.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ExtensionState()
        {
        }
        public static new ExtensionState Empty => new ExtensionState();
    }
}
