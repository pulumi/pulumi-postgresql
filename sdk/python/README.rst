|Build Status|

postgresql Resource Provider
============================

The postgresql resource provider for Pulumi lets you manage postgresql
resources in your cloud programs. To use this package, please `install
the Pulumi CLI first <https://pulumi.io/>`__.

Installing
----------

This package is available in many languages in the standard packaging
formats.

Node.js (Java/TypeScript)
~~~~~~~~~~~~~~~~~~~~~~~~~

To use from JavaScript or TypeScript in Node.js, install using either
``npm``:

::

   $ npm install @pulumi/postgresql

or ``yarn``:

::

   $ yarn add @pulumi/postgresql

Python
~~~~~~

To use from Python, install using ``pip``:

::

   $ pip install pulumi_postgresql

Go
~~

To use from Go, use ``go get`` to grab the latest version of the library

::

   $ go get github.com/pulumi/pulumi-postgresql/sdk/go/...

Configuration
-------------

The following configuration points are available:

-  ``postgresql:host`` - (required) The address for the postgresql
   server connection.
-  ``postgresql:port`` - (optional) The port for the postgresql server
   connection. The default is 5432.
-  ``postgresql:database`` - (optional) Database to connect to. The
   default is postgres.
-  ``postgresql:username`` - (required) Username for the server
   connection.
-  ``postgresql:password`` - (optional) Password for the server
   connection.
-  ``postgresql:database_username`` - (optional) Username of the user in
   the database if different than connection username (See user name
   maps).
-  ``postgresql:superuser`` - (optional) Should be set to false if the
   user to connect is not a PostgreSQL superuser (as is the case in
   RDS). In this case, some features might be disabled (e.g.: Refreshing
   state password from database).
-  ``postgresql:sslmode`` - (optional) Set the priority for an SSL
   connection to the server. Valid values for sslmode are (note: prefer
   is not supported by Go’s lib/pq):

   -  ``disable`` - No ssl
   -  ``require`` - always SSL (the default, also skip verification)
   -  ``verify-ca`` - always SSL (verify that the certificate presented
      by the server was signed by a trusted CA)
   -  ``verify-full`` - Always SSL (verify that the certification
      presented by the server was signed by a trusted CA and the server
      host name matches the one in the certificate) Additional
      information on the options and their implications can be seen in
      the libpq(3) SSL guide.

-  ``postgresql:connect_timeout`` - (optional) Maximum wait for
   connection, in seconds. The default is 180s. Zero or not specified
   means wait indefinitely.
-  ``postgresql:max_connections`` - (optional) Set the maximum number of
   open connections to the database. The default is 4. Zero means
   unlimited open connections.
-  ``postgresql:expected_version`` - (optional) Specify a hint to
   Terraform regarding the expected version that the provider will be
   talking with. This is a required hint in order for Terraform to talk
   with an ancient version of PostgreSQL. This parameter is expected to
   be a PostgreSQL Version or current. Once a connection has been
   established, Terraform will fingerprint the actual version. Default:
   9.0.0.

Reference
---------

For detailed reference documentation, please visit `the API
docs <https://pulumi.io/reference/pkg/nodejs/@pulumi/postgresql/index.html>`__.

.. |Build Status| image:: https://travis-ci.com/pulumi/pulumi-postgresql.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master
   :target: https://travis-ci.com/pulumi/pulumi-postgresql
