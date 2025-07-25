// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql.Inputs
{

    public sealed class SchemaPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        /// </summary>
        [Input("create")]
        public Input<bool>? Create { get; set; }

        /// <summary>
        /// Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        /// </summary>
        [Input("createWithGrant")]
        public Input<bool>? CreateWithGrant { get; set; }

        /// <summary>
        /// The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        /// </summary>
        [Input("usage")]
        public Input<bool>? Usage { get; set; }

        /// <summary>
        /// Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
        /// 
        /// &gt; **NOTE on `policy`:** The permissions of a role specified in multiple policy blocks is cumulative.  For example, if the same role is specified in two different `policy` each with different permissions (e.g. `create` and `usage_with_grant`, respectively), then the specified role with have both `create` and `usage_with_grant` privileges.
        /// </summary>
        [Input("usageWithGrant")]
        public Input<bool>? UsageWithGrant { get; set; }

        public SchemaPolicyArgs()
        {
        }
        public static new SchemaPolicyArgs Empty => new SchemaPolicyArgs();
    }
}
