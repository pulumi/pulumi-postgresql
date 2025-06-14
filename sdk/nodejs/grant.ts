// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.Grant`` resource creates and manages privileges given to a user for a database schema.
 *
 * See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-grant.html)
 *
 * > **Note:** This resource needs Postgresql version 9 or above.
 * **Note:** Using column & table grants on the _same_ table with the _same_ privileges can lead to unexpected behaviours.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * // Grant SELECT privileges on 2 tables
 * const readonlyTables = new postgresql.Grant("readonly_tables", {
 *     database: "test_db",
 *     role: "test_role",
 *     schema: "public",
 *     objectType: "table",
 *     objects: [
 *         "table1",
 *         "table2",
 *     ],
 *     privileges: ["SELECT"],
 * });
 * // Grant SELECT & INSERT privileges on 2 columns in 1 table
 * const readInsertColumn = new postgresql.Grant("read_insert_column", {
 *     database: "test_db",
 *     role: "test_role",
 *     schema: "public",
 *     objectType: "column",
 *     objects: ["table1"],
 *     columns: [
 *         "col1",
 *         "col2",
 *     ],
 *     privileges: [
 *         "UPDATE",
 *         "INSERT",
 *     ],
 * });
 * ```
 *
 * ## Examples
 *
 * Revoke default accesses for public schema:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const revokePublic = new postgresql.Grant("revoke_public", {
 *     database: "test_db",
 *     role: "public",
 *     schema: "public",
 *     objectType: "schema",
 *     privileges: [],
 * });
 * ```
 */
export class Grant extends pulumi.CustomResource {
    /**
     * Get an existing Grant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrantState, opts?: pulumi.CustomResourceOptions): Grant {
        return new Grant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/grant:Grant';

    /**
     * Returns true if the given object is an instance of Grant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Grant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Grant.__pulumiType;
    }

    /**
     * The columns upon which to grant the privileges. Required when `objectType` is `column`. You cannot specify this option if the `objectType` is not `column`.
     */
    public readonly columns!: pulumi.Output<string[] | undefined>;
    /**
     * The database to grant privileges on for this role.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
     */
    public readonly objectType!: pulumi.Output<string>;
    /**
     * The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`. When `objectType` is `column`, only one value is allowed.
     */
    public readonly objects!: pulumi.Output<string[] | undefined>;
    /**
     * The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     */
    public readonly privileges!: pulumi.Output<string[]>;
    /**
     * The name of the role to grant privileges on, Set it to "public" for all roles.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The database schema to grant privileges on for this role (Required except if objectType is "database")
     */
    public readonly schema!: pulumi.Output<string | undefined>;
    /**
     * Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Grant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrantArgs | GrantState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrantState | undefined;
            resourceInputs["columns"] = state ? state.columns : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["objectType"] = state ? state.objectType : undefined;
            resourceInputs["objects"] = state ? state.objects : undefined;
            resourceInputs["privileges"] = state ? state.privileges : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
            resourceInputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as GrantArgs | undefined;
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.objectType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectType'");
            }
            if ((!args || args.privileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privileges'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["columns"] = args ? args.columns : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["objectType"] = args ? args.objectType : undefined;
            resourceInputs["objects"] = args ? args.objects : undefined;
            resourceInputs["privileges"] = args ? args.privileges : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Grant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Grant resources.
 */
export interface GrantState {
    /**
     * The columns upon which to grant the privileges. Required when `objectType` is `column`. You cannot specify this option if the `objectType` is not `column`.
     */
    columns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database to grant privileges on for this role.
     */
    database?: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
     */
    objectType?: pulumi.Input<string>;
    /**
     * The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`. When `objectType` is `column`, only one value is allowed.
     */
    objects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     */
    privileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to grant privileges on, Set it to "public" for all roles.
     */
    role?: pulumi.Input<string>;
    /**
     * The database schema to grant privileges on for this role (Required except if objectType is "database")
     */
    schema?: pulumi.Input<string>;
    /**
     * Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Grant resource.
 */
export interface GrantArgs {
    /**
     * The columns upon which to grant the privileges. Required when `objectType` is `column`. You cannot specify this option if the `objectType` is not `column`.
     */
    columns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database to grant privileges on for this role.
     */
    database: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
     */
    objectType: pulumi.Input<string>;
    /**
     * The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`. When `objectType` is `column`, only one value is allowed.
     */
    objects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     */
    privileges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to grant privileges on, Set it to "public" for all roles.
     */
    role: pulumi.Input<string>;
    /**
     * The database schema to grant privileges on for this role (Required except if objectType is "database")
     */
    schema?: pulumi.Input<string>;
    /**
     * Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     */
    withGrantOption?: pulumi.Input<boolean>;
}
