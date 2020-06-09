// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the postgresql package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'postgresql';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        inputs["clientcert"] = pulumi.output(args ? args.clientcert : undefined).apply(JSON.stringify);
        inputs["connectTimeout"] = pulumi.output((args ? args.connectTimeout : undefined) || (<any>utilities.getEnvNumber("PGCONNECT_TIMEOUT") || 180)).apply(JSON.stringify);
        inputs["database"] = (args ? args.database : undefined) || (utilities.getEnv("PGDATABASE") || "postgres");
        inputs["databaseUsername"] = args ? args.databaseUsername : undefined;
        inputs["expectedVersion"] = args ? args.expectedVersion : undefined;
        inputs["host"] = (args ? args.host : undefined) || utilities.getEnv("PGHOST");
        inputs["maxConnections"] = pulumi.output(args ? args.maxConnections : undefined).apply(JSON.stringify);
        inputs["password"] = (args ? args.password : undefined) || utilities.getEnv("PGPASSWORD");
        inputs["port"] = pulumi.output((args ? args.port : undefined) || (<any>utilities.getEnvNumber("PGPORT") || 5432)).apply(JSON.stringify);
        inputs["sslMode"] = args ? args.sslMode : undefined;
        inputs["sslmode"] = (args ? args.sslmode : undefined) || utilities.getEnv("PGSSLMODE");
        inputs["sslrootcert"] = args ? args.sslrootcert : undefined;
        inputs["superuser"] = pulumi.output(args ? args.superuser : undefined).apply(JSON.stringify);
        inputs["username"] = (args ? args.username : undefined) || (utilities.getEnv("PGUSER") || "postgres");
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * SSL client certificate if required by the database.
     */
    readonly clientcert?: pulumi.Input<inputs.ProviderClientcert>;
    /**
     * Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
     */
    readonly connectTimeout?: pulumi.Input<number>;
    /**
     * The name of the database to connect to in order to conenct to (defaults to `postgres`).
     */
    readonly database?: pulumi.Input<string>;
    /**
     * Database username associated to the connected user (for user name maps)
     */
    readonly databaseUsername?: pulumi.Input<string>;
    /**
     * Specify the expected version of PostgreSQL.
     */
    readonly expectedVersion?: pulumi.Input<string>;
    /**
     * Name of PostgreSQL server address to connect to
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Maximum number of connections to establish to the database. Zero means unlimited.
     */
    readonly maxConnections?: pulumi.Input<number>;
    /**
     * Password to be used if the PostgreSQL server demands password authentication
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
     */
    readonly port?: pulumi.Input<number>;
    /**
     * @deprecated Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     */
    readonly sslMode?: pulumi.Input<string>;
    /**
     * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     */
    readonly sslmode?: pulumi.Input<string>;
    /**
     * The SSL server root certificate file path. The file must contain PEM encoded data.
     */
    readonly sslrootcert?: pulumi.Input<string>;
    /**
     * Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
     * Refreshing state password from Postgres)
     */
    readonly superuser?: pulumi.Input<boolean>;
    /**
     * PostgreSQL user name to connect as
     */
    readonly username?: pulumi.Input<string>;
}
