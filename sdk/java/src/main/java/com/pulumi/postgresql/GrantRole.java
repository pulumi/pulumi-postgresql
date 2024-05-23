// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.GrantRoleArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.GrantRoleState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``postgresql.GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.
 * 
 * When using ``postgresql.GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.
 * 
 * &gt; **Note:** This resource needs PostgreSQL version 9 or above.
 * 
 * ## Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.GrantRole;
 * import com.pulumi.postgresql.GrantRoleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var grantRoot = new GrantRole("grantRoot", GrantRoleArgs.builder()
 *             .role("root")
 *             .grantRole("application")
 *             .withAdminOption(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; **Note:** If you use `postgresql.GrantRole` for a role that you also manage with a `postgresql.Role` resource, you need to ignore the changes of the `roles` attribute in the `postgresql.Role` resource or they will fight over what your role grants should be. e.g.:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.Role;
 * import com.pulumi.postgresql.RoleArgs;
 * import com.pulumi.postgresql.GrantRole;
 * import com.pulumi.postgresql.GrantRoleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var bob = new Role("bob", RoleArgs.builder()
 *             .name("bob")
 *             .build());
 * 
 *         var bobAdmin = new GrantRole("bobAdmin", GrantRoleArgs.builder()
 *             .role("bob")
 *             .grantRole("admin")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="postgresql:index/grantRole:GrantRole")
public class GrantRole extends com.pulumi.resources.CustomResource {
    /**
     * The name of the role that is added to `role`.
     * 
     */
    @Export(name="grantRole", refs={String.class}, tree="[0]")
    private Output<String> grantRole;

    /**
     * @return The name of the role that is added to `role`.
     * 
     */
    public Output<String> grantRole() {
        return this.grantRole;
    }
    /**
     * The name of the role that is granted a new membership.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name of the role that is granted a new membership.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * Giving ability to grant membership to others or not for `role`. (Default: false)
     * 
     */
    @Export(name="withAdminOption", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> withAdminOption;

    /**
     * @return Giving ability to grant membership to others or not for `role`. (Default: false)
     * 
     */
    public Output<Optional<Boolean>> withAdminOption() {
        return Codegen.optional(this.withAdminOption);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GrantRole(String name) {
        this(name, GrantRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GrantRole(String name, GrantRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GrantRole(String name, GrantRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/grantRole:GrantRole", name, args == null ? GrantRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GrantRole(String name, Output<String> id, @Nullable GrantRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/grantRole:GrantRole", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static GrantRole get(String name, Output<String> id, @Nullable GrantRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GrantRole(name, id, state, options);
    }
}
