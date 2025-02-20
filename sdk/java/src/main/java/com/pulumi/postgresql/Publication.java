// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.PublicationArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.PublicationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `postgresql.Publication` resource creates and manages a publication on a PostgreSQL
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
 * import com.pulumi.postgresql.Publication;
 * import com.pulumi.postgresql.PublicationArgs;
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
 *         var publication = new Publication("publication", PublicationArgs.builder()
 *             .name("publication")
 *             .tables(            
 *                 "public.test",
 *                 "another_schema.test")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import Example
 * 
 * Publication can be imported using this format:
 * 
 */
@ResourceType(type="postgresql:index/publication:Publication")
public class Publication extends com.pulumi.resources.CustomResource {
    /**
     * Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
     * 
     */
    @Export(name="allTables", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allTables;

    /**
     * @return Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
     * 
     */
    public Output<Boolean> allTables() {
        return this.allTables;
    }
    /**
     * Which database to create the publication on. Defaults to provider database.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output<String> database;

    /**
     * @return Which database to create the publication on. Defaults to provider database.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
     * 
     */
    @Export(name="dropCascade", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dropCascade;

    /**
     * @return Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> dropCascade() {
        return Codegen.optional(this.dropCascade);
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
     * Who owns the publication. Defaults to provider user.
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output<String> owner;

    /**
     * @return Who owns the publication. Defaults to provider user.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
     * 
     */
    @Export(name="publishParams", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> publishParams;

    /**
     * @return Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
     * 
     */
    public Output<List<String>> publishParams() {
        return this.publishParams;
    }
    /**
     * Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
     * 
     */
    @Export(name="publishViaPartitionRootParam", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> publishViaPartitionRootParam;

    /**
     * @return Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> publishViaPartitionRootParam() {
        return Codegen.optional(this.publishViaPartitionRootParam);
    }
    /**
     * Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
     * 
     */
    @Export(name="tables", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tables;

    /**
     * @return Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
     * 
     */
    public Output<List<String>> tables() {
        return this.tables;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Publication(java.lang.String name) {
        this(name, PublicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Publication(java.lang.String name, @Nullable PublicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Publication(java.lang.String name, @Nullable PublicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/publication:Publication", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Publication(java.lang.String name, Output<java.lang.String> id, @Nullable PublicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/publication:Publication", name, state, makeResourceOptions(options, id), false);
    }

    private static PublicationArgs makeArgs(@Nullable PublicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PublicationArgs.Empty : args;
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
    public static Publication get(java.lang.String name, Output<java.lang.String> id, @Nullable PublicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Publication(name, id, state, options);
    }
}
