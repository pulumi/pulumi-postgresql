// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``postgresql.getTables`` data source retrieves a list of table names from a specified PostgreSQL database.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const myTables = postgresql.getTables({
 *     database: "my_database",
 * });
 * ```
 */
export function getTables(args: GetTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetTablesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("postgresql:index/getTables:getTables", {
        "database": args.database,
        "likeAllPatterns": args.likeAllPatterns,
        "likeAnyPatterns": args.likeAnyPatterns,
        "notLikeAllPatterns": args.notLikeAllPatterns,
        "regexPattern": args.regexPattern,
        "schemas": args.schemas,
        "tableTypes": args.tableTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getTables.
 */
export interface GetTablesArgs {
    /**
     * The PostgreSQL database which will be queried for table names.
     */
    database: string;
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
     */
    likeAllPatterns?: string[];
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
     */
    likeAnyPatterns?: string[];
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     */
    notLikeAllPatterns?: string[];
    /**
     * Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
     */
    regexPattern?: string;
    /**
     * List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
     */
    schemas?: string[];
    /**
     * List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use 'BASE TABLE' for normal tables only.
     */
    tableTypes?: string[];
}

/**
 * A collection of values returned by getTables.
 */
export interface GetTablesResult {
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly likeAllPatterns?: string[];
    readonly likeAnyPatterns?: string[];
    readonly notLikeAllPatterns?: string[];
    readonly regexPattern?: string;
    readonly schemas?: string[];
    readonly tableTypes?: string[];
    /**
     * A list of PostgreSQL tables retrieved by this data source. Each table consists of the fields documented below.
     * ___
     */
    readonly tables: outputs.GetTablesTable[];
}
/**
 * The ``postgresql.getTables`` data source retrieves a list of table names from a specified PostgreSQL database.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const myTables = postgresql.getTables({
 *     database: "my_database",
 * });
 * ```
 */
export function getTablesOutput(args: GetTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTablesResult> {
    return pulumi.output(args).apply((a: any) => getTables(a, opts))
}

/**
 * A collection of arguments for invoking getTables.
 */
export interface GetTablesOutputArgs {
    /**
     * The PostgreSQL database which will be queried for table names.
     */
    database: pulumi.Input<string>;
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
     */
    likeAllPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
     */
    likeAnyPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     */
    notLikeAllPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
     */
    regexPattern?: pulumi.Input<string>;
    /**
     * List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
     */
    schemas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use 'BASE TABLE' for normal tables only.
     */
    tableTypes?: pulumi.Input<pulumi.Input<string>[]>;
}
