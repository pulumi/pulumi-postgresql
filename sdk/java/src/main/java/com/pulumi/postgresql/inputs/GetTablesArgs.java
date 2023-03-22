// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTablesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTablesArgs Empty = new GetTablesArgs();

    /**
     * The PostgreSQL database which will be queried for table names.
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return The PostgreSQL database which will be queried for table names.
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
     * 
     */
    @Import(name="likeAllPatterns")
    private @Nullable Output<List<String>> likeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
     * 
     */
    public Optional<Output<List<String>>> likeAllPatterns() {
        return Optional.ofNullable(this.likeAllPatterns);
    }

    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
     * 
     */
    @Import(name="likeAnyPatterns")
    private @Nullable Output<List<String>> likeAnyPatterns;

    /**
     * @return List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
     * 
     */
    public Optional<Output<List<String>>> likeAnyPatterns() {
        return Optional.ofNullable(this.likeAnyPatterns);
    }

    /**
     * List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     * 
     */
    @Import(name="notLikeAllPatterns")
    private @Nullable Output<List<String>> notLikeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     * 
     */
    public Optional<Output<List<String>>> notLikeAllPatterns() {
        return Optional.ofNullable(this.notLikeAllPatterns);
    }

    /**
     * Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
     * 
     */
    @Import(name="regexPattern")
    private @Nullable Output<String> regexPattern;

    /**
     * @return Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
     * 
     */
    public Optional<Output<String>> regexPattern() {
        return Optional.ofNullable(this.regexPattern);
    }

    /**
     * List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
     * 
     */
    @Import(name="schemas")
    private @Nullable Output<List<String>> schemas;

    /**
     * @return List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
     * 
     */
    public Optional<Output<List<String>>> schemas() {
        return Optional.ofNullable(this.schemas);
    }

    /**
     * List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use &#39;BASE TABLE&#39; for normal tables only.
     * 
     */
    @Import(name="tableTypes")
    private @Nullable Output<List<String>> tableTypes;

    /**
     * @return List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use &#39;BASE TABLE&#39; for normal tables only.
     * 
     */
    public Optional<Output<List<String>>> tableTypes() {
        return Optional.ofNullable(this.tableTypes);
    }

    private GetTablesArgs() {}

    private GetTablesArgs(GetTablesArgs $) {
        this.database = $.database;
        this.likeAllPatterns = $.likeAllPatterns;
        this.likeAnyPatterns = $.likeAnyPatterns;
        this.notLikeAllPatterns = $.notLikeAllPatterns;
        this.regexPattern = $.regexPattern;
        this.schemas = $.schemas;
        this.tableTypes = $.tableTypes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTablesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTablesArgs $;

        public Builder() {
            $ = new GetTablesArgs();
        }

        public Builder(GetTablesArgs defaults) {
            $ = new GetTablesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The PostgreSQL database which will be queried for table names.
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The PostgreSQL database which will be queried for table names.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(@Nullable Output<List<String>> likeAllPatterns) {
            $.likeAllPatterns = likeAllPatterns;
            return this;
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(List<String> likeAllPatterns) {
            return likeAllPatterns(Output.of(likeAllPatterns));
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(String... likeAllPatterns) {
            return likeAllPatterns(List.of(likeAllPatterns));
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(@Nullable Output<List<String>> likeAnyPatterns) {
            $.likeAnyPatterns = likeAnyPatterns;
            return this;
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(List<String> likeAnyPatterns) {
            return likeAnyPatterns(Output.of(likeAnyPatterns));
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(String... likeAnyPatterns) {
            return likeAnyPatterns(List.of(likeAnyPatterns));
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(@Nullable Output<List<String>> notLikeAllPatterns) {
            $.notLikeAllPatterns = notLikeAllPatterns;
            return this;
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(List<String> notLikeAllPatterns) {
            return notLikeAllPatterns(Output.of(notLikeAllPatterns));
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(String... notLikeAllPatterns) {
            return notLikeAllPatterns(List.of(notLikeAllPatterns));
        }

        /**
         * @param regexPattern Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
         * 
         * @return builder
         * 
         */
        public Builder regexPattern(@Nullable Output<String> regexPattern) {
            $.regexPattern = regexPattern;
            return this;
        }

        /**
         * @param regexPattern Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
         * 
         * @return builder
         * 
         */
        public Builder regexPattern(String regexPattern) {
            return regexPattern(Output.of(regexPattern));
        }

        /**
         * @param schemas List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
         * 
         * @return builder
         * 
         */
        public Builder schemas(@Nullable Output<List<String>> schemas) {
            $.schemas = schemas;
            return this;
        }

        /**
         * @param schemas List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
         * 
         * @return builder
         * 
         */
        public Builder schemas(List<String> schemas) {
            return schemas(Output.of(schemas));
        }

        /**
         * @param schemas List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
         * 
         * @return builder
         * 
         */
        public Builder schemas(String... schemas) {
            return schemas(List.of(schemas));
        }

        /**
         * @param tableTypes List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use &#39;BASE TABLE&#39; for normal tables only.
         * 
         * @return builder
         * 
         */
        public Builder tableTypes(@Nullable Output<List<String>> tableTypes) {
            $.tableTypes = tableTypes;
            return this;
        }

        /**
         * @param tableTypes List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use &#39;BASE TABLE&#39; for normal tables only.
         * 
         * @return builder
         * 
         */
        public Builder tableTypes(List<String> tableTypes) {
            return tableTypes(Output.of(tableTypes));
        }

        /**
         * @param tableTypes List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use &#39;BASE TABLE&#39; for normal tables only.
         * 
         * @return builder
         * 
         */
        public Builder tableTypes(String... tableTypes) {
            return tableTypes(List.of(tableTypes));
        }

        public GetTablesArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            return $;
        }
    }

}
