// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.config.inputs.Clientcert;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("postgresql");
/**
 * Use rds_iam instead of password authentication (see:
 * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
 * 
 */
    public Optional<Boolean> awsRdsIamAuth() {
        return Codegen.booleanProp("awsRdsIamAuth").config(config).get();
    }
/**
 * AWS profile to use for IAM auth
 * 
 */
    public Optional<String> awsRdsIamProfile() {
        return Codegen.stringProp("awsRdsIamProfile").config(config).get();
    }
/**
 * AWS IAM role to assume for IAM auth
 * 
 */
    public Optional<String> awsRdsIamProviderRoleArn() {
        return Codegen.stringProp("awsRdsIamProviderRoleArn").config(config).get();
    }
/**
 * AWS region to use for IAM auth
 * 
 */
    public Optional<String> awsRdsIamRegion() {
        return Codegen.stringProp("awsRdsIamRegion").config(config).get();
    }
/**
 * Use MS Azure identity OAuth token (see:
 * https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
 * 
 */
    public Optional<Boolean> azureIdentityAuth() {
        return Codegen.booleanProp("azureIdentityAuth").config(config).get();
    }
    public Optional<String> azureTenantId() {
        return Codegen.stringProp("azureTenantId").config(config).get();
    }
/**
 * SSL client certificate if required by the database.
 * 
 */
    public Optional<Clientcert> clientcert() {
        return Codegen.objectProp("clientcert", Clientcert.class).config(config).get();
    }
/**
 * Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
 * 
 */
    public Optional<Integer> connectTimeout() {
        return Codegen.integerProp("connectTimeout").config(config).env("PGCONNECT_TIMEOUT").def(180).get();
    }
/**
 * The name of the database to connect to in order to conenct to (defaults to `postgres`).
 * 
 */
    public Optional<String> database() {
        return Codegen.stringProp("database").config(config).get();
    }
/**
 * Database username associated to the connected user (for user name maps)
 * 
 */
    public Optional<String> databaseUsername() {
        return Codegen.stringProp("databaseUsername").config(config).get();
    }
/**
 * Specify the expected version of PostgreSQL.
 * 
 */
    public Optional<String> expectedVersion() {
        return Codegen.stringProp("expectedVersion").config(config).get();
    }
/**
 * Service account to impersonate when using GCP IAM authentication.
 * 
 */
    public Optional<String> gcpIamImpersonateServiceAccount() {
        return Codegen.stringProp("gcpIamImpersonateServiceAccount").config(config).get();
    }
/**
 * Name of PostgreSQL server address to connect to
 * 
 */
    public Optional<String> host() {
        return Codegen.stringProp("host").config(config).get();
    }
/**
 * Maximum number of connections to establish to the database. Zero means unlimited.
 * 
 */
    public Optional<Integer> maxConnections() {
        return Codegen.integerProp("maxConnections").config(config).get();
    }
/**
 * Password to be used if the PostgreSQL server demands password authentication
 * 
 */
    public Optional<String> password() {
        return Codegen.stringProp("password").config(config).get();
    }
/**
 * The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
 * 
 */
    public Optional<Integer> port() {
        return Codegen.integerProp("port").config(config).get();
    }
    public Optional<String> scheme() {
        return Codegen.stringProp("scheme").config(config).get();
    }
    public Optional<String> sslMode() {
        return Codegen.stringProp("sslMode").config(config).get();
    }
/**
 * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
 * PostgreSQL server
 * 
 */
    public Optional<String> sslmode() {
        return Codegen.stringProp("sslmode").config(config).env("PGSSLMODE").get();
    }
/**
 * The SSL server root certificate file path. The file must contain PEM encoded data.
 * 
 */
    public Optional<String> sslrootcert() {
        return Codegen.stringProp("sslrootcert").config(config).get();
    }
/**
 * Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
 * Refreshing state password from Postgres)
 * 
 */
    public Optional<Boolean> superuser() {
        return Codegen.booleanProp("superuser").config(config).get();
    }
/**
 * PostgreSQL user name to connect as
 * 
 */
    public Optional<String> username() {
        return Codegen.stringProp("username").config(config).get();
    }
}
