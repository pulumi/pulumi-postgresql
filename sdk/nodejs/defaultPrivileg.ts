// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.
 *
 * > **Note:** This resource needs Postgresql version 9 or above.
 *
 * ## Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const readOnlyTables = new postgresql.DefaultPrivileges("readOnlyTables", {
 *     database: "test_db",
 *     objectType: "table",
 *     owner: "db_owner",
 *     privileges: ["SELECT"],
 *     role: "test_role",
 *     schema: "public",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Examples
 *
 * Revoke default privileges for functions for "public" role:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const revokePublic = new postgresql.DefaultPrivileges("revokePublic", {
 *     database: postgresql_database.example_db.name,
 *     role: "public",
 *     owner: "object_owner",
 *     objectType: "function",
 *     privileges: [],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * @deprecated postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges
 */
export class DefaultPrivileg extends pulumi.CustomResource {
    /**
     * Get an existing DefaultPrivileg resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultPrivilegState, opts?: pulumi.CustomResourceOptions): DefaultPrivileg {
        pulumi.log.warn("DefaultPrivileg is deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges")
        return new DefaultPrivileg(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/defaultPrivileg:DefaultPrivileg';

    /**
     * Returns true if the given object is an instance of DefaultPrivileg.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultPrivileg {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultPrivileg.__pulumiType;
    }

    /**
     * The database to grant default privileges for this role.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    public readonly objectType!: pulumi.Output<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    public readonly privileges!: pulumi.Output<string[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    public readonly schema!: pulumi.Output<string | undefined>;
    /**
     * Permit the grant recipient to grant it to others
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DefaultPrivileg resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges */
    constructor(name: string, args: DefaultPrivilegArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges */
    constructor(name: string, argsOrState?: DefaultPrivilegArgs | DefaultPrivilegState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DefaultPrivileg is deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultPrivilegState | undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["objectType"] = state ? state.objectType : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["privileges"] = state ? state.privileges : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
            resourceInputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as DefaultPrivilegArgs | undefined;
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.objectType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectType'");
            }
            if ((!args || args.owner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'owner'");
            }
            if ((!args || args.privileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privileges'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["objectType"] = args ? args.objectType : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["privileges"] = args ? args.privileges : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultPrivileg.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultPrivileg resources.
 */
export interface DefaultPrivilegState {
    /**
     * The database to grant default privileges for this role.
     */
    database?: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    objectType?: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    owner?: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    privileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    role?: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    schema?: pulumi.Input<string>;
    /**
     * Permit the grant recipient to grant it to others
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DefaultPrivileg resource.
 */
export interface DefaultPrivilegArgs {
    /**
     * The database to grant default privileges for this role.
     */
    database: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    objectType: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    owner: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    privileges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    role: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    schema?: pulumi.Input<string>;
    /**
     * Permit the grant recipient to grant it to others
     */
    withGrantOption?: pulumi.Input<boolean>;
}
