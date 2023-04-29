// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql.Outputs
{

    [OutputType]
    public sealed class GetTablesTableResult
    {
        /// <summary>
        /// The table name.
        /// </summary>
        public readonly string ObjectName;
        /// <summary>
        /// The parent schema.
        /// </summary>
        public readonly string SchemaName;
        /// <summary>
        /// The table type as defined in ``information_schema.tables``.
        /// </summary>
        public readonly string TableType;

        [OutputConstructor]
        private GetTablesTableResult(
            string objectName,

            string schemaName,

            string tableType)
        {
            ObjectName = objectName;
            SchemaName = schemaName;
            TableType = tableType;
        }
    }
}