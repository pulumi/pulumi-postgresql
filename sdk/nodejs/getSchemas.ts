// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.getSchemas`` data source retrieves a list of schema names from a specified PostgreSQL database.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const mySchemas = postgresql.getSchemas({
 *     database: "my_database",
 * });
 * ```
 */
export function getSchemas(args: GetSchemasArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("postgresql:index/getSchemas:getSchemas", {
        "database": args.database,
        "includeSystemSchemas": args.includeSystemSchemas,
        "likeAllPatterns": args.likeAllPatterns,
        "likeAnyPatterns": args.likeAnyPatterns,
        "notLikeAllPatterns": args.notLikeAllPatterns,
        "regexPattern": args.regexPattern,
    }, opts);
}

/**
 * A collection of arguments for invoking getSchemas.
 */
export interface GetSchemasArgs {
    /**
     * The PostgreSQL database which will be queried for schema names.
     */
    database: string;
    /**
     * Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
     */
    includeSystemSchemas?: boolean;
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
     */
    likeAllPatterns?: string[];
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
     */
    likeAnyPatterns?: string[];
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     */
    notLikeAllPatterns?: string[];
    /**
     * Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
     *
     * Note that all optional arguments can be used in conjunction.
     */
    regexPattern?: string;
}

/**
 * A collection of values returned by getSchemas.
 */
export interface GetSchemasResult {
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeSystemSchemas?: boolean;
    readonly likeAllPatterns?: string[];
    readonly likeAnyPatterns?: string[];
    readonly notLikeAllPatterns?: string[];
    readonly regexPattern?: string;
    /**
     * A list of full names of found schemas.
     */
    readonly schemas: string[];
}
/**
 * The ``postgresql.getSchemas`` data source retrieves a list of schema names from a specified PostgreSQL database.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const mySchemas = postgresql.getSchemas({
 *     database: "my_database",
 * });
 * ```
 */
export function getSchemasOutput(args: GetSchemasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("postgresql:index/getSchemas:getSchemas", {
        "database": args.database,
        "includeSystemSchemas": args.includeSystemSchemas,
        "likeAllPatterns": args.likeAllPatterns,
        "likeAnyPatterns": args.likeAnyPatterns,
        "notLikeAllPatterns": args.notLikeAllPatterns,
        "regexPattern": args.regexPattern,
    }, opts);
}

/**
 * A collection of arguments for invoking getSchemas.
 */
export interface GetSchemasOutputArgs {
    /**
     * The PostgreSQL database which will be queried for schema names.
     */
    database: pulumi.Input<string>;
    /**
     * Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
     */
    includeSystemSchemas?: pulumi.Input<boolean>;
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
     */
    likeAllPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
     */
    likeAnyPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     */
    notLikeAllPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
     *
     * Note that all optional arguments can be used in conjunction.
     */
    regexPattern?: pulumi.Input<string>;
}
