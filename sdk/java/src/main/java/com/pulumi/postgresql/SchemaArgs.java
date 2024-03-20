// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.postgresql.inputs.SchemaPolicyArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SchemaArgs extends com.pulumi.resources.ResourceArgs {

    public static final SchemaArgs Empty = new SchemaArgs();

    /**
     * The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * When true, will also drop all the objects that are contained in the schema. (Default: false)
     * 
     */
    @Import(name="dropCascade")
    private @Nullable Output<Boolean> dropCascade;

    /**
     * @return When true, will also drop all the objects that are contained in the schema. (Default: false)
     * 
     */
    public Optional<Output<Boolean>> dropCascade() {
        return Optional.ofNullable(this.dropCascade);
    }

    /**
     * When true, use the existing schema if it exists. (Default: true)
     * 
     */
    @Import(name="ifNotExists")
    private @Nullable Output<Boolean> ifNotExists;

    /**
     * @return When true, use the existing schema if it exists. (Default: true)
     * 
     */
    public Optional<Output<Boolean>> ifNotExists() {
        return Optional.ofNullable(this.ifNotExists);
    }

    /**
     * The name of the schema. Must be unique in the PostgreSQL
     * database instance where it is configured.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the schema. Must be unique in the PostgreSQL
     * database instance where it is configured.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ROLE who owns the schema.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return The ROLE who owns the schema.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Can be specified multiple times for each policy.  Each
     * policy block supports fields documented below.
     * 
     * @deprecated
     * Use postgresql.Grant resource instead (with object_type=&#34;schema&#34;)
     * 
     */
    @Deprecated /* Use postgresql.Grant resource instead (with object_type=""schema"") */
    @Import(name="policies")
    private @Nullable Output<List<SchemaPolicyArgs>> policies;

    /**
     * @return Can be specified multiple times for each policy.  Each
     * policy block supports fields documented below.
     * 
     * @deprecated
     * Use postgresql.Grant resource instead (with object_type=&#34;schema&#34;)
     * 
     */
    @Deprecated /* Use postgresql.Grant resource instead (with object_type=""schema"") */
    public Optional<Output<List<SchemaPolicyArgs>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    private SchemaArgs() {}

    private SchemaArgs(SchemaArgs $) {
        this.database = $.database;
        this.dropCascade = $.dropCascade;
        this.ifNotExists = $.ifNotExists;
        this.name = $.name;
        this.owner = $.owner;
        this.policies = $.policies;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SchemaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SchemaArgs $;

        public Builder() {
            $ = new SchemaArgs();
        }

        public Builder(SchemaArgs defaults) {
            $ = new SchemaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param dropCascade When true, will also drop all the objects that are contained in the schema. (Default: false)
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(@Nullable Output<Boolean> dropCascade) {
            $.dropCascade = dropCascade;
            return this;
        }

        /**
         * @param dropCascade When true, will also drop all the objects that are contained in the schema. (Default: false)
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(Boolean dropCascade) {
            return dropCascade(Output.of(dropCascade));
        }

        /**
         * @param ifNotExists When true, use the existing schema if it exists. (Default: true)
         * 
         * @return builder
         * 
         */
        public Builder ifNotExists(@Nullable Output<Boolean> ifNotExists) {
            $.ifNotExists = ifNotExists;
            return this;
        }

        /**
         * @param ifNotExists When true, use the existing schema if it exists. (Default: true)
         * 
         * @return builder
         * 
         */
        public Builder ifNotExists(Boolean ifNotExists) {
            return ifNotExists(Output.of(ifNotExists));
        }

        /**
         * @param name The name of the schema. Must be unique in the PostgreSQL
         * database instance where it is configured.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the schema. Must be unique in the PostgreSQL
         * database instance where it is configured.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param owner The ROLE who owns the schema.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The ROLE who owns the schema.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param policies Can be specified multiple times for each policy.  Each
         * policy block supports fields documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * Use postgresql.Grant resource instead (with object_type=&#34;schema&#34;)
         * 
         */
        @Deprecated /* Use postgresql.Grant resource instead (with object_type=""schema"") */
        public Builder policies(@Nullable Output<List<SchemaPolicyArgs>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies Can be specified multiple times for each policy.  Each
         * policy block supports fields documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * Use postgresql.Grant resource instead (with object_type=&#34;schema&#34;)
         * 
         */
        @Deprecated /* Use postgresql.Grant resource instead (with object_type=""schema"") */
        public Builder policies(List<SchemaPolicyArgs> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies Can be specified multiple times for each policy.  Each
         * policy block supports fields documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * Use postgresql.Grant resource instead (with object_type=&#34;schema&#34;)
         * 
         */
        @Deprecated /* Use postgresql.Grant resource instead (with object_type=""schema"") */
        public Builder policies(SchemaPolicyArgs... policies) {
            return policies(List.of(policies));
        }

        public SchemaArgs build() {
            return $;
        }
    }

}
