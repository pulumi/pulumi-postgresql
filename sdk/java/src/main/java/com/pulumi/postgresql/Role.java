// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.RoleArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.RoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="postgresql:index/role:Role")
public class Role extends com.pulumi.resources.CustomResource {
    /**
     * Defines the role to switch to at login via [`SET ROLE`](https://www.postgresql.org/docs/current/sql-set-role.html).
     * 
     */
    @Export(name="assumeRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> assumeRole;

    /**
     * @return Defines the role to switch to at login via [`SET ROLE`](https://www.postgresql.org/docs/current/sql-set-role.html).
     * 
     */
    public Output<Optional<String>> assumeRole() {
        return Codegen.optional(this.assumeRole);
    }
    /**
     * Defines whether a role bypasses every
     * row-level security (RLS) policy.  Default value is `false`.
     * 
     */
    @Export(name="bypassRowLevelSecurity", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bypassRowLevelSecurity;

    /**
     * @return Defines whether a role bypasses every
     * row-level security (RLS) policy.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> bypassRowLevelSecurity() {
        return Codegen.optional(this.bypassRowLevelSecurity);
    }
    /**
     * If this role can log in, this specifies how
     * many concurrent connections the role can establish. `-1` (the default) means no
     * limit.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> connectionLimit;

    /**
     * @return If this role can log in, this specifies how
     * many concurrent connections the role can establish. `-1` (the default) means no
     * limit.
     * 
     */
    public Output<Optional<Integer>> connectionLimit() {
        return Codegen.optional(this.connectionLimit);
    }
    /**
     * Defines a role&#39;s ability to execute `CREATE
     * DATABASE`.  Default value is `false`.
     * 
     */
    @Export(name="createDatabase", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createDatabase;

    /**
     * @return Defines a role&#39;s ability to execute `CREATE
     * DATABASE`.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> createDatabase() {
        return Codegen.optional(this.createDatabase);
    }
    /**
     * Defines a role&#39;s ability to execute `CREATE ROLE`.
     * A role with this privilege can also alter and drop other roles.  Default value
     * is `false`.
     * 
     */
    @Export(name="createRole", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createRole;

    /**
     * @return Defines a role&#39;s ability to execute `CREATE ROLE`.
     * A role with this privilege can also alter and drop other roles.  Default value
     * is `false`.
     * 
     */
    public Output<Optional<Boolean>> createRole() {
        return Codegen.optional(this.createRole);
    }
    /**
     * @deprecated
     * Rename PostgreSQL role resource attribute &#34;encrypted&#34; to &#34;encrypted_password&#34;
     * 
     */
    @Deprecated /* Rename PostgreSQL role resource attribute ""encrypted"" to ""encrypted_password"" */
    @Export(name="encrypted", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encrypted;

    public Output<Optional<String>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    /**
     * Defines whether the password is stored
     * encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
     * is always set (to the conservative and safe value), but may interfere with the
     * behavior of
     * [PostgreSQL&#39;s `password_encryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
     * 
     */
    @Export(name="encryptedPassword", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encryptedPassword;

    /**
     * @return Defines whether the password is stored
     * encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
     * is always set (to the conservative and safe value), but may interfere with the
     * behavior of
     * [PostgreSQL&#39;s `password_encryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
     * 
     */
    public Output<Optional<Boolean>> encryptedPassword() {
        return Codegen.optional(this.encryptedPassword);
    }
    /**
     * Terminate any session with an open transaction that has been idle for longer than the specified duration in milliseconds
     * 
     */
    @Export(name="idleInTransactionSessionTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> idleInTransactionSessionTimeout;

    /**
     * @return Terminate any session with an open transaction that has been idle for longer than the specified duration in milliseconds
     * 
     */
    public Output<Optional<Integer>> idleInTransactionSessionTimeout() {
        return Codegen.optional(this.idleInTransactionSessionTimeout);
    }
    /**
     * Defines whether a role &#34;inherits&#34; the privileges of
     * roles it is a member of.  Default value is `true`.
     * 
     */
    @Export(name="inherit", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> inherit;

    /**
     * @return Defines whether a role &#34;inherits&#34; the privileges of
     * roles it is a member of.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> inherit() {
        return Codegen.optional(this.inherit);
    }
    /**
     * Defines whether role is allowed to log in.  Roles without
     * this attribute are useful for managing database privileges, but are not users
     * in the usual sense of the word.  Default value is `false`.
     * 
     */
    @Export(name="login", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> login;

    /**
     * @return Defines whether role is allowed to log in.  Roles without
     * this attribute are useful for managing database privileges, but are not users
     * in the usual sense of the word.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> login() {
        return Codegen.optional(this.login);
    }
    /**
     * The name of the role. Must be unique on the PostgreSQL
     * server instance where it is configured.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the role. Must be unique on the PostgreSQL
     * server instance where it is configured.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Sets the role&#39;s password. A password is only of use
     * for roles having the `login` attribute set to true.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Sets the role&#39;s password. A password is only of use
     * for roles having the `login` attribute set to true.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Defines whether a role is allowed to initiate
     * streaming replication or put the system in and out of backup mode.  Default
     * value is `false`
     * 
     */
    @Export(name="replication", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> replication;

    /**
     * @return Defines whether a role is allowed to initiate
     * streaming replication or put the system in and out of backup mode.  Default
     * value is `false`
     * 
     */
    public Output<Optional<Boolean>> replication() {
        return Codegen.optional(this.replication);
    }
    /**
     * Defines list of roles which will be granted to this new role.
     * 
     */
    @Export(name="roles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> roles;

    /**
     * @return Defines list of roles which will be granted to this new role.
     * 
     */
    public Output<Optional<List<String>>> roles() {
        return Codegen.optional(this.roles);
    }
    /**
     * Alters the search path of this new role. Note that
     * due to limitations in the implementation, values cannot contain the substring
     * `&#34;, &#34;`.
     * 
     */
    @Export(name="searchPaths", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> searchPaths;

    /**
     * @return Alters the search path of this new role. Note that
     * due to limitations in the implementation, values cannot contain the substring
     * `&#34;, &#34;`.
     * 
     */
    public Output<Optional<List<String>>> searchPaths() {
        return Codegen.optional(this.searchPaths);
    }
    /**
     * When a PostgreSQL ROLE exists in multiple
     * databases and the ROLE is dropped, the
     * [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
     * in each of the respective databases must occur before the ROLE can be dropped
     * from the catalog.  Set this option to true when there are multiple databases
     * in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
     * This is the third and final step taken when removing a ROLE from a database.
     * 
     */
    @Export(name="skipDropRole", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipDropRole;

    /**
     * @return When a PostgreSQL ROLE exists in multiple
     * databases and the ROLE is dropped, the
     * [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
     * in each of the respective databases must occur before the ROLE can be dropped
     * from the catalog.  Set this option to true when there are multiple databases
     * in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
     * This is the third and final step taken when removing a ROLE from a database.
     * 
     */
    public Output<Optional<Boolean>> skipDropRole() {
        return Codegen.optional(this.skipDropRole);
    }
    /**
     * When a PostgreSQL ROLE exists in multiple
     * databases and the ROLE is dropped, a
     * [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
     * must be executed on each of the respective databases before the `DROP ROLE`
     * can be executed to dropped the ROLE from the catalog.  This is the first and
     * second steps taken when removing a ROLE from a database (the second step being
     * an implicit
     * [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
     * 
     */
    @Export(name="skipReassignOwned", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipReassignOwned;

    /**
     * @return When a PostgreSQL ROLE exists in multiple
     * databases and the ROLE is dropped, a
     * [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
     * must be executed on each of the respective databases before the `DROP ROLE`
     * can be executed to dropped the ROLE from the catalog.  This is the first and
     * second steps taken when removing a ROLE from a database (the second step being
     * an implicit
     * [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
     * 
     */
    public Output<Optional<Boolean>> skipReassignOwned() {
        return Codegen.optional(this.skipReassignOwned);
    }
    /**
     * Defines [`statement_timeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
     * 
     */
    @Export(name="statementTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> statementTimeout;

    /**
     * @return Defines [`statement_timeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
     * 
     */
    public Output<Optional<Integer>> statementTimeout() {
        return Codegen.optional(this.statementTimeout);
    }
    /**
     * Defines whether the role is a &#34;superuser&#34;, and
     * therefore can override all access restrictions within the database.  Default
     * value is `false`.
     * 
     */
    @Export(name="superuser", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> superuser;

    /**
     * @return Defines whether the role is a &#34;superuser&#34;, and
     * therefore can override all access restrictions within the database.  Default
     * value is `false`.
     * 
     */
    public Output<Optional<Boolean>> superuser() {
        return Codegen.optional(this.superuser);
    }
    /**
     * Defines the date and time after which the role&#39;s
     * password is no longer valid.  Established connections past this `valid_time`
     * will have to be manually terminated.  This value corresponds to a PostgreSQL
     * datetime. If omitted or the magic value `NULL` is used, `valid_until` will be
     * set to `infinity`.  Default is `NULL`, therefore `infinity`.
     * 
     */
    @Export(name="validUntil", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> validUntil;

    /**
     * @return Defines the date and time after which the role&#39;s
     * password is no longer valid.  Established connections past this `valid_time`
     * will have to be manually terminated.  This value corresponds to a PostgreSQL
     * datetime. If omitted or the magic value `NULL` is used, `valid_until` will be
     * set to `infinity`.  Default is `NULL`, therefore `infinity`.
     * 
     */
    public Output<Optional<String>> validUntil() {
        return Codegen.optional(this.validUntil);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Role(java.lang.String name) {
        this(name, RoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Role(java.lang.String name, @Nullable RoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Role(java.lang.String name, @Nullable RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/role:Role", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Role(java.lang.String name, Output<java.lang.String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/role:Role", name, state, makeResourceOptions(options, id), false);
    }

    private static RoleArgs makeArgs(@Nullable RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RoleArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static Role get(java.lang.String name, Output<java.lang.String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Role(name, id, state, options);
    }
}
