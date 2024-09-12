// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.inputs.ProviderClientcertArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * Use rds_iam instead of password authentication (see:
     * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
     * 
     */
    @Import(name="awsRdsIamAuth", json=true)
    private @Nullable Output<Boolean> awsRdsIamAuth;

    /**
     * @return Use rds_iam instead of password authentication (see:
     * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
     * 
     */
    public Optional<Output<Boolean>> awsRdsIamAuth() {
        return Optional.ofNullable(this.awsRdsIamAuth);
    }

    /**
     * AWS profile to use for IAM auth
     * 
     */
    @Import(name="awsRdsIamProfile")
    private @Nullable Output<String> awsRdsIamProfile;

    /**
     * @return AWS profile to use for IAM auth
     * 
     */
    public Optional<Output<String>> awsRdsIamProfile() {
        return Optional.ofNullable(this.awsRdsIamProfile);
    }

    /**
     * AWS region to use for IAM auth
     * 
     */
    @Import(name="awsRdsIamRegion")
    private @Nullable Output<String> awsRdsIamRegion;

    /**
     * @return AWS region to use for IAM auth
     * 
     */
    public Optional<Output<String>> awsRdsIamRegion() {
        return Optional.ofNullable(this.awsRdsIamRegion);
    }

    /**
     * Use MS Azure identity OAuth token (see:
     * https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
     * 
     */
    @Import(name="azureIdentityAuth", json=true)
    private @Nullable Output<Boolean> azureIdentityAuth;

    /**
     * @return Use MS Azure identity OAuth token (see:
     * https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
     * 
     */
    public Optional<Output<Boolean>> azureIdentityAuth() {
        return Optional.ofNullable(this.azureIdentityAuth);
    }

    @Import(name="azureTenantId")
    private @Nullable Output<String> azureTenantId;

    public Optional<Output<String>> azureTenantId() {
        return Optional.ofNullable(this.azureTenantId);
    }

    /**
     * SSL client certificate if required by the database.
     * 
     */
    @Import(name="clientcert", json=true)
    private @Nullable Output<ProviderClientcertArgs> clientcert;

    /**
     * @return SSL client certificate if required by the database.
     * 
     */
    public Optional<Output<ProviderClientcertArgs>> clientcert() {
        return Optional.ofNullable(this.clientcert);
    }

    /**
     * Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
     * 
     */
    @Import(name="connectTimeout", json=true)
    private @Nullable Output<Integer> connectTimeout;

    /**
     * @return Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
     * 
     */
    public Optional<Output<Integer>> connectTimeout() {
        return Optional.ofNullable(this.connectTimeout);
    }

    /**
     * The name of the database to connect to in order to conenct to (defaults to `postgres`).
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return The name of the database to connect to in order to conenct to (defaults to `postgres`).
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * Database username associated to the connected user (for user name maps)
     * 
     */
    @Import(name="databaseUsername")
    private @Nullable Output<String> databaseUsername;

    /**
     * @return Database username associated to the connected user (for user name maps)
     * 
     */
    public Optional<Output<String>> databaseUsername() {
        return Optional.ofNullable(this.databaseUsername);
    }

    /**
     * Specify the expected version of PostgreSQL.
     * 
     */
    @Import(name="expectedVersion")
    private @Nullable Output<String> expectedVersion;

    /**
     * @return Specify the expected version of PostgreSQL.
     * 
     */
    public Optional<Output<String>> expectedVersion() {
        return Optional.ofNullable(this.expectedVersion);
    }

    /**
     * Service account to impersonate when using GCP IAM authentication.
     * 
     */
    @Import(name="gcpIamImpersonateServiceAccount")
    private @Nullable Output<String> gcpIamImpersonateServiceAccount;

    /**
     * @return Service account to impersonate when using GCP IAM authentication.
     * 
     */
    public Optional<Output<String>> gcpIamImpersonateServiceAccount() {
        return Optional.ofNullable(this.gcpIamImpersonateServiceAccount);
    }

    /**
     * Name of PostgreSQL server address to connect to
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return Name of PostgreSQL server address to connect to
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Maximum number of connections to establish to the database. Zero means unlimited.
     * 
     */
    @Import(name="maxConnections", json=true)
    private @Nullable Output<Integer> maxConnections;

    /**
     * @return Maximum number of connections to establish to the database. Zero means unlimited.
     * 
     */
    public Optional<Output<Integer>> maxConnections() {
        return Optional.ofNullable(this.maxConnections);
    }

    /**
     * Password to be used if the PostgreSQL server demands password authentication
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password to be used if the PostgreSQL server demands password authentication
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
     * 
     */
    @Import(name="port", json=true)
    private @Nullable Output<Integer> port;

    /**
     * @return The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    @Import(name="scheme")
    private @Nullable Output<String> scheme;

    public Optional<Output<String>> scheme() {
        return Optional.ofNullable(this.scheme);
    }

    /**
     * @deprecated
     * Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     * 
     */
    @Deprecated /* Rename PostgreSQL provider `ssl_mode` attribute to `sslmode` */
    @Import(name="sslMode")
    private @Nullable Output<String> sslMode;

    /**
     * @deprecated
     * Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     * 
     */
    @Deprecated /* Rename PostgreSQL provider `ssl_mode` attribute to `sslmode` */
    public Optional<Output<String>> sslMode() {
        return Optional.ofNullable(this.sslMode);
    }

    /**
     * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     * 
     */
    @Import(name="sslmode")
    private @Nullable Output<String> sslmode;

    /**
     * @return This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     * 
     */
    public Optional<Output<String>> sslmode() {
        return Optional.ofNullable(this.sslmode);
    }

    /**
     * The SSL server root certificate file path. The file must contain PEM encoded data.
     * 
     */
    @Import(name="sslrootcert")
    private @Nullable Output<String> sslrootcert;

    /**
     * @return The SSL server root certificate file path. The file must contain PEM encoded data.
     * 
     */
    public Optional<Output<String>> sslrootcert() {
        return Optional.ofNullable(this.sslrootcert);
    }

    /**
     * Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
     * Refreshing state password from Postgres)
     * 
     */
    @Import(name="superuser", json=true)
    private @Nullable Output<Boolean> superuser;

    /**
     * @return Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
     * Refreshing state password from Postgres)
     * 
     */
    public Optional<Output<Boolean>> superuser() {
        return Optional.ofNullable(this.superuser);
    }

    /**
     * PostgreSQL user name to connect as
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return PostgreSQL user name to connect as
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.awsRdsIamAuth = $.awsRdsIamAuth;
        this.awsRdsIamProfile = $.awsRdsIamProfile;
        this.awsRdsIamRegion = $.awsRdsIamRegion;
        this.azureIdentityAuth = $.azureIdentityAuth;
        this.azureTenantId = $.azureTenantId;
        this.clientcert = $.clientcert;
        this.connectTimeout = $.connectTimeout;
        this.database = $.database;
        this.databaseUsername = $.databaseUsername;
        this.expectedVersion = $.expectedVersion;
        this.gcpIamImpersonateServiceAccount = $.gcpIamImpersonateServiceAccount;
        this.host = $.host;
        this.maxConnections = $.maxConnections;
        this.password = $.password;
        this.port = $.port;
        this.scheme = $.scheme;
        this.sslMode = $.sslMode;
        this.sslmode = $.sslmode;
        this.sslrootcert = $.sslrootcert;
        this.superuser = $.superuser;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsRdsIamAuth Use rds_iam instead of password authentication (see:
         * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamAuth(@Nullable Output<Boolean> awsRdsIamAuth) {
            $.awsRdsIamAuth = awsRdsIamAuth;
            return this;
        }

        /**
         * @param awsRdsIamAuth Use rds_iam instead of password authentication (see:
         * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamAuth(Boolean awsRdsIamAuth) {
            return awsRdsIamAuth(Output.of(awsRdsIamAuth));
        }

        /**
         * @param awsRdsIamProfile AWS profile to use for IAM auth
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamProfile(@Nullable Output<String> awsRdsIamProfile) {
            $.awsRdsIamProfile = awsRdsIamProfile;
            return this;
        }

        /**
         * @param awsRdsIamProfile AWS profile to use for IAM auth
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamProfile(String awsRdsIamProfile) {
            return awsRdsIamProfile(Output.of(awsRdsIamProfile));
        }

        /**
         * @param awsRdsIamRegion AWS region to use for IAM auth
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamRegion(@Nullable Output<String> awsRdsIamRegion) {
            $.awsRdsIamRegion = awsRdsIamRegion;
            return this;
        }

        /**
         * @param awsRdsIamRegion AWS region to use for IAM auth
         * 
         * @return builder
         * 
         */
        public Builder awsRdsIamRegion(String awsRdsIamRegion) {
            return awsRdsIamRegion(Output.of(awsRdsIamRegion));
        }

        /**
         * @param azureIdentityAuth Use MS Azure identity OAuth token (see:
         * https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
         * 
         * @return builder
         * 
         */
        public Builder azureIdentityAuth(@Nullable Output<Boolean> azureIdentityAuth) {
            $.azureIdentityAuth = azureIdentityAuth;
            return this;
        }

        /**
         * @param azureIdentityAuth Use MS Azure identity OAuth token (see:
         * https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
         * 
         * @return builder
         * 
         */
        public Builder azureIdentityAuth(Boolean azureIdentityAuth) {
            return azureIdentityAuth(Output.of(azureIdentityAuth));
        }

        public Builder azureTenantId(@Nullable Output<String> azureTenantId) {
            $.azureTenantId = azureTenantId;
            return this;
        }

        public Builder azureTenantId(String azureTenantId) {
            return azureTenantId(Output.of(azureTenantId));
        }

        /**
         * @param clientcert SSL client certificate if required by the database.
         * 
         * @return builder
         * 
         */
        public Builder clientcert(@Nullable Output<ProviderClientcertArgs> clientcert) {
            $.clientcert = clientcert;
            return this;
        }

        /**
         * @param clientcert SSL client certificate if required by the database.
         * 
         * @return builder
         * 
         */
        public Builder clientcert(ProviderClientcertArgs clientcert) {
            return clientcert(Output.of(clientcert));
        }

        /**
         * @param connectTimeout Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(@Nullable Output<Integer> connectTimeout) {
            $.connectTimeout = connectTimeout;
            return this;
        }

        /**
         * @param connectTimeout Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(Integer connectTimeout) {
            return connectTimeout(Output.of(connectTimeout));
        }

        /**
         * @param database The name of the database to connect to in order to conenct to (defaults to `postgres`).
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The name of the database to connect to in order to conenct to (defaults to `postgres`).
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param databaseUsername Database username associated to the connected user (for user name maps)
         * 
         * @return builder
         * 
         */
        public Builder databaseUsername(@Nullable Output<String> databaseUsername) {
            $.databaseUsername = databaseUsername;
            return this;
        }

        /**
         * @param databaseUsername Database username associated to the connected user (for user name maps)
         * 
         * @return builder
         * 
         */
        public Builder databaseUsername(String databaseUsername) {
            return databaseUsername(Output.of(databaseUsername));
        }

        /**
         * @param expectedVersion Specify the expected version of PostgreSQL.
         * 
         * @return builder
         * 
         */
        public Builder expectedVersion(@Nullable Output<String> expectedVersion) {
            $.expectedVersion = expectedVersion;
            return this;
        }

        /**
         * @param expectedVersion Specify the expected version of PostgreSQL.
         * 
         * @return builder
         * 
         */
        public Builder expectedVersion(String expectedVersion) {
            return expectedVersion(Output.of(expectedVersion));
        }

        /**
         * @param gcpIamImpersonateServiceAccount Service account to impersonate when using GCP IAM authentication.
         * 
         * @return builder
         * 
         */
        public Builder gcpIamImpersonateServiceAccount(@Nullable Output<String> gcpIamImpersonateServiceAccount) {
            $.gcpIamImpersonateServiceAccount = gcpIamImpersonateServiceAccount;
            return this;
        }

        /**
         * @param gcpIamImpersonateServiceAccount Service account to impersonate when using GCP IAM authentication.
         * 
         * @return builder
         * 
         */
        public Builder gcpIamImpersonateServiceAccount(String gcpIamImpersonateServiceAccount) {
            return gcpIamImpersonateServiceAccount(Output.of(gcpIamImpersonateServiceAccount));
        }

        /**
         * @param host Name of PostgreSQL server address to connect to
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Name of PostgreSQL server address to connect to
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param maxConnections Maximum number of connections to establish to the database. Zero means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder maxConnections(@Nullable Output<Integer> maxConnections) {
            $.maxConnections = maxConnections;
            return this;
        }

        /**
         * @param maxConnections Maximum number of connections to establish to the database. Zero means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder maxConnections(Integer maxConnections) {
            return maxConnections(Output.of(maxConnections));
        }

        /**
         * @param password Password to be used if the PostgreSQL server demands password authentication
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password to be used if the PostgreSQL server demands password authentication
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param port The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public Builder scheme(@Nullable Output<String> scheme) {
            $.scheme = scheme;
            return this;
        }

        public Builder scheme(String scheme) {
            return scheme(Output.of(scheme));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
         * 
         */
        @Deprecated /* Rename PostgreSQL provider `ssl_mode` attribute to `sslmode` */
        public Builder sslMode(@Nullable Output<String> sslMode) {
            $.sslMode = sslMode;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
         * 
         */
        @Deprecated /* Rename PostgreSQL provider `ssl_mode` attribute to `sslmode` */
        public Builder sslMode(String sslMode) {
            return sslMode(Output.of(sslMode));
        }

        /**
         * @param sslmode This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
         * PostgreSQL server
         * 
         * @return builder
         * 
         */
        public Builder sslmode(@Nullable Output<String> sslmode) {
            $.sslmode = sslmode;
            return this;
        }

        /**
         * @param sslmode This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
         * PostgreSQL server
         * 
         * @return builder
         * 
         */
        public Builder sslmode(String sslmode) {
            return sslmode(Output.of(sslmode));
        }

        /**
         * @param sslrootcert The SSL server root certificate file path. The file must contain PEM encoded data.
         * 
         * @return builder
         * 
         */
        public Builder sslrootcert(@Nullable Output<String> sslrootcert) {
            $.sslrootcert = sslrootcert;
            return this;
        }

        /**
         * @param sslrootcert The SSL server root certificate file path. The file must contain PEM encoded data.
         * 
         * @return builder
         * 
         */
        public Builder sslrootcert(String sslrootcert) {
            return sslrootcert(Output.of(sslrootcert));
        }

        /**
         * @param superuser Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
         * Refreshing state password from Postgres)
         * 
         * @return builder
         * 
         */
        public Builder superuser(@Nullable Output<Boolean> superuser) {
            $.superuser = superuser;
            return this;
        }

        /**
         * @param superuser Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
         * Refreshing state password from Postgres)
         * 
         * @return builder
         * 
         */
        public Builder superuser(Boolean superuser) {
            return superuser(Output.of(superuser));
        }

        /**
         * @param username PostgreSQL user name to connect as
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username PostgreSQL user name to connect as
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderArgs build() {
            $.connectTimeout = Codegen.integerProp("connectTimeout").output().arg($.connectTimeout).env("PGCONNECT_TIMEOUT").def(180).getNullable();
            $.sslmode = Codegen.stringProp("sslmode").output().arg($.sslmode).env("PGSSLMODE").getNullable();
            return $;
        }
    }

}
