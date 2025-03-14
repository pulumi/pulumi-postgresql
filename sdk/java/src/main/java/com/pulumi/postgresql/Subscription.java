// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.SubscriptionArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.SubscriptionState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `postgresql.Subscription` resource creates and manages a publication on a PostgreSQL
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
 * import com.pulumi.postgresql.Subscription;
 * import com.pulumi.postgresql.SubscriptionArgs;
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
 *         var subscription = new Subscription("subscription", SubscriptionArgs.builder()
 *             .name("subscription")
 *             .conninfo("host=localhost port=5432 dbname=mydb user=postgres password=postgres")
 *             .publications("publication")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Postgres documentation
 * 
 * - https://www.postgresql.org/docs/current/sql-createsubscription.html
 * 
 */
@ResourceType(type="postgresql:index/subscription:Subscription")
public class Subscription extends com.pulumi.resources.CustomResource {
    /**
     * The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
     * 
     */
    @Export(name="conninfo", refs={String.class}, tree="[0]")
    private Output<String> conninfo;

    /**
     * @return The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
     * 
     */
    public Output<String> conninfo() {
        return this.conninfo;
    }
    /**
     * Specifies whether the command should create the replication slot on the publisher. Default behavior is true
     * 
     */
    @Export(name="createSlot", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createSlot;

    /**
     * @return Specifies whether the command should create the replication slot on the publisher. Default behavior is true
     * 
     */
    public Output<Optional<Boolean>> createSlot() {
        return Codegen.optional(this.createSlot);
    }
    /**
     * Which database to create the subscription on. Defaults to provider database.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output<String> database;

    /**
     * @return Which database to create the subscription on. Defaults to provider database.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * The name of the publication.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the publication.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Names of the publications on the publisher to subscribe to
     * 
     */
    @Export(name="publications", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> publications;

    /**
     * @return Names of the publications on the publisher to subscribe to
     * 
     */
    public Output<List<String>> publications() {
        return this.publications;
    }
    /**
     * Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
     * 
     */
    @Export(name="slotName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> slotName;

    /**
     * @return Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
     * 
     */
    public Output<Optional<String>> slotName() {
        return Codegen.optional(this.slotName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Subscription(java.lang.String name) {
        this(name, SubscriptionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Subscription(java.lang.String name, SubscriptionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Subscription(java.lang.String name, SubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/subscription:Subscription", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Subscription(java.lang.String name, Output<java.lang.String> id, @Nullable SubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/subscription:Subscription", name, state, makeResourceOptions(options, id), false);
    }

    private static SubscriptionArgs makeArgs(SubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SubscriptionArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "conninfo"
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
    public static Subscription get(java.lang.String name, Output<java.lang.String> id, @Nullable SubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Subscription(name, id, state, options);
    }
}
