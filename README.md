[![Actions Status](https://github.com/pulumi/pulumi-postgresql/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-postgresql/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fpostgresql.svg)](https://www.npmjs.com/package/@pulumi/postgresql)
[![Python version](https://badge.fury.io/py/pulumi-postgresql.svg)](https://pypi.org/project/pulumi-postgresql)
[![NuGet version](https://badge.fury.io/nu/pulumi.postgresql.svg)](https://badge.fury.io/nu/pulumi.postgresql)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-postgresql/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-postgresql/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-postgresql/blob/master/LICENSE)

# postgresql Resource Provider

The postgresql resource provider for Pulumi lets you manage postgresql resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/postgresql

or `yarn`:

    $ yarn add @pulumi/postgresql

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_postgresql

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-postgresql/sdk/v2/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Postgresql

## Configuration

The following configuration points are available:

- `postgresql:host` - (required) The address for the postgresql server connection. Can also be specified with the `PGHOST`
   environment variable.
- `postgresql:port` - (optional) The port for the postgresql server connection. The default is `5432`.  Can also be specified 
   with the `PGPORT` environment variable.
- `postgresql:database` - (optional) Database to connect to. The default is `postgres`. Can also be specified 
   with the `PGDATABASE` environment variable.
- `postgresql:username` - (required) Username for the server connection. The default is `postgres`. Can also be specified 
   with the `PGUSER` environment variable.
- `postgresql:password` - (optional) Password for the server connection. Can also be specified with the `PGPASSWORD` environment variable.
- `postgresql:database_username` - (optional) Username of the user in the database if different than connection username (See user name maps).
- `postgresql:superuser` - (optional) Should be set to false if the user to connect is not a PostgreSQL superuser (as is the case in RDS). 
   In this case, some features might be disabled (e.g.: Refreshing state password from database).
- `postgresql:sslmode` - (optional) Set the priority for an SSL connection to the server. Valid values for sslmode are (note: prefer is not supported by Go's lib/pq):
    * `disable` - No ssl
    * `require` - always SSL (the default, also skip verification)
    * `verify-ca` - always SSL (verify that the certificate presented by the server was signed by a trusted CA)
    * `verify-full` - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server 
       host name matches the one in the certificate) Additional information on the options and their implications can be seen in the libpq(3) SSL guide.
  Can also be specified with the `PGSSLMODE` environment variable. 
- `postgresql:connect_timeout` - (optional) Maximum wait for connection, in seconds. The default is `180s`. Zero or not specified means wait indefinitely. 
  Can also be specified with the `PGCONNECT_TIMEOUT` environment variable.
- `postgresql:max_connections` - (optional) Set the maximum number of open connections to the database. The default is `4`. Zero means unlimited open connections.
- `postgresql:expected_version` - (optional) Specify a hint to Terraform regarding the expected version that the provider will be talking with. This is a 
   required hint in order for the provider to talk with an ancient version of PostgreSQL. This parameter is expected to be a PostgreSQL Version or current. 
   Once a connection has been established, the provider will fingerprint the actual version. Default: 9.0.0.
- `postgresql:clientcert` - (optional) Clientcert block for configuring SSL certificate. 
  - `postgresql:clientcert.cert` - (required) The SSL client certificate file path. The file must contain PEM encoded data.
  - `postgresql:clientcert.key` - (required) The SSL client certificate private key file path. The file must contain PEM encoded data.


## Reference

For further information, please visit [the postgresql provider docs](https://www.pulumi.com/docs/intro/cloud-providers/postgresql) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/postgresql).
