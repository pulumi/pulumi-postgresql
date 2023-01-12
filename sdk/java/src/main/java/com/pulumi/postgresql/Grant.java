// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.GrantArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.GrantState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``postgresql.Grant`` resource creates and manages privileges given to a user for a database schema.
 * 
 * See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-grant.html)
 * 
 * &gt; **Note:** This resource needs Postgresql version 9 or above.
 * 
 * ## Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.Grant;
 * import com.pulumi.postgresql.GrantArgs;
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
 *         var readonlyTables = new Grant(&#34;readonlyTables&#34;, GrantArgs.builder()        
 *             .database(&#34;test_db&#34;)
 *             .objectType(&#34;table&#34;)
 *             .objects(            
 *                 &#34;table1&#34;,
 *                 &#34;table2&#34;)
 *             .privileges(&#34;SELECT&#34;)
 *             .role(&#34;test_role&#34;)
 *             .schema(&#34;public&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Examples
 * 
 * Revoke default accesses for public schema:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.Grant;
 * import com.pulumi.postgresql.GrantArgs;
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
 *         var revokePublic = new Grant(&#34;revokePublic&#34;, GrantArgs.builder()        
 *             .database(&#34;test_db&#34;)
 *             .objectType(&#34;schema&#34;)
 *             .privileges()
 *             .role(&#34;public&#34;)
 *             .schema(&#34;public&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="postgresql:index/grant:Grant")
public class Grant extends com.pulumi.resources.CustomResource {
    /**
     * The database to grant privileges on for this role.
     * 
     */
    @Export(name="database", type=String.class, parameters={})
    private Output<String> database;

    /**
     * @return The database to grant privileges on for this role.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
     * 
     */
    @Export(name="objectType", type=String.class, parameters={})
    private Output<String> objectType;

    /**
     * @return The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
     * 
     */
    public Output<String> objectType() {
        return this.objectType;
    }
    /**
     * The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`.
     * 
     */
    @Export(name="objects", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> objects;

    /**
     * @return The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`.
     * 
     */
    public Output<Optional<List<String>>> objects() {
        return Codegen.optional(this.objects);
    }
    /**
     * The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     * 
     */
    @Export(name="privileges", type=List.class, parameters={String.class})
    private Output<List<String>> privileges;

    /**
     * @return The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     * 
     */
    public Output<List<String>> privileges() {
        return this.privileges;
    }
    /**
     * The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
     * 
     */
    @Export(name="schema", type=String.class, parameters={})
    private Output</* @Nullable */ String> schema;

    /**
     * @return The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
     * 
     */
    public Output<Optional<String>> schema() {
        return Codegen.optional(this.schema);
    }
    /**
     * Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     * 
     */
    @Export(name="withGrantOption", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> withGrantOption;

    /**
     * @return Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> withGrantOption() {
        return Codegen.optional(this.withGrantOption);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Grant(String name) {
        this(name, GrantArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Grant(String name, GrantArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Grant(String name, GrantArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/grant:Grant", name, args == null ? GrantArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Grant(String name, Output<String> id, @Nullable GrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/grant:Grant", name, state, makeResourceOptions(options, id));
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
    public static Grant get(String name, Output<String> id, @Nullable GrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Grant(name, id, state, options);
    }
}