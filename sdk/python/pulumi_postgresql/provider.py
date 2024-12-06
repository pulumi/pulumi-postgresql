# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from ._inputs import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 aws_rds_iam_auth: Optional[pulumi.Input[bool]] = None,
                 aws_rds_iam_profile: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_provider_role_arn: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_region: Optional[pulumi.Input[str]] = None,
                 azure_identity_auth: Optional[pulumi.Input[bool]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input['ProviderClientcertArgs']] = None,
                 connect_timeout: Optional[pulumi.Input[int]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 database_username: Optional[pulumi.Input[str]] = None,
                 expected_version: Optional[pulumi.Input[str]] = None,
                 gcp_iam_impersonate_service_account: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 sslmode: Optional[pulumi.Input[str]] = None,
                 sslrootcert: Optional[pulumi.Input[str]] = None,
                 superuser: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[bool] aws_rds_iam_auth: Use rds_iam instead of password authentication (see:
               https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
        :param pulumi.Input[str] aws_rds_iam_profile: AWS profile to use for IAM auth
        :param pulumi.Input[str] aws_rds_iam_provider_role_arn: AWS IAM role to assume for IAM auth
        :param pulumi.Input[str] aws_rds_iam_region: AWS region to use for IAM auth
        :param pulumi.Input[bool] azure_identity_auth: Use MS Azure identity OAuth token (see:
               https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
        :param pulumi.Input['ProviderClientcertArgs'] clientcert: SSL client certificate if required by the database.
        :param pulumi.Input[int] connect_timeout: Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        :param pulumi.Input[str] database: The name of the database to connect to in order to conenct to (defaults to `postgres`).
        :param pulumi.Input[str] database_username: Database username associated to the connected user (for user name maps)
        :param pulumi.Input[str] expected_version: Specify the expected version of PostgreSQL.
        :param pulumi.Input[str] gcp_iam_impersonate_service_account: Service account to impersonate when using GCP IAM authentication.
        :param pulumi.Input[str] host: Name of PostgreSQL server address to connect to
        :param pulumi.Input[int] max_connections: Maximum number of connections to establish to the database. Zero means unlimited.
        :param pulumi.Input[str] password: Password to be used if the PostgreSQL server demands password authentication
        :param pulumi.Input[int] port: The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
        :param pulumi.Input[str] sslmode: This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
               PostgreSQL server
        :param pulumi.Input[str] sslrootcert: The SSL server root certificate file path. The file must contain PEM encoded data.
        :param pulumi.Input[bool] superuser: Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
               Refreshing state password from Postgres)
        :param pulumi.Input[str] username: PostgreSQL user name to connect as
        """
        if aws_rds_iam_auth is not None:
            pulumi.set(__self__, "aws_rds_iam_auth", aws_rds_iam_auth)
        if aws_rds_iam_profile is not None:
            pulumi.set(__self__, "aws_rds_iam_profile", aws_rds_iam_profile)
        if aws_rds_iam_provider_role_arn is not None:
            pulumi.set(__self__, "aws_rds_iam_provider_role_arn", aws_rds_iam_provider_role_arn)
        if aws_rds_iam_region is not None:
            pulumi.set(__self__, "aws_rds_iam_region", aws_rds_iam_region)
        if azure_identity_auth is not None:
            pulumi.set(__self__, "azure_identity_auth", azure_identity_auth)
        if azure_tenant_id is not None:
            pulumi.set(__self__, "azure_tenant_id", azure_tenant_id)
        if clientcert is not None:
            pulumi.set(__self__, "clientcert", clientcert)
        if connect_timeout is None:
            connect_timeout = (_utilities.get_env_int('PGCONNECT_TIMEOUT') or 180)
        if connect_timeout is not None:
            pulumi.set(__self__, "connect_timeout", connect_timeout)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if database_username is not None:
            pulumi.set(__self__, "database_username", database_username)
        if expected_version is not None:
            pulumi.set(__self__, "expected_version", expected_version)
        if gcp_iam_impersonate_service_account is not None:
            pulumi.set(__self__, "gcp_iam_impersonate_service_account", gcp_iam_impersonate_service_account)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if scheme is not None:
            pulumi.set(__self__, "scheme", scheme)
        if ssl_mode is not None:
            warnings.warn("""Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`""", DeprecationWarning)
            pulumi.log.warn("""ssl_mode is deprecated: Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`""")
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)
        if sslmode is None:
            sslmode = _utilities.get_env('PGSSLMODE')
        if sslmode is not None:
            pulumi.set(__self__, "sslmode", sslmode)
        if sslrootcert is not None:
            pulumi.set(__self__, "sslrootcert", sslrootcert)
        if superuser is not None:
            pulumi.set(__self__, "superuser", superuser)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="awsRdsIamAuth")
    def aws_rds_iam_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        Use rds_iam instead of password authentication (see:
        https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
        """
        return pulumi.get(self, "aws_rds_iam_auth")

    @aws_rds_iam_auth.setter
    def aws_rds_iam_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "aws_rds_iam_auth", value)

    @property
    @pulumi.getter(name="awsRdsIamProfile")
    def aws_rds_iam_profile(self) -> Optional[pulumi.Input[str]]:
        """
        AWS profile to use for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_profile")

    @aws_rds_iam_profile.setter
    def aws_rds_iam_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_rds_iam_profile", value)

    @property
    @pulumi.getter(name="awsRdsIamProviderRoleArn")
    def aws_rds_iam_provider_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        AWS IAM role to assume for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_provider_role_arn")

    @aws_rds_iam_provider_role_arn.setter
    def aws_rds_iam_provider_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_rds_iam_provider_role_arn", value)

    @property
    @pulumi.getter(name="awsRdsIamRegion")
    def aws_rds_iam_region(self) -> Optional[pulumi.Input[str]]:
        """
        AWS region to use for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_region")

    @aws_rds_iam_region.setter
    def aws_rds_iam_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_rds_iam_region", value)

    @property
    @pulumi.getter(name="azureIdentityAuth")
    def azure_identity_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        Use MS Azure identity OAuth token (see:
        https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
        """
        return pulumi.get(self, "azure_identity_auth")

    @azure_identity_auth.setter
    def azure_identity_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "azure_identity_auth", value)

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "azure_tenant_id")

    @azure_tenant_id.setter
    def azure_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_tenant_id", value)

    @property
    @pulumi.getter
    def clientcert(self) -> Optional[pulumi.Input['ProviderClientcertArgs']]:
        """
        SSL client certificate if required by the database.
        """
        return pulumi.get(self, "clientcert")

    @clientcert.setter
    def clientcert(self, value: Optional[pulumi.Input['ProviderClientcertArgs']]):
        pulumi.set(self, "clientcert", value)

    @property
    @pulumi.getter(name="connectTimeout")
    def connect_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        """
        return pulumi.get(self, "connect_timeout")

    @connect_timeout.setter
    def connect_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connect_timeout", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database to connect to in order to conenct to (defaults to `postgres`).
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="databaseUsername")
    def database_username(self) -> Optional[pulumi.Input[str]]:
        """
        Database username associated to the connected user (for user name maps)
        """
        return pulumi.get(self, "database_username")

    @database_username.setter
    def database_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_username", value)

    @property
    @pulumi.getter(name="expectedVersion")
    def expected_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the expected version of PostgreSQL.
        """
        return pulumi.get(self, "expected_version")

    @expected_version.setter
    def expected_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_version", value)

    @property
    @pulumi.getter(name="gcpIamImpersonateServiceAccount")
    def gcp_iam_impersonate_service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Service account to impersonate when using GCP IAM authentication.
        """
        return pulumi.get(self, "gcp_iam_impersonate_service_account")

    @gcp_iam_impersonate_service_account.setter
    def gcp_iam_impersonate_service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcp_iam_impersonate_service_account", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Name of PostgreSQL server address to connect to
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of connections to establish to the database. Zero means unlimited.
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password to be used if the PostgreSQL server demands password authentication
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def scheme(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scheme")

    @scheme.setter
    def scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scheme", value)

    @property
    @pulumi.getter(name="sslMode")
    @_utilities.deprecated("""Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`""")
    def ssl_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_mode", value)

    @property
    @pulumi.getter
    def sslmode(self) -> Optional[pulumi.Input[str]]:
        """
        This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
        PostgreSQL server
        """
        return pulumi.get(self, "sslmode")

    @sslmode.setter
    def sslmode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sslmode", value)

    @property
    @pulumi.getter
    def sslrootcert(self) -> Optional[pulumi.Input[str]]:
        """
        The SSL server root certificate file path. The file must contain PEM encoded data.
        """
        return pulumi.get(self, "sslrootcert")

    @sslrootcert.setter
    def sslrootcert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sslrootcert", value)

    @property
    @pulumi.getter
    def superuser(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
        Refreshing state password from Postgres)
        """
        return pulumi.get(self, "superuser")

    @superuser.setter
    def superuser(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "superuser", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        PostgreSQL user name to connect as
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_rds_iam_auth: Optional[pulumi.Input[bool]] = None,
                 aws_rds_iam_profile: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_provider_role_arn: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_region: Optional[pulumi.Input[str]] = None,
                 azure_identity_auth: Optional[pulumi.Input[bool]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input[Union['ProviderClientcertArgs', 'ProviderClientcertArgsDict']]] = None,
                 connect_timeout: Optional[pulumi.Input[int]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 database_username: Optional[pulumi.Input[str]] = None,
                 expected_version: Optional[pulumi.Input[str]] = None,
                 gcp_iam_impersonate_service_account: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 sslmode: Optional[pulumi.Input[str]] = None,
                 sslrootcert: Optional[pulumi.Input[str]] = None,
                 superuser: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the postgresql package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] aws_rds_iam_auth: Use rds_iam instead of password authentication (see:
               https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
        :param pulumi.Input[str] aws_rds_iam_profile: AWS profile to use for IAM auth
        :param pulumi.Input[str] aws_rds_iam_provider_role_arn: AWS IAM role to assume for IAM auth
        :param pulumi.Input[str] aws_rds_iam_region: AWS region to use for IAM auth
        :param pulumi.Input[bool] azure_identity_auth: Use MS Azure identity OAuth token (see:
               https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
        :param pulumi.Input[Union['ProviderClientcertArgs', 'ProviderClientcertArgsDict']] clientcert: SSL client certificate if required by the database.
        :param pulumi.Input[int] connect_timeout: Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        :param pulumi.Input[str] database: The name of the database to connect to in order to conenct to (defaults to `postgres`).
        :param pulumi.Input[str] database_username: Database username associated to the connected user (for user name maps)
        :param pulumi.Input[str] expected_version: Specify the expected version of PostgreSQL.
        :param pulumi.Input[str] gcp_iam_impersonate_service_account: Service account to impersonate when using GCP IAM authentication.
        :param pulumi.Input[str] host: Name of PostgreSQL server address to connect to
        :param pulumi.Input[int] max_connections: Maximum number of connections to establish to the database. Zero means unlimited.
        :param pulumi.Input[str] password: Password to be used if the PostgreSQL server demands password authentication
        :param pulumi.Input[int] port: The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
        :param pulumi.Input[str] sslmode: This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
               PostgreSQL server
        :param pulumi.Input[str] sslrootcert: The SSL server root certificate file path. The file must contain PEM encoded data.
        :param pulumi.Input[bool] superuser: Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
               Refreshing state password from Postgres)
        :param pulumi.Input[str] username: PostgreSQL user name to connect as
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the postgresql package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_rds_iam_auth: Optional[pulumi.Input[bool]] = None,
                 aws_rds_iam_profile: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_provider_role_arn: Optional[pulumi.Input[str]] = None,
                 aws_rds_iam_region: Optional[pulumi.Input[str]] = None,
                 azure_identity_auth: Optional[pulumi.Input[bool]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input[Union['ProviderClientcertArgs', 'ProviderClientcertArgsDict']]] = None,
                 connect_timeout: Optional[pulumi.Input[int]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 database_username: Optional[pulumi.Input[str]] = None,
                 expected_version: Optional[pulumi.Input[str]] = None,
                 gcp_iam_impersonate_service_account: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 sslmode: Optional[pulumi.Input[str]] = None,
                 sslrootcert: Optional[pulumi.Input[str]] = None,
                 superuser: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["aws_rds_iam_auth"] = pulumi.Output.from_input(aws_rds_iam_auth).apply(pulumi.runtime.to_json) if aws_rds_iam_auth is not None else None
            __props__.__dict__["aws_rds_iam_profile"] = aws_rds_iam_profile
            __props__.__dict__["aws_rds_iam_provider_role_arn"] = aws_rds_iam_provider_role_arn
            __props__.__dict__["aws_rds_iam_region"] = aws_rds_iam_region
            __props__.__dict__["azure_identity_auth"] = pulumi.Output.from_input(azure_identity_auth).apply(pulumi.runtime.to_json) if azure_identity_auth is not None else None
            __props__.__dict__["azure_tenant_id"] = azure_tenant_id
            __props__.__dict__["clientcert"] = pulumi.Output.from_input(clientcert).apply(pulumi.runtime.to_json) if clientcert is not None else None
            if connect_timeout is None:
                connect_timeout = (_utilities.get_env_int('PGCONNECT_TIMEOUT') or 180)
            __props__.__dict__["connect_timeout"] = pulumi.Output.from_input(connect_timeout).apply(pulumi.runtime.to_json) if connect_timeout is not None else None
            __props__.__dict__["database"] = database
            __props__.__dict__["database_username"] = database_username
            __props__.__dict__["expected_version"] = expected_version
            __props__.__dict__["gcp_iam_impersonate_service_account"] = gcp_iam_impersonate_service_account
            __props__.__dict__["host"] = host
            __props__.__dict__["max_connections"] = pulumi.Output.from_input(max_connections).apply(pulumi.runtime.to_json) if max_connections is not None else None
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = pulumi.Output.from_input(port).apply(pulumi.runtime.to_json) if port is not None else None
            __props__.__dict__["scheme"] = scheme
            __props__.__dict__["ssl_mode"] = ssl_mode
            if sslmode is None:
                sslmode = _utilities.get_env('PGSSLMODE')
            __props__.__dict__["sslmode"] = sslmode
            __props__.__dict__["sslrootcert"] = sslrootcert
            __props__.__dict__["superuser"] = pulumi.Output.from_input(superuser).apply(pulumi.runtime.to_json) if superuser is not None else None
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'postgresql',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="awsRdsIamProfile")
    def aws_rds_iam_profile(self) -> pulumi.Output[Optional[str]]:
        """
        AWS profile to use for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_profile")

    @property
    @pulumi.getter(name="awsRdsIamProviderRoleArn")
    def aws_rds_iam_provider_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        AWS IAM role to assume for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_provider_role_arn")

    @property
    @pulumi.getter(name="awsRdsIamRegion")
    def aws_rds_iam_region(self) -> pulumi.Output[Optional[str]]:
        """
        AWS region to use for IAM auth
        """
        return pulumi.get(self, "aws_rds_iam_region")

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "azure_tenant_id")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the database to connect to in order to conenct to (defaults to `postgres`).
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="databaseUsername")
    def database_username(self) -> pulumi.Output[Optional[str]]:
        """
        Database username associated to the connected user (for user name maps)
        """
        return pulumi.get(self, "database_username")

    @property
    @pulumi.getter(name="expectedVersion")
    def expected_version(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the expected version of PostgreSQL.
        """
        return pulumi.get(self, "expected_version")

    @property
    @pulumi.getter(name="gcpIamImpersonateServiceAccount")
    def gcp_iam_impersonate_service_account(self) -> pulumi.Output[Optional[str]]:
        """
        Service account to impersonate when using GCP IAM authentication.
        """
        return pulumi.get(self, "gcp_iam_impersonate_service_account")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        Name of PostgreSQL server address to connect to
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password to be used if the PostgreSQL server demands password authentication
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def scheme(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "scheme")

    @property
    @pulumi.getter(name="sslMode")
    @_utilities.deprecated("""Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`""")
    def ssl_mode(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ssl_mode")

    @property
    @pulumi.getter
    def sslmode(self) -> pulumi.Output[Optional[str]]:
        """
        This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
        PostgreSQL server
        """
        return pulumi.get(self, "sslmode")

    @property
    @pulumi.getter
    def sslrootcert(self) -> pulumi.Output[Optional[str]]:
        """
        The SSL server root certificate file path. The file must contain PEM encoded data.
        """
        return pulumi.get(self, "sslrootcert")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        PostgreSQL user name to connect as
        """
        return pulumi.get(self, "username")

