// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.UserMappingArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.UserMappingState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `postgresql.UserMapping` resource creates and manages a user mapping on a PostgreSQL server.
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
 * import com.pulumi.postgresql.Extension;
 * import com.pulumi.postgresql.ExtensionArgs;
 * import com.pulumi.postgresql.Server;
 * import com.pulumi.postgresql.ServerArgs;
 * import com.pulumi.postgresql.Role;
 * import com.pulumi.postgresql.RoleArgs;
 * import com.pulumi.postgresql.UserMapping;
 * import com.pulumi.postgresql.UserMappingArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var extPostgresFdw = new Extension("extPostgresFdw", ExtensionArgs.builder()
 *             .name("postgres_fdw")
 *             .build());
 * 
 *         var myserverPostgres = new Server("myserverPostgres", ServerArgs.builder()
 *             .serverName("myserver_postgres")
 *             .fdwName("postgres_fdw")
 *             .options(Map.ofEntries(
 *                 Map.entry("host", "foo"),
 *                 Map.entry("dbname", "foodb"),
 *                 Map.entry("port", "5432")
 *             ))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(extPostgresFdw)
 *                 .build());
 * 
 *         var remote = new Role("remote", RoleArgs.builder()
 *             .name("remote")
 *             .build());
 * 
 *         var remoteUserMapping = new UserMapping("remoteUserMapping", UserMappingArgs.builder()
 *             .serverName(myserverPostgres.serverName())
 *             .userName(remote.name())
 *             .options(Map.ofEntries(
 *                 Map.entry("user", "admin"),
 *                 Map.entry("password", "pass")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="postgresql:index/userMapping:UserMapping")
public class UserMapping extends com.pulumi.resources.CustomResource {
    /**
     * This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server&#39;s foreign-data wrapper.
     * 
     */
    @Export(name="options", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> options;

    /**
     * @return This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server&#39;s foreign-data wrapper.
     * 
     */
    public Output<Optional<Map<String,String>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * The name of an existing server for which the user mapping is to be created.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     * 
     */
    @Export(name="serverName", refs={String.class}, tree="[0]")
    private Output<String> serverName;

    /**
     * @return The name of an existing server for which the user mapping is to be created.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     * 
     */
    public Output<String> serverName() {
        return this.serverName;
    }
    /**
     * The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    /**
     * @return The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserMapping(java.lang.String name) {
        this(name, UserMappingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserMapping(java.lang.String name, UserMappingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserMapping(java.lang.String name, UserMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/userMapping:UserMapping", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UserMapping(java.lang.String name, Output<java.lang.String> id, @Nullable UserMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/userMapping:UserMapping", name, state, makeResourceOptions(options, id), false);
    }

    private static UserMappingArgs makeArgs(UserMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserMappingArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static UserMapping get(java.lang.String name, Output<java.lang.String> id, @Nullable UserMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserMapping(name, id, state, options);
    }
}
