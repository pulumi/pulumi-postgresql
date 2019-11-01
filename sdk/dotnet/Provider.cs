// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Postgresql
{
    /// <summary>
    /// The provider type for the postgresql package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/index.html.markdown.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, ResourceOptions? options = null)
            : base("postgresql", name, args, MakeResourceOptions(options, ""))
        {
        }

        private static ResourceOptions MakeResourceOptions(ResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        /// </summary>
        [Input("connectTimeout", json: true)]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// The name of the database to connect to in order to conenct to (defaults to `postgres`).
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Database username associated to the connected user (for user name maps)
        /// </summary>
        [Input("databaseUsername")]
        public Input<string>? DatabaseUsername { get; set; }

        /// <summary>
        /// Specify the expected version of PostgreSQL.
        /// </summary>
        [Input("expectedVersion")]
        public Input<string>? ExpectedVersion { get; set; }

        /// <summary>
        /// Name of PostgreSQL server address to connect to
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Maximum number of connections to establish to the database. Zero means unlimited.
        /// </summary>
        [Input("maxConnections", json: true)]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Password to be used if the PostgreSQL server demands password authentication
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain
        /// connections
        /// </summary>
        [Input("port", json: true)]
        public Input<int>? Port { get; set; }

        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        /// <summary>
        /// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with
        /// the PostgreSQL server
        /// </summary>
        [Input("sslmode")]
        public Input<string>? Sslmode { get; set; }

        /// <summary>
        /// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled
        /// (e.g.: Refreshing state password from Postgres)
        /// </summary>
        [Input("superuser", json: true)]
        public Input<bool>? Superuser { get; set; }

        /// <summary>
        /// PostgreSQL user name to connect as
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProviderArgs()
        {
        }
    }
}
