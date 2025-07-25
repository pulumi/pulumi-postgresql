// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `postgresql.Publication` resource creates and manages a publication on a PostgreSQL
 * server.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const publication = new postgresql.Publication("publication", {
 *     name: "publication",
 *     tables: [
 *         "public.test",
 *         "another_schema.test",
 *     ],
 * });
 * ```
 *
 * ## Import Example
 *
 * Publication can be imported using this format:
 */
export class Publication extends pulumi.CustomResource {
    /**
     * Get an existing Publication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicationState, opts?: pulumi.CustomResourceOptions): Publication {
        return new Publication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/publication:Publication';

    /**
     * Returns true if the given object is an instance of Publication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Publication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Publication.__pulumiType;
    }

    /**
     * Should be ALL TABLES added to the publication. Defaults to 'false'
     */
    public readonly allTables!: pulumi.Output<boolean>;
    /**
     * Which database to create the publication on. Defaults to provider database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * Should all subsequent resources of the publication be dropped. Defaults to 'false'
     */
    public readonly dropCascade!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the publication.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Who owns the publication. Defaults to provider user.
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * Which 'publish' options should be turned on. Default to 'insert','update','delete'
     */
    public readonly publishParams!: pulumi.Output<string[]>;
    /**
     * Should be option 'publish_via_partition_root' be turned on. Default to 'false'
     */
    public readonly publishViaPartitionRootParam!: pulumi.Output<boolean | undefined>;
    /**
     * Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
     */
    public readonly tables!: pulumi.Output<string[]>;

    /**
     * Create a Publication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PublicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicationArgs | PublicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicationState | undefined;
            resourceInputs["allTables"] = state ? state.allTables : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["dropCascade"] = state ? state.dropCascade : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["publishParams"] = state ? state.publishParams : undefined;
            resourceInputs["publishViaPartitionRootParam"] = state ? state.publishViaPartitionRootParam : undefined;
            resourceInputs["tables"] = state ? state.tables : undefined;
        } else {
            const args = argsOrState as PublicationArgs | undefined;
            resourceInputs["allTables"] = args ? args.allTables : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["dropCascade"] = args ? args.dropCascade : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["publishParams"] = args ? args.publishParams : undefined;
            resourceInputs["publishViaPartitionRootParam"] = args ? args.publishViaPartitionRootParam : undefined;
            resourceInputs["tables"] = args ? args.tables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Publication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Publication resources.
 */
export interface PublicationState {
    /**
     * Should be ALL TABLES added to the publication. Defaults to 'false'
     */
    allTables?: pulumi.Input<boolean>;
    /**
     * Which database to create the publication on. Defaults to provider database.
     */
    database?: pulumi.Input<string>;
    /**
     * Should all subsequent resources of the publication be dropped. Defaults to 'false'
     */
    dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the publication.
     */
    name?: pulumi.Input<string>;
    /**
     * Who owns the publication. Defaults to provider user.
     */
    owner?: pulumi.Input<string>;
    /**
     * Which 'publish' options should be turned on. Default to 'insert','update','delete'
     */
    publishParams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should be option 'publish_via_partition_root' be turned on. Default to 'false'
     */
    publishViaPartitionRootParam?: pulumi.Input<boolean>;
    /**
     * Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
     */
    tables?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Publication resource.
 */
export interface PublicationArgs {
    /**
     * Should be ALL TABLES added to the publication. Defaults to 'false'
     */
    allTables?: pulumi.Input<boolean>;
    /**
     * Which database to create the publication on. Defaults to provider database.
     */
    database?: pulumi.Input<string>;
    /**
     * Should all subsequent resources of the publication be dropped. Defaults to 'false'
     */
    dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the publication.
     */
    name?: pulumi.Input<string>;
    /**
     * Who owns the publication. Defaults to provider user.
     */
    owner?: pulumi.Input<string>;
    /**
     * Which 'publish' options should be turned on. Default to 'insert','update','delete'
     */
    publishParams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should be option 'publish_via_partition_root' be turned on. Default to 'false'
     */
    publishViaPartitionRootParam?: pulumi.Input<boolean>;
    /**
     * Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
     */
    tables?: pulumi.Input<pulumi.Input<string>[]>;
}
