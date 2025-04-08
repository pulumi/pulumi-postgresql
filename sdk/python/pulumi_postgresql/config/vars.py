# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from .. import _utilities
from . import outputs

import types

__config__ = pulumi.Config('postgresql')


class _ExportableConfig(types.ModuleType):
    @property
    def aws_rds_iam_auth(self) -> Optional[bool]:
        """
        Use rds_iam instead of password authentication (see:
        https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
        """
        return __config__.get_bool('awsRdsIamAuth')

    @property
    def aws_rds_iam_profile(self) -> Optional[str]:
        """
        AWS profile to use for IAM auth
        """
        return __config__.get('awsRdsIamProfile')

    @property
    def aws_rds_iam_provider_role_arn(self) -> Optional[str]:
        """
        AWS IAM role to assume for IAM auth
        """
        return __config__.get('awsRdsIamProviderRoleArn')

    @property
    def aws_rds_iam_region(self) -> Optional[str]:
        """
        AWS region to use for IAM auth
        """
        return __config__.get('awsRdsIamRegion')

    @property
    def azure_identity_auth(self) -> Optional[bool]:
        """
        Use MS Azure identity OAuth token (see:
        https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication)
        """
        return __config__.get_bool('azureIdentityAuth')

    @property
    def azure_tenant_id(self) -> Optional[str]:
        return __config__.get('azureTenantId')

    @property
    def clientcert(self) -> Optional[str]:
        """
        SSL client certificate if required by the database.
        """
        return __config__.get('clientcert')

    @property
    def connect_timeout(self) -> int:
        """
        Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        """
        return __config__.get_int('connectTimeout') or (_utilities.get_env_int('PGCONNECT_TIMEOUT') or 180)

    @property
    def database(self) -> Optional[str]:
        """
        The name of the database to connect to in order to conenct to (defaults to `postgres`).
        """
        return __config__.get('database')

    @property
    def database_username(self) -> Optional[str]:
        """
        Database username associated to the connected user (for user name maps)
        """
        return __config__.get('databaseUsername')

    @property
    def expected_version(self) -> Optional[str]:
        """
        Specify the expected version of PostgreSQL.
        """
        return __config__.get('expectedVersion')

    @property
    def gcp_iam_impersonate_service_account(self) -> Optional[str]:
        """
        Service account to impersonate when using GCP IAM authentication.
        """
        return __config__.get('gcpIamImpersonateServiceAccount')

    @property
    def host(self) -> Optional[str]:
        """
        Name of PostgreSQL server address to connect to
        """
        return __config__.get('host')

    @property
    def max_connections(self) -> Optional[int]:
        """
        Maximum number of connections to establish to the database. Zero means unlimited.
        """
        return __config__.get_int('maxConnections')

    @property
    def password(self) -> Optional[str]:
        """
        Password to be used if the PostgreSQL server demands password authentication
        """
        return __config__.get('password')

    @property
    def port(self) -> Optional[int]:
        """
        The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
        """
        return __config__.get_int('port')

    @property
    def scheme(self) -> Optional[str]:
        return __config__.get('scheme')

    @property
    def ssl_mode(self) -> Optional[str]:
        return __config__.get('sslMode')

    @property
    def sslmode(self) -> Optional[str]:
        """
        This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
        PostgreSQL server
        """
        return __config__.get('sslmode') or _utilities.get_env('PGSSLMODE')

    @property
    def sslrootcert(self) -> Optional[str]:
        """
        The SSL server root certificate file path. The file must contain PEM encoded data.
        """
        return __config__.get('sslrootcert')

    @property
    def superuser(self) -> Optional[bool]:
        """
        Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
        Refreshing state password from Postgres)
        """
        return __config__.get_bool('superuser')

    @property
    def username(self) -> Optional[str]:
        """
        PostgreSQL user name to connect as
        """
        return __config__.get('username')

