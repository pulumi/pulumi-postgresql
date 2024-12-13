// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.SecurityLabel`` resource creates and manages security labels.
 *
 * See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-security-label.html)
 *
 * > **Note:** This resource needs Postgresql version 11 or above.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const myRole = new postgresql.Role("my_role", {
 *     name: "my_role",
 *     login: true,
 * });
 * const workload = new postgresql.SecurityLabel("workload", {
 *     objectType: "role",
 *     objectName: myRole.name,
 *     labelProvider: "pgaadauth",
 *     label: "aadauth,oid=00000000-0000-0000-0000-000000000000,type=service",
 * });
 * ```
 *
 * ## Import
 *
 * Security label is an attribute that can be added multiple times, so no import is needed, simply apply again.
 */
export class SecurityLabel extends pulumi.CustomResource {
    /**
     * Get an existing SecurityLabel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityLabelState, opts?: pulumi.CustomResourceOptions): SecurityLabel {
        return new SecurityLabel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/securityLabel:SecurityLabel';

    /**
     * Returns true if the given object is an instance of SecurityLabel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityLabel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityLabel.__pulumiType;
    }

    /**
     * The value of the security label.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * The name of the provider with which this label is to be associated.
     */
    public readonly labelProvider!: pulumi.Output<string>;
    /**
     * The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
     */
    public readonly objectName!: pulumi.Output<string>;
    /**
     * The PostgreSQL object type to apply this security label to.
     */
    public readonly objectType!: pulumi.Output<string>;

    /**
     * Create a SecurityLabel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityLabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityLabelArgs | SecurityLabelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityLabelState | undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["labelProvider"] = state ? state.labelProvider : undefined;
            resourceInputs["objectName"] = state ? state.objectName : undefined;
            resourceInputs["objectType"] = state ? state.objectType : undefined;
        } else {
            const args = argsOrState as SecurityLabelArgs | undefined;
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            if ((!args || args.labelProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labelProvider'");
            }
            if ((!args || args.objectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectName'");
            }
            if ((!args || args.objectType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectType'");
            }
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["labelProvider"] = args ? args.labelProvider : undefined;
            resourceInputs["objectName"] = args ? args.objectName : undefined;
            resourceInputs["objectType"] = args ? args.objectType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityLabel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityLabel resources.
 */
export interface SecurityLabelState {
    /**
     * The value of the security label.
     */
    label?: pulumi.Input<string>;
    /**
     * The name of the provider with which this label is to be associated.
     */
    labelProvider?: pulumi.Input<string>;
    /**
     * The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
     */
    objectName?: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to apply this security label to.
     */
    objectType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityLabel resource.
 */
export interface SecurityLabelArgs {
    /**
     * The value of the security label.
     */
    label: pulumi.Input<string>;
    /**
     * The name of the provider with which this label is to be associated.
     */
    labelProvider: pulumi.Input<string>;
    /**
     * The name of the object to be labeled. Names of objects that reside in schemas (tables, functions, etc.) can be schema-qualified.
     */
    objectName: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to apply this security label to.
     */
    objectType: pulumi.Input<string>;
}