// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.DefaultPrivilegesArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.DefaultPrivilegesState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``postgresql.DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.
 * 
 * &gt; **Note:** This resource needs Postgresql version 9 or above.
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
 * import com.pulumi.postgresql.DefaultPrivileges;
 * import com.pulumi.postgresql.DefaultPrivilegesArgs;
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
 *         var readOnlyTables = new DefaultPrivileges("readOnlyTables", DefaultPrivilegesArgs.builder()
 *             .role("test_role")
 *             .database("test_db")
 *             .schema("public")
 *             .owner("db_owner")
 *             .objectType("table")
 *             .privileges("SELECT")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Examples
 * 
 * Revoke default privileges for functions for &#34;public&#34; role:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.DefaultPrivileges;
 * import com.pulumi.postgresql.DefaultPrivilegesArgs;
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
 *         var revokePublic = new DefaultPrivileges("revokePublic", DefaultPrivilegesArgs.builder()
 *             .database(exampleDb.name())
 *             .role("public")
 *             .owner("object_owner")
 *             .objectType("function")
 *             .privileges()
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="postgresql:index/defaultPrivileges:DefaultPrivileges")
public class DefaultPrivileges extends com.pulumi.resources.CustomResource {
    /**
     * The database to grant default privileges for this role.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output<String> database;

    /**
     * @return The database to grant default privileges for this role.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     * 
     */
    @Export(name="objectType", refs={String.class}, tree="[0]")
    private Output<String> objectType;

    /**
     * @return The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     * 
     */
    public Output<String> objectType() {
        return this.objectType;
    }
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output<String> owner;

    /**
     * @return Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     * 
     */
    @Export(name="privileges", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> privileges;

    /**
     * @return The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     * 
     */
    public Output<List<String>> privileges() {
        return this.privileges;
    }
    /**
     * The name of the role to which grant default privileges on.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name of the role to which grant default privileges on.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The database schema to set default privileges for this role.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> schema;

    /**
     * @return The database schema to set default privileges for this role.
     * 
     */
    public Output<Optional<String>> schema() {
        return Codegen.optional(this.schema);
    }
    /**
     * Permit the grant recipient to grant it to others
     * 
     */
    @Export(name="withGrantOption", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> withGrantOption;

    /**
     * @return Permit the grant recipient to grant it to others
     * 
     */
    public Output<Optional<Boolean>> withGrantOption() {
        return Codegen.optional(this.withGrantOption);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DefaultPrivileges(String name) {
        this(name, DefaultPrivilegesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DefaultPrivileges(String name, DefaultPrivilegesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DefaultPrivileges(String name, DefaultPrivilegesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/defaultPrivileges:DefaultPrivileges", name, args == null ? DefaultPrivilegesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DefaultPrivileges(String name, Output<String> id, @Nullable DefaultPrivilegesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/defaultPrivileges:DefaultPrivileges", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("postgresql:index/defaultPrivileg:DefaultPrivileg").build())
            ))
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
    public static DefaultPrivileges get(String name, Output<String> id, @Nullable DefaultPrivilegesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DefaultPrivileges(name, id, state, options);
    }
}
