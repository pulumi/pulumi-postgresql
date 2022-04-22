// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.
 *
 * When using ``postgresql.GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.
 *
 * > **Note:** This resource needs PostgreSQL version 9 or above.
 */
export class GrantRole extends pulumi.CustomResource {
    /**
     * Get an existing GrantRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrantRoleState, opts?: pulumi.CustomResourceOptions): GrantRole {
        return new GrantRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/grantRole:GrantRole';

    /**
     * Returns true if the given object is an instance of GrantRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrantRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrantRole.__pulumiType;
    }

    /**
     * The name of the role that is added to `role`.
     */
    public readonly grantRole!: pulumi.Output<string>;
    /**
     * The name of the role that is granted a new membership.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Giving ability to grant membership to others or not for `role`. (Default: false)
     */
    public readonly withAdminOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GrantRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrantRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrantRoleArgs | GrantRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrantRoleState | undefined;
            resourceInputs["grantRole"] = state ? state.grantRole : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["withAdminOption"] = state ? state.withAdminOption : undefined;
        } else {
            const args = argsOrState as GrantRoleArgs | undefined;
            if ((!args || args.grantRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'grantRole'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["grantRole"] = args ? args.grantRole : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["withAdminOption"] = args ? args.withAdminOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrantRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrantRole resources.
 */
export interface GrantRoleState {
    /**
     * The name of the role that is added to `role`.
     */
    grantRole?: pulumi.Input<string>;
    /**
     * The name of the role that is granted a new membership.
     */
    role?: pulumi.Input<string>;
    /**
     * Giving ability to grant membership to others or not for `role`. (Default: false)
     */
    withAdminOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GrantRole resource.
 */
export interface GrantRoleArgs {
    /**
     * The name of the role that is added to `role`.
     */
    grantRole: pulumi.Input<string>;
    /**
     * The name of the role that is granted a new membership.
     */
    role: pulumi.Input<string>;
    /**
     * Giving ability to grant membership to others or not for `role`. (Default: false)
     */
    withAdminOption?: pulumi.Input<boolean>;
}
