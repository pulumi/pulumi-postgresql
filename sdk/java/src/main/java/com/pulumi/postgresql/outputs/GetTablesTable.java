// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTablesTable {
    /**
     * @return The table name.
     * 
     */
    private String objectName;
    /**
     * @return The parent schema.
     * 
     */
    private String schemaName;
    /**
     * @return The table type as defined in ``information_schema.tables``.
     * 
     */
    private String tableType;

    private GetTablesTable() {}
    /**
     * @return The table name.
     * 
     */
    public String objectName() {
        return this.objectName;
    }
    /**
     * @return The parent schema.
     * 
     */
    public String schemaName() {
        return this.schemaName;
    }
    /**
     * @return The table type as defined in ``information_schema.tables``.
     * 
     */
    public String tableType() {
        return this.tableType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTablesTable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String objectName;
        private String schemaName;
        private String tableType;
        public Builder() {}
        public Builder(GetTablesTable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.objectName = defaults.objectName;
    	      this.schemaName = defaults.schemaName;
    	      this.tableType = defaults.tableType;
        }

        @CustomType.Setter
        public Builder objectName(String objectName) {
            if (objectName == null) {
              throw new MissingRequiredPropertyException("GetTablesTable", "objectName");
            }
            this.objectName = objectName;
            return this;
        }
        @CustomType.Setter
        public Builder schemaName(String schemaName) {
            if (schemaName == null) {
              throw new MissingRequiredPropertyException("GetTablesTable", "schemaName");
            }
            this.schemaName = schemaName;
            return this;
        }
        @CustomType.Setter
        public Builder tableType(String tableType) {
            if (tableType == null) {
              throw new MissingRequiredPropertyException("GetTablesTable", "tableType");
            }
            this.tableType = tableType;
            return this;
        }
        public GetTablesTable build() {
            final var _resultValue = new GetTablesTable();
            _resultValue.objectName = objectName;
            _resultValue.schemaName = schemaName;
            _resultValue.tableType = tableType;
            return _resultValue;
        }
    }
}
