// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSequencesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSequencesPlainArgs Empty = new GetSequencesPlainArgs();

    /**
     * The PostgreSQL database which will be queried for sequence names.
     * 
     */
    @Import(name="database", required=true)
    private String database;

    /**
     * @return The PostgreSQL database which will be queried for sequence names.
     * 
     */
    public String database() {
        return this.database;
    }

    /**
     * List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ALL` operators.
     * 
     */
    @Import(name="likeAllPatterns")
    private @Nullable List<String> likeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ALL` operators.
     * 
     */
    public Optional<List<String>> likeAllPatterns() {
        return Optional.ofNullable(this.likeAllPatterns);
    }

    /**
     * List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ANY` operators.
     * 
     */
    @Import(name="likeAnyPatterns")
    private @Nullable List<String> likeAnyPatterns;

    /**
     * @return List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ANY` operators.
     * 
     */
    public Optional<List<String>> likeAnyPatterns() {
        return Optional.ofNullable(this.likeAnyPatterns);
    }

    /**
     * List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `NOT LIKE ALL` operators.
     * 
     */
    @Import(name="notLikeAllPatterns")
    private @Nullable List<String> notLikeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `NOT LIKE ALL` operators.
     * 
     */
    public Optional<List<String>> notLikeAllPatterns() {
        return Optional.ofNullable(this.notLikeAllPatterns);
    }

    /**
     * Expression which will be pattern matched against sequence names in the query using the PostgreSQL `~` (regular expression match) operator.
     * 
     * Note that all optional arguments can be used in conjunction.
     * 
     */
    @Import(name="regexPattern")
    private @Nullable String regexPattern;

    /**
     * @return Expression which will be pattern matched against sequence names in the query using the PostgreSQL `~` (regular expression match) operator.
     * 
     * Note that all optional arguments can be used in conjunction.
     * 
     */
    public Optional<String> regexPattern() {
        return Optional.ofNullable(this.regexPattern);
    }

    /**
     * List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
     * 
     */
    @Import(name="schemas")
    private @Nullable List<String> schemas;

    /**
     * @return List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
     * 
     */
    public Optional<List<String>> schemas() {
        return Optional.ofNullable(this.schemas);
    }

    private GetSequencesPlainArgs() {}

    private GetSequencesPlainArgs(GetSequencesPlainArgs $) {
        this.database = $.database;
        this.likeAllPatterns = $.likeAllPatterns;
        this.likeAnyPatterns = $.likeAnyPatterns;
        this.notLikeAllPatterns = $.notLikeAllPatterns;
        this.regexPattern = $.regexPattern;
        this.schemas = $.schemas;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSequencesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSequencesPlainArgs $;

        public Builder() {
            $ = new GetSequencesPlainArgs();
        }

        public Builder(GetSequencesPlainArgs defaults) {
            $ = new GetSequencesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The PostgreSQL database which will be queried for sequence names.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            $.database = database;
            return this;
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ALL` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(@Nullable List<String> likeAllPatterns) {
            $.likeAllPatterns = likeAllPatterns;
            return this;
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ALL` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(String... likeAllPatterns) {
            return likeAllPatterns(List.of(likeAllPatterns));
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ANY` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(@Nullable List<String> likeAnyPatterns) {
            $.likeAnyPatterns = likeAnyPatterns;
            return this;
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `LIKE ANY` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(String... likeAnyPatterns) {
            return likeAnyPatterns(List.of(likeAnyPatterns));
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `NOT LIKE ALL` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(@Nullable List<String> notLikeAllPatterns) {
            $.notLikeAllPatterns = notLikeAllPatterns;
            return this;
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched against sequence names in the query using the PostgreSQL `NOT LIKE ALL` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(String... notLikeAllPatterns) {
            return notLikeAllPatterns(List.of(notLikeAllPatterns));
        }

        /**
         * @param regexPattern Expression which will be pattern matched against sequence names in the query using the PostgreSQL `~` (regular expression match) operator.
         * 
         * Note that all optional arguments can be used in conjunction.
         * 
         * @return builder
         * 
         */
        public Builder regexPattern(@Nullable String regexPattern) {
            $.regexPattern = regexPattern;
            return this;
        }

        /**
         * @param schemas List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
         * 
         * @return builder
         * 
         */
        public Builder schemas(@Nullable List<String> schemas) {
            $.schemas = schemas;
            return this;
        }

        /**
         * @param schemas List of PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
         * 
         * @return builder
         * 
         */
        public Builder schemas(String... schemas) {
            return schemas(List.of(schemas));
        }

        public GetSequencesPlainArgs build() {
            if ($.database == null) {
                throw new MissingRequiredPropertyException("GetSequencesPlainArgs", "database");
            }
            return $;
        }
    }

}
