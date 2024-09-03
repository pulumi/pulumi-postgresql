// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * layout: "postgresql"
 * page_title: "PostgreSQL: postgresqlSusbcription"
 * sidebar_current: "docs-postgresql-resource-postgresql_subscription"
 * description: |-
 * Creates and manages a subscription in a PostgreSQL server database.
 * <!-- yaml: line 6: could not find expected ':' -->
 *
 * # postgresql.Subscription
 *
 * The `postgresql.Subscription` resource creates and manages a publication on a PostgreSQL
 * server.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const subscription = new postgresql.Subscription("subscription", {
 *     name: "subscription",
 *     conninfo: "host=localhost port=5432 dbname=mydb user=postgres password=postgres",
 *     publications: ["publication"],
 * });
 * ```
 *
 * ## Postgres documentation
 *
 * - <https://www.postgresql.org/docs/current/sql-createsubscription.html>
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/subscription:Subscription';

    /**
     * Returns true if the given object is an instance of Subscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subscription.__pulumiType;
    }

    /**
     * The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
     */
    public readonly conninfo!: pulumi.Output<string>;
    /**
     * Specifies whether the command should create the replication slot on the publisher. Default behavior is true
     */
    public readonly createSlot!: pulumi.Output<boolean | undefined>;
    /**
     * Which database to create the subscription on. Defaults to provider database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The name of the publication.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Names of the publications on the publisher to subscribe to
     */
    public readonly publications!: pulumi.Output<string[]>;
    /**
     * Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
     */
    public readonly slotName!: pulumi.Output<string | undefined>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionState | undefined;
            resourceInputs["conninfo"] = state ? state.conninfo : undefined;
            resourceInputs["createSlot"] = state ? state.createSlot : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publications"] = state ? state.publications : undefined;
            resourceInputs["slotName"] = state ? state.slotName : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if ((!args || args.conninfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conninfo'");
            }
            if ((!args || args.publications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publications'");
            }
            resourceInputs["conninfo"] = args?.conninfo ? pulumi.secret(args.conninfo) : undefined;
            resourceInputs["createSlot"] = args ? args.createSlot : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publications"] = args ? args.publications : undefined;
            resourceInputs["slotName"] = args ? args.slotName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["conninfo"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Subscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
     */
    conninfo?: pulumi.Input<string>;
    /**
     * Specifies whether the command should create the replication slot on the publisher. Default behavior is true
     */
    createSlot?: pulumi.Input<boolean>;
    /**
     * Which database to create the subscription on. Defaults to provider database.
     */
    database?: pulumi.Input<string>;
    /**
     * The name of the publication.
     */
    name?: pulumi.Input<string>;
    /**
     * Names of the publications on the publisher to subscribe to
     */
    publications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
     */
    slotName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
     */
    conninfo: pulumi.Input<string>;
    /**
     * Specifies whether the command should create the replication slot on the publisher. Default behavior is true
     */
    createSlot?: pulumi.Input<boolean>;
    /**
     * Which database to create the subscription on. Defaults to provider database.
     */
    database?: pulumi.Input<string>;
    /**
     * The name of the publication.
     */
    name?: pulumi.Input<string>;
    /**
     * Names of the publications on the publisher to subscribe to
     */
    publications: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
     */
    slotName?: pulumi.Input<string>;
}
