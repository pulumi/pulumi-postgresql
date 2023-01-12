// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DefaultPrivilegState extends com.pulumi.resources.ResourceArgs {

    public static final DefaultPrivilegState Empty = new DefaultPrivilegState();

    /**
     * The database to grant default privileges for this role.
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return The database to grant default privileges for this role.
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
     * 
     */
    @Import(name="objectType")
    private @Nullable Output<String> objectType;

    /**
     * @return The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
     * 
     */
    public Optional<Output<String>> objectType() {
        return Optional.ofNullable(this.objectType);
    }

    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     * 
     */
    @Import(name="privileges")
    private @Nullable Output<List<String>> privileges;

    /**
     * @return The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     * 
     */
    public Optional<Output<List<String>>> privileges() {
        return Optional.ofNullable(this.privileges);
    }

    /**
     * The name of the role to which grant default privileges on.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The name of the role to which grant default privileges on.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * The database schema to set default privileges for this role.
     * 
     */
    @Import(name="schema")
    private @Nullable Output<String> schema;

    /**
     * @return The database schema to set default privileges for this role.
     * 
     */
    public Optional<Output<String>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * Permit the grant recipient to grant it to others
     * 
     */
    @Import(name="withGrantOption")
    private @Nullable Output<Boolean> withGrantOption;

    /**
     * @return Permit the grant recipient to grant it to others
     * 
     */
    public Optional<Output<Boolean>> withGrantOption() {
        return Optional.ofNullable(this.withGrantOption);
    }

    private DefaultPrivilegState() {}

    private DefaultPrivilegState(DefaultPrivilegState $) {
        this.database = $.database;
        this.objectType = $.objectType;
        this.owner = $.owner;
        this.privileges = $.privileges;
        this.role = $.role;
        this.schema = $.schema;
        this.withGrantOption = $.withGrantOption;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DefaultPrivilegState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DefaultPrivilegState $;

        public Builder() {
            $ = new DefaultPrivilegState();
        }

        public Builder(DefaultPrivilegState defaults) {
            $ = new DefaultPrivilegState(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The database to grant default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The database to grant default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param objectType The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
         * 
         * @return builder
         * 
         */
        public Builder objectType(@Nullable Output<String> objectType) {
            $.objectType = objectType;
            return this;
        }

        /**
         * @param objectType The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
         * 
         * @return builder
         * 
         */
        public Builder objectType(String objectType) {
            return objectType(Output.of(objectType));
        }

        /**
         * @param owner Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param privileges The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(@Nullable Output<List<String>> privileges) {
            $.privileges = privileges;
            return this;
        }

        /**
         * @param privileges The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(List<String> privileges) {
            return privileges(Output.of(privileges));
        }

        /**
         * @param privileges The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(String... privileges) {
            return privileges(List.of(privileges));
        }

        /**
         * @param role The name of the role to which grant default privileges on.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role to which grant default privileges on.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param schema The database schema to set default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<String> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema The database schema to set default privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder schema(String schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param withGrantOption Permit the grant recipient to grant it to others
         * 
         * @return builder
         * 
         */
        public Builder withGrantOption(@Nullable Output<Boolean> withGrantOption) {
            $.withGrantOption = withGrantOption;
            return this;
        }

        /**
         * @param withGrantOption Permit the grant recipient to grant it to others
         * 
         * @return builder
         * 
         */
        public Builder withGrantOption(Boolean withGrantOption) {
            return withGrantOption(Output.of(withGrantOption));
        }

        public DefaultPrivilegState build() {
            return $;
        }
    }

}