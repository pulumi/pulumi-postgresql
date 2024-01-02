// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionArgArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionArgArgs Empty = new FunctionArgArgs();

    /**
     * An expression to be used as default value if the parameter is not specified.
     * 
     */
    @Import(name="default")
    private @Nullable Output<String> default_;

    /**
     * @return An expression to be used as default value if the parameter is not specified.
     * 
     */
    public Optional<Output<String>> default_() {
        return Optional.ofNullable(this.default_);
    }

    /**
     * Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The name of the argument.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the argument.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The type of the argument.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the argument.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private FunctionArgArgs() {}

    private FunctionArgArgs(FunctionArgArgs $) {
        this.default_ = $.default_;
        this.mode = $.mode;
        this.name = $.name;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionArgArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionArgArgs $;

        public Builder() {
            $ = new FunctionArgArgs();
        }

        public Builder(FunctionArgArgs defaults) {
            $ = new FunctionArgArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param default_ An expression to be used as default value if the parameter is not specified.
         * 
         * @return builder
         * 
         */
        public Builder default_(@Nullable Output<String> default_) {
            $.default_ = default_;
            return this;
        }

        /**
         * @param default_ An expression to be used as default value if the parameter is not specified.
         * 
         * @return builder
         * 
         */
        public Builder default_(String default_) {
            return default_(Output.of(default_));
        }

        /**
         * @param mode Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name The name of the argument.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the argument.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param type The type of the argument.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the argument.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public FunctionArgArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("FunctionArgArgs", "type");
            }
            return $;
        }
    }

}
