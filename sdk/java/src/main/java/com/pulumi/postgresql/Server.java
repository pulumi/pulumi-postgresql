// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.ServerArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.ServerState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``postgresql.Server`` resource creates and manages a foreign server on a PostgreSQL server.
 * 
 * ## Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.Extension;
 * import com.pulumi.postgresql.ExtensionArgs;
 * import com.pulumi.postgresql.Server;
 * import com.pulumi.postgresql.ServerArgs;
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
 *         var extPostgresFdw = new Extension(&#34;extPostgresFdw&#34;, ExtensionArgs.builder()        
 *             .name(&#34;postgres_fdw&#34;)
 *             .build());
 * 
 *         var myserverPostgres = new Server(&#34;myserverPostgres&#34;, ServerArgs.builder()        
 *             .serverName(&#34;myserver_postgres&#34;)
 *             .fdwName(&#34;postgres_fdw&#34;)
 *             .options(Map.ofEntries(
 *                 Map.entry(&#34;host&#34;, &#34;foo&#34;),
 *                 Map.entry(&#34;dbname&#34;, &#34;foodb&#34;),
 *                 Map.entry(&#34;port&#34;, &#34;5432&#34;)
 *             ))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(extPostgresFdw)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.postgresql.Extension;
 * import com.pulumi.postgresql.ExtensionArgs;
 * import com.pulumi.postgresql.Server;
 * import com.pulumi.postgresql.ServerArgs;
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
 *         var extFileFdw = new Extension(&#34;extFileFdw&#34;, ExtensionArgs.builder()        
 *             .name(&#34;file_fdw&#34;)
 *             .build());
 * 
 *         var myserverFile = new Server(&#34;myserverFile&#34;, ServerArgs.builder()        
 *             .serverName(&#34;myserver_file&#34;)
 *             .fdwName(&#34;file_fdw&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(extFileFdw)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="postgresql:index/server:Server")
public class Server extends com.pulumi.resources.CustomResource {
    /**
     * When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
     * 
     */
    @Export(name="dropCascade", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dropCascade;

    /**
     * @return When true, will drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects . (Default: false)
     * 
     */
    public Output<Optional<Boolean>> dropCascade() {
        return Codegen.optional(this.dropCascade);
    }
    /**
     * The name of the foreign-data wrapper that manages the server.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     * 
     */
    @Export(name="fdwName", refs={String.class}, tree="[0]")
    private Output<String> fdwName;

    /**
     * @return The name of the foreign-data wrapper that manages the server.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     * 
     */
    public Output<String> fdwName() {
        return this.fdwName;
    }
    /**
     * This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server&#39;s foreign-data wrapper.
     * 
     */
    @Export(name="options", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> options;

    /**
     * @return This clause specifies the options for the server. The options typically define the connection details of the server, but the actual names and values are dependent on the server&#39;s foreign-data wrapper.
     * 
     */
    public Output<Optional<Map<String,String>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * The name of the foreign server to be created.
     * 
     */
    @Export(name="serverName", refs={String.class}, tree="[0]")
    private Output<String> serverName;

    /**
     * @return The name of the foreign server to be created.
     * 
     */
    public Output<String> serverName() {
        return this.serverName;
    }
    /**
     * By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
     * 
     */
    @Export(name="serverOwner", refs={String.class}, tree="[0]")
    private Output<String> serverOwner;

    /**
     * @return By default, the user who defines the server becomes its owner. Set this value to configure the new owner of the foreign server.
     * 
     */
    public Output<String> serverOwner() {
        return this.serverOwner;
    }
    /**
     * Optional server type, potentially useful to foreign-data wrappers.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     * 
     */
    @Export(name="serverType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverType;

    /**
     * @return Optional server type, potentially useful to foreign-data wrappers.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the foreign server is created.
     * 
     */
    public Output<Optional<String>> serverType() {
        return Codegen.optional(this.serverType);
    }
    /**
     * Optional server version, potentially useful to foreign-data wrappers.
     * 
     */
    @Export(name="serverVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverVersion;

    /**
     * @return Optional server version, potentially useful to foreign-data wrappers.
     * 
     */
    public Output<Optional<String>> serverVersion() {
        return Codegen.optional(this.serverVersion);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Server(String name) {
        this(name, ServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Server(String name, ServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Server(String name, ServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/server:Server", name, args == null ? ServerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Server(String name, Output<String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/server:Server", name, state, makeResourceOptions(options, id));
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
    public static Server get(String name, Output<String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Server(name, id, state, options);
    }
}
