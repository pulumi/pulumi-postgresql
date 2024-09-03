// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.GetSchemasArgs;
import com.pulumi.postgresql.inputs.GetSchemasPlainArgs;
import com.pulumi.postgresql.inputs.GetSequencesArgs;
import com.pulumi.postgresql.inputs.GetSequencesPlainArgs;
import com.pulumi.postgresql.inputs.GetTablesArgs;
import com.pulumi.postgresql.inputs.GetTablesPlainArgs;
import com.pulumi.postgresql.outputs.GetSchemasResult;
import com.pulumi.postgresql.outputs.GetSequencesResult;
import com.pulumi.postgresql.outputs.GetTablesResult;
import java.util.concurrent.CompletableFuture;

public final class PostgresqlFunctions {
    /**
     * The `postgresql.getSchemas` data source retrieves a list of schema names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSchemasArgs;
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
     *         final var mySchemas = PostgresqlFunctions.getSchemas(GetSchemasArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSchemasResult> getSchemas(GetSchemasArgs args) {
        return getSchemas(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getSchemas` data source retrieves a list of schema names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSchemasArgs;
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
     *         final var mySchemas = PostgresqlFunctions.getSchemas(GetSchemasArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSchemasResult> getSchemasPlain(GetSchemasPlainArgs args) {
        return getSchemasPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getSchemas` data source retrieves a list of schema names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSchemasArgs;
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
     *         final var mySchemas = PostgresqlFunctions.getSchemas(GetSchemasArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSchemasResult> getSchemas(GetSchemasArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("postgresql:index/getSchemas:getSchemas", TypeShape.of(GetSchemasResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `postgresql.getSchemas` data source retrieves a list of schema names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSchemasArgs;
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
     *         final var mySchemas = PostgresqlFunctions.getSchemas(GetSchemasArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSchemasResult> getSchemasPlain(GetSchemasPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("postgresql:index/getSchemas:getSchemas", TypeShape.of(GetSchemasResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `postgresql.getSequences` data source retrieves a list of sequence names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSequencesArgs;
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
     *         final var mySequences = PostgresqlFunctions.getSequences(GetSequencesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSequencesResult> getSequences(GetSequencesArgs args) {
        return getSequences(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getSequences` data source retrieves a list of sequence names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSequencesArgs;
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
     *         final var mySequences = PostgresqlFunctions.getSequences(GetSequencesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSequencesResult> getSequencesPlain(GetSequencesPlainArgs args) {
        return getSequencesPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getSequences` data source retrieves a list of sequence names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSequencesArgs;
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
     *         final var mySequences = PostgresqlFunctions.getSequences(GetSequencesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSequencesResult> getSequences(GetSequencesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("postgresql:index/getSequences:getSequences", TypeShape.of(GetSequencesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `postgresql.getSequences` data source retrieves a list of sequence names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetSequencesArgs;
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
     *         final var mySequences = PostgresqlFunctions.getSequences(GetSequencesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSequencesResult> getSequencesPlain(GetSequencesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("postgresql:index/getSequences:getSequences", TypeShape.of(GetSequencesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `postgresql.getTables` data source retrieves a list of table names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetTablesArgs;
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
     *         final var myTables = PostgresqlFunctions.getTables(GetTablesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTablesResult> getTables(GetTablesArgs args) {
        return getTables(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getTables` data source retrieves a list of table names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetTablesArgs;
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
     *         final var myTables = PostgresqlFunctions.getTables(GetTablesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTablesResult> getTablesPlain(GetTablesPlainArgs args) {
        return getTablesPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `postgresql.getTables` data source retrieves a list of table names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetTablesArgs;
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
     *         final var myTables = PostgresqlFunctions.getTables(GetTablesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTablesResult> getTables(GetTablesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("postgresql:index/getTables:getTables", TypeShape.of(GetTablesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `postgresql.getTables` data source retrieves a list of table names from a specified PostgreSQL database.
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
     * import com.pulumi.postgresql.PostgresqlFunctions;
     * import com.pulumi.postgresql.inputs.GetTablesArgs;
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
     *         final var myTables = PostgresqlFunctions.getTables(GetTablesArgs.builder()
     *             .database("my_database")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTablesResult> getTablesPlain(GetTablesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("postgresql:index/getTables:getTables", TypeShape.of(GetTablesResult.class), args, Utilities.withVersion(options));
    }
}
