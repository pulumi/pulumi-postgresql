// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.FunctionArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.FunctionState;
import com.pulumi.postgresql.outputs.FunctionArg;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * layout: &#34;postgresql&#34;
 * page_title: &#34;PostgreSQL: postgresql.Function&#34;
 * sidebar_current: &#34;docs-postgresql-resource-postgresql_function&#34;
 * description: |-
 * Creates and manages a function on a PostgreSQL server.
 * &lt;!-- yaml: line 6: could not find expected &#39;:&#39; --&gt;
 * 
 * # postgresql\_function
 * 
 * The `postgresql.Function` resource creates and manages a function on a PostgreSQL
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
 * import com.pulumi.postgresql.Function;
 * import com.pulumi.postgresql.FunctionArgs;
 * import com.pulumi.postgresql.inputs.FunctionArgArgs;
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
 *         var increment = new Function("increment", FunctionArgs.builder()
 *             .name("increment")
 *             .args(FunctionArgArgs.builder()
 *                 .name("i")
 *                 .type("integer")
 *                 .build())
 *             .returns("integer")
 *             .language("plpgsql")
 *             .body("""
 * BEGIN
 *     RETURN i + 1;
 * END;
 *             """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * It is possible to import a `postgresql_function` resource with the following
 * command:
 * 
 * ```sh
 * $ pulumi import postgresql:index/function:Function function_foo &#34;my_database.my_schema.my_function_name(arguments)&#34;
 * ```
 * Where `my_database` is the name of the database containing the schema,
 * `my_schema` is the name of the schema in the PostgreSQL database, `my_function_name` is the function name to be imported, `arguments` is the argument signature of the function including all non OUT types and
 * `postgresql_schema.function_foo` is the name of the resource whose state will be
 * populated as a result of the command.
 * 
 */
@ResourceType(type="postgresql:index/function:Function")
public class Function extends com.pulumi.resources.CustomResource {
    /**
     * List of arguments for the function.
     * 
     */
    @Export(name="args", refs={List.class,FunctionArg.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FunctionArg>> args;

    /**
     * @return List of arguments for the function.
     * 
     */
    public Output<Optional<List<FunctionArg>>> args() {
        return Codegen.optional(this.args);
    }
    /**
     * Function body.
     * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
     * 
     */
    @Export(name="body", refs={String.class}, tree="[0]")
    private Output<String> body;

    /**
     * @return Function body.
     * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
     * 
     */
    public Output<String> body() {
        return this.body;
    }
    /**
     * The database where the function is located.
     * If not specified, the function is created in the current database.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output<String> database;

    /**
     * @return The database where the function is located.
     * If not specified, the function is created in the current database.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * True to automatically drop objects that depend on the function (such as
     * operators or triggers), and in turn all objects that depend on those objects. Default is false.
     * 
     */
    @Export(name="dropCascade", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dropCascade;

    /**
     * @return True to automatically drop objects that depend on the function (such as
     * operators or triggers), and in turn all objects that depend on those objects. Default is false.
     * 
     */
    public Output<Optional<Boolean>> dropCascade() {
        return Codegen.optional(this.dropCascade);
    }
    /**
     * The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
     * 
     */
    @Export(name="language", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> language;

    /**
     * @return The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
     * 
     */
    public Output<Optional<String>> language() {
        return Codegen.optional(this.language);
    }
    /**
     * The name of the function.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the function.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
     * 
     */
    @Export(name="parallel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parallel;

    /**
     * @return Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
     * 
     */
    public Output<Optional<String>> parallel() {
        return Codegen.optional(this.parallel);
    }
    /**
     * Type that the function returns. It can be computed from the OUT arguments. Default is void.
     * 
     */
    @Export(name="returns", refs={String.class}, tree="[0]")
    private Output<String> returns;

    /**
     * @return Type that the function returns. It can be computed from the OUT arguments. Default is void.
     * 
     */
    public Output<String> returns() {
        return this.returns;
    }
    /**
     * The schema where the function is located.
     * If not specified, the function is created in the current schema.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The schema where the function is located.
     * If not specified, the function is created in the current schema.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
     * 
     */
    @Export(name="securityDefiner", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> securityDefiner;

    /**
     * @return If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
     * 
     */
    public Output<Optional<Boolean>> securityDefiner() {
        return Codegen.optional(this.securityDefiner);
    }
    /**
     * If the function should always return NULL when any of the inputs is NULL. Default is false.
     * 
     */
    @Export(name="strict", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> strict;

    /**
     * @return If the function should always return NULL when any of the inputs is NULL. Default is false.
     * 
     */
    public Output<Optional<Boolean>> strict() {
        return Codegen.optional(this.strict);
    }
    /**
     * Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
     * 
     */
    @Export(name="volatility", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volatility;

    /**
     * @return Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
     * 
     */
    public Output<Optional<String>> volatility() {
        return Codegen.optional(this.volatility);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Function(java.lang.String name) {
        this(name, FunctionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Function(java.lang.String name, FunctionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Function(java.lang.String name, FunctionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/function:Function", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Function(java.lang.String name, Output<java.lang.String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/function:Function", name, state, makeResourceOptions(options, id), false);
    }

    private static FunctionArgs makeArgs(FunctionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FunctionArgs.Empty : args;
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
    public static Function get(java.lang.String name, Output<java.lang.String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Function(name, id, state, options);
    }
}
