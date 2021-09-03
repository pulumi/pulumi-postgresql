// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.Extension`` resource creates and manages an extension on a PostgreSQL
 * server.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const myExtension = new postgresql.Extension("my_extension", {});
 * ```
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionState, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/extension:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    /**
     * When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
     */
    public readonly createCascade!: pulumi.Output<boolean | undefined>;
    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
     */
    public readonly dropCascade!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the extension.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Sets the schema of an extension.
     */
    public readonly schema!: pulumi.Output<string>;
    /**
     * Sets the version number of the extension.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionArgs | ExtensionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionState | undefined;
            inputs["createCascade"] = state ? state.createCascade : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["dropCascade"] = state ? state.dropCascade : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ExtensionArgs | undefined;
            inputs["createCascade"] = args ? args.createCascade : undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["dropCascade"] = args ? args.dropCascade : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Extension.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extension resources.
 */
export interface ExtensionState {
    /**
     * When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
     */
    readonly createCascade?: pulumi.Input<boolean>;
    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
     */
    readonly dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the extension.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets the schema of an extension.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * Sets the version number of the extension.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    /**
     * When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
     */
    readonly createCascade?: pulumi.Input<boolean>;
    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
     */
    readonly dropCascade?: pulumi.Input<boolean>;
    /**
     * The name of the extension.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets the schema of an extension.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * Sets the version number of the extension.
     */
    readonly version?: pulumi.Input<string>;
}
