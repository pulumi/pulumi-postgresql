// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.Server`` resource creates and manages a foreign server on a PostgreSQL server.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const extPostgresFdw = new postgresql.Extension("ext_postgres_fdw", {name: "postgres_fdw"});
 * const myserverPostgres = new postgresql.Server("myserver_postgres", {
 *     serverName: "myserver_postgres",
 *     fdwName: "postgres_fdw",
 *     options: {
 *         host: "foo",
 *         dbname: "foodb",
 *         port: "5432",
 *     },
 * }, {
 *     dependsOn: [extPostgresFdw],
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const extFileFdw = new postgresql.Extension("ext_file_fdw", {name: "file_fdw"});
 * const myserverFile = new postgresql.Server("myserver_file", {
 *     serverName: "myserver_file",
 *     fdwName: "file_fdw",
 * }, {
 *     dependsOn: [extFileFdw],
 * });
 * ```
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
     */
    public readonly dropCascade!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the foreign-data wrapper that manages the server.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    public readonly fdwName!: pulumi.Output<string>;
    /**
     * This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
     */
    public readonly options!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the foreign server to be created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
     */
    public readonly serverOwner!: pulumi.Output<string>;
    /**
     * Optional server type, potentially useful to foreign-data wrappers.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    public readonly serverType!: pulumi.Output<string | undefined>;
    /**
     * Optional server version, potentially useful to foreign-data wrappers.
     */
    public readonly serverVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            resourceInputs["dropCascade"] = state ? state.dropCascade : undefined;
            resourceInputs["fdwName"] = state ? state.fdwName : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["serverOwner"] = state ? state.serverOwner : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["serverVersion"] = state ? state.serverVersion : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if ((!args || args.fdwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fdwName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            resourceInputs["dropCascade"] = args ? args.dropCascade : undefined;
            resourceInputs["fdwName"] = args ? args.fdwName : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["serverOwner"] = args ? args.serverOwner : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["serverVersion"] = args ? args.serverVersion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
     */
    dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the foreign-data wrapper that manages the server.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    fdwName?: pulumi.Input<string>;
    /**
     * This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the foreign server to be created.
     */
    serverName?: pulumi.Input<string>;
    /**
     * By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
     */
    serverOwner?: pulumi.Input<string>;
    /**
     * Optional server type, potentially useful to foreign-data wrappers.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Optional server version, potentially useful to foreign-data wrappers.
     */
    serverVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
     */
    dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the foreign-data wrapper that manages the server.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    fdwName: pulumi.Input<string>;
    /**
     * This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper.
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the foreign server to be created.
     */
    serverName: pulumi.Input<string>;
    /**
     * By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
     */
    serverOwner?: pulumi.Input<string>;
    /**
     * Optional server type, potentially useful to foreign-data wrappers.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Optional server version, potentially useful to foreign-data wrappers.
     */
    serverVersion?: pulumi.Input<string>;
}
