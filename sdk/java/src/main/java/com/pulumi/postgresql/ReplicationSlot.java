// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.ReplicationSlotArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.ReplicationSlotState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The ``postgresql.ReplicationSlot`` resource creates and manages a replication slot on a PostgreSQL
 * server.
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
 * import com.pulumi.postgresql.ReplicationSlot;
 * import com.pulumi.postgresql.ReplicationSlotArgs;
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
 *         var mySlot = new ReplicationSlot("mySlot", ReplicationSlotArgs.builder()
 *             .name("my_slot")
 *             .plugin("test_decoding")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="postgresql:index/replicationSlot:ReplicationSlot")
public class ReplicationSlot extends com.pulumi.resources.CustomResource {
    /**
     * Which database to create the replication slot on. Defaults to provider database.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output<String> database;

    /**
     * @return Which database to create the replication slot on. Defaults to provider database.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * The name of the replication slot.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the replication slot.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Sets the output plugin.
     * 
     */
    @Export(name="plugin", refs={String.class}, tree="[0]")
    private Output<String> plugin;

    /**
     * @return Sets the output plugin.
     * 
     */
    public Output<String> plugin() {
        return this.plugin;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReplicationSlot(java.lang.String name) {
        this(name, ReplicationSlotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicationSlot(java.lang.String name, ReplicationSlotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicationSlot(java.lang.String name, ReplicationSlotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/replicationSlot:ReplicationSlot", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReplicationSlot(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicationSlotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/replicationSlot:ReplicationSlot", name, state, makeResourceOptions(options, id), false);
    }

    private static ReplicationSlotArgs makeArgs(ReplicationSlotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReplicationSlotArgs.Empty : args;
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
    public static ReplicationSlot get(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicationSlotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicationSlot(name, id, state, options);
    }
}
