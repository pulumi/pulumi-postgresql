[![Build Status](https://travis-ci.com/pulumi/pulumi-postgresql.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-postgresql)

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

    $ go get github.com/pulumi/pulumi-postgresql/sdk/go/...

## Configuration

The following configuration points are available:

- `postgresql:endpoint` (required) - The address of the postgresql server to use. Most often a "hostname:port" pair, but may also be an absolute path to a Unix socket when the host OS is Unix-compatible.
- `postgresql:username` (required) - Username to use to authenticate with the server.
- `postgresql:password` - Password for the given user, if that user has a password.
- `postgresql:tls` - The TLS configuration. One of false, true, or skip-verify. Defaults to false.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/postgresql/index.html).
