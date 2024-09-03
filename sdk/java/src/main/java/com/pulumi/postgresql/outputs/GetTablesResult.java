// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.postgresql.outputs.GetTablesTable;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTablesResult {
    private String database;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable List<String> likeAllPatterns;
    private @Nullable List<String> likeAnyPatterns;
    private @Nullable List<String> notLikeAllPatterns;
    private @Nullable String regexPattern;
    private @Nullable List<String> schemas;
    private @Nullable List<String> tableTypes;
    /**
     * @return A list of PostgreSQL tables retrieved by this data source. Each table consists of the fields documented below.
     * ---
     * 
     */
    private List<GetTablesTable> tables;

    private GetTablesResult() {}
    public String database() {
        return this.database;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> likeAllPatterns() {
        return this.likeAllPatterns == null ? List.of() : this.likeAllPatterns;
    }
    public List<String> likeAnyPatterns() {
        return this.likeAnyPatterns == null ? List.of() : this.likeAnyPatterns;
    }
    public List<String> notLikeAllPatterns() {
        return this.notLikeAllPatterns == null ? List.of() : this.notLikeAllPatterns;
    }
    public Optional<String> regexPattern() {
        return Optional.ofNullable(this.regexPattern);
    }
    public List<String> schemas() {
        return this.schemas == null ? List.of() : this.schemas;
    }
    public List<String> tableTypes() {
        return this.tableTypes == null ? List.of() : this.tableTypes;
    }
    /**
     * @return A list of PostgreSQL tables retrieved by this data source. Each table consists of the fields documented below.
     * ---
     * 
     */
    public List<GetTablesTable> tables() {
        return this.tables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTablesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String database;
        private String id;
        private @Nullable List<String> likeAllPatterns;
        private @Nullable List<String> likeAnyPatterns;
        private @Nullable List<String> notLikeAllPatterns;
        private @Nullable String regexPattern;
        private @Nullable List<String> schemas;
        private @Nullable List<String> tableTypes;
        private List<GetTablesTable> tables;
        public Builder() {}
        public Builder(GetTablesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.database = defaults.database;
    	      this.id = defaults.id;
    	      this.likeAllPatterns = defaults.likeAllPatterns;
    	      this.likeAnyPatterns = defaults.likeAnyPatterns;
    	      this.notLikeAllPatterns = defaults.notLikeAllPatterns;
    	      this.regexPattern = defaults.regexPattern;
    	      this.schemas = defaults.schemas;
    	      this.tableTypes = defaults.tableTypes;
    	      this.tables = defaults.tables;
        }

        @CustomType.Setter
        public Builder database(String database) {
            if (database == null) {
              throw new MissingRequiredPropertyException("GetTablesResult", "database");
            }
            this.database = database;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTablesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder likeAllPatterns(@Nullable List<String> likeAllPatterns) {

            this.likeAllPatterns = likeAllPatterns;
            return this;
        }
        public Builder likeAllPatterns(String... likeAllPatterns) {
            return likeAllPatterns(List.of(likeAllPatterns));
        }
        @CustomType.Setter
        public Builder likeAnyPatterns(@Nullable List<String> likeAnyPatterns) {

            this.likeAnyPatterns = likeAnyPatterns;
            return this;
        }
        public Builder likeAnyPatterns(String... likeAnyPatterns) {
            return likeAnyPatterns(List.of(likeAnyPatterns));
        }
        @CustomType.Setter
        public Builder notLikeAllPatterns(@Nullable List<String> notLikeAllPatterns) {

            this.notLikeAllPatterns = notLikeAllPatterns;
            return this;
        }
        public Builder notLikeAllPatterns(String... notLikeAllPatterns) {
            return notLikeAllPatterns(List.of(notLikeAllPatterns));
        }
        @CustomType.Setter
        public Builder regexPattern(@Nullable String regexPattern) {

            this.regexPattern = regexPattern;
            return this;
        }
        @CustomType.Setter
        public Builder schemas(@Nullable List<String> schemas) {

            this.schemas = schemas;
            return this;
        }
        public Builder schemas(String... schemas) {
            return schemas(List.of(schemas));
        }
        @CustomType.Setter
        public Builder tableTypes(@Nullable List<String> tableTypes) {

            this.tableTypes = tableTypes;
            return this;
        }
        public Builder tableTypes(String... tableTypes) {
            return tableTypes(List.of(tableTypes));
        }
        @CustomType.Setter
        public Builder tables(List<GetTablesTable> tables) {
            if (tables == null) {
              throw new MissingRequiredPropertyException("GetTablesResult", "tables");
            }
            this.tables = tables;
            return this;
        }
        public Builder tables(GetTablesTable... tables) {
            return tables(List.of(tables));
        }
        public GetTablesResult build() {
            final var _resultValue = new GetTablesResult();
            _resultValue.database = database;
            _resultValue.id = id;
            _resultValue.likeAllPatterns = likeAllPatterns;
            _resultValue.likeAnyPatterns = likeAnyPatterns;
            _resultValue.notLikeAllPatterns = notLikeAllPatterns;
            _resultValue.regexPattern = regexPattern;
            _resultValue.schemas = schemas;
            _resultValue.tableTypes = tableTypes;
            _resultValue.tables = tables;
            return _resultValue;
        }
    }
}
