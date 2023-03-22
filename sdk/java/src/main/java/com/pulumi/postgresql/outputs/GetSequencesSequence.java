// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSequencesSequence {
    /**
     * @return The sequence&#39;s data type as defined in ``information_schema.sequences``.
     * 
     */
    private String dataType;
    /**
     * @return The sequence name.
     * 
     */
    private String objectName;
    /**
     * @return The parent schema.
     * 
     */
    private String schemaName;

    private GetSequencesSequence() {}
    /**
     * @return The sequence&#39;s data type as defined in ``information_schema.sequences``.
     * 
     */
    public String dataType() {
        return this.dataType;
    }
    /**
     * @return The sequence name.
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSequencesSequence defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dataType;
        private String objectName;
        private String schemaName;
        public Builder() {}
        public Builder(GetSequencesSequence defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dataType = defaults.dataType;
    	      this.objectName = defaults.objectName;
    	      this.schemaName = defaults.schemaName;
        }

        @CustomType.Setter
        public Builder dataType(String dataType) {
            this.dataType = Objects.requireNonNull(dataType);
            return this;
        }
        @CustomType.Setter
        public Builder objectName(String objectName) {
            this.objectName = Objects.requireNonNull(objectName);
            return this;
        }
        @CustomType.Setter
        public Builder schemaName(String schemaName) {
            this.schemaName = Objects.requireNonNull(schemaName);
            return this;
        }
        public GetSequencesSequence build() {
            final var o = new GetSequencesSequence();
            o.dataType = dataType;
            o.objectName = objectName;
            o.schemaName = schemaName;
            return o;
        }
    }
}
