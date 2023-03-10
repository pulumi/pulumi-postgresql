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
     * AWS profile to use for IAM auth
     */
    public readonly awsRdsIamProfile!: pulumi.Output<string | undefined>;
    /**
     * The name of the database to connect to in order to conenct to (defaults to `postgres`).
     */
    public readonly database!: pulumi.Output<string | undefined>;
    /**
     * Database username associated to the connected user (for user name maps)
     */
    public readonly databaseUsername!: pulumi.Output<string | undefined>;
    /**
     * Specify the expected version of PostgreSQL.
     */
    public readonly expectedVersion!: pulumi.Output<string | undefined>;
    /**
     * Name of PostgreSQL server address to connect to
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * Password to be used if the PostgreSQL server demands password authentication
     */
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly scheme!: pulumi.Output<string | undefined>;
    /**
     * @deprecated Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     */
    public readonly sslMode!: pulumi.Output<string | undefined>;
    /**
     * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     */
    public readonly sslmode!: pulumi.Output<string | undefined>;
    /**
     * The SSL server root certificate file path. The file must contain PEM encoded data.
     */
    public readonly sslrootcert!: pulumi.Output<string | undefined>;
    /**
     * PostgreSQL user name to connect as
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["awsRdsIamAuth"] = pulumi.output(args ? args.awsRdsIamAuth : undefined).apply(JSON.stringify);
            resourceInputs["awsRdsIamProfile"] = args ? args.awsRdsIamProfile : undefined;
            resourceInputs["clientcert"] = pulumi.output(args ? args.clientcert : undefined).apply(JSON.stringify);
            resourceInputs["connectTimeout"] = pulumi.output((args ? args.connectTimeout : undefined) ?? (utilities.getEnvNumber("PGCONNECT_TIMEOUT") || 180)).apply(JSON.stringify);
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["databaseUsername"] = args ? args.databaseUsername : undefined;
            resourceInputs["expectedVersion"] = args ? args.expectedVersion : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["maxConnections"] = pulumi.output(args ? args.maxConnections : undefined).apply(JSON.stringify);
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["port"] = pulumi.output(args ? args.port : undefined).apply(JSON.stringify);
            resourceInputs["scheme"] = args ? args.scheme : undefined;
            resourceInputs["sslMode"] = args ? args.sslMode : undefined;
            resourceInputs["sslmode"] = (args ? args.sslmode : undefined) ?? utilities.getEnv("PGSSLMODE");
            resourceInputs["sslrootcert"] = args ? args.sslrootcert : undefined;
            resourceInputs["superuser"] = pulumi.output(args ? args.superuser : undefined).apply(JSON.stringify);
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Use rds_iam instead of password authentication (see:
     * https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
     */
    awsRdsIamAuth?: pulumi.Input<boolean>;
    /**
     * AWS profile to use for IAM auth
     */
    awsRdsIamProfile?: pulumi.Input<string>;
    /**
     * SSL client certificate if required by the database.
     */
    clientcert?: pulumi.Input<inputs.ProviderClientcert>;
    /**
     * Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * The name of the database to connect to in order to conenct to (defaults to `postgres`).
     */
    database?: pulumi.Input<string>;
    /**
     * Database username associated to the connected user (for user name maps)
     */
    databaseUsername?: pulumi.Input<string>;
    /**
     * Specify the expected version of PostgreSQL.
     */
    expectedVersion?: pulumi.Input<string>;
    /**
     * Name of PostgreSQL server address to connect to
     */
    host?: pulumi.Input<string>;
    /**
     * Maximum number of connections to establish to the database. Zero means unlimited.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Password to be used if the PostgreSQL server demands password authentication
     */
    password?: pulumi.Input<string>;
    /**
     * The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
     */
    port?: pulumi.Input<number>;
    scheme?: pulumi.Input<string>;
    /**
     * @deprecated Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     */
    sslMode?: pulumi.Input<string>;
    /**
     * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     */
    sslmode?: pulumi.Input<string>;
    /**
     * The SSL server root certificate file path. The file must contain PEM encoded data.
     */
    sslrootcert?: pulumi.Input<string>;
    /**
     * Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
     * Refreshing state password from Postgres)
     */
    superuser?: pulumi.Input<boolean>;
    /**
     * PostgreSQL user name to connect as
     */
    username?: pulumi.Input<string>;
}
