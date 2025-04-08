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
from . import _utilities

__all__ = ['DefaultPrivilegArgs', 'DefaultPrivileg']

@pulumi.input_type
class DefaultPrivilegArgs:
    def __init__(__self__, *,
                 database: pulumi.Input[builtins.str],
                 object_type: pulumi.Input[builtins.str],
                 owner: pulumi.Input[builtins.str],
                 privileges: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 role: pulumi.Input[builtins.str],
                 schema: Optional[pulumi.Input[builtins.str]] = None,
                 with_grant_option: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a DefaultPrivileg resource.
        :param pulumi.Input[builtins.str] database: The database to grant default privileges for this role.
        :param pulumi.Input[builtins.str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        :param pulumi.Input[builtins.str] owner: Specifies the role that creates objects for which the default privileges will be applied.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] privileges: List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        :param pulumi.Input[builtins.str] role: The role that will automatically be granted the specified privileges on new objects created by the owner.
        :param pulumi.Input[builtins.str] schema: The database schema to set default privileges for this role.
        :param pulumi.Input[builtins.bool] with_grant_option: Permit the grant recipient to grant it to others
        """
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "object_type", object_type)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "privileges", privileges)
        pulumi.set(__self__, "role", role)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Input[builtins.str]:
        """
        The database to grant default privileges for this role.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> pulumi.Input[builtins.str]:
        """
        The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        """
        return pulumi.get(self, "object_type")

    @object_type.setter
    def object_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "object_type", value)

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the role that creates objects for which the default privileges will be applied.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[builtins.str]:
        """
        The role that will automatically be granted the specified privileges on new objects created by the owner.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The database schema to set default privileges for this role.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Permit the grant recipient to grant it to others
        """
        return pulumi.get(self, "with_grant_option")

    @with_grant_option.setter
    def with_grant_option(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "with_grant_option", value)


@pulumi.input_type
class _DefaultPrivilegState:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[builtins.str]] = None,
                 object_type: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[pulumi.Input[builtins.str]] = None,
                 with_grant_option: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering DefaultPrivileg resources.
        :param pulumi.Input[builtins.str] database: The database to grant default privileges for this role.
        :param pulumi.Input[builtins.str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        :param pulumi.Input[builtins.str] owner: Specifies the role that creates objects for which the default privileges will be applied.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] privileges: List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        :param pulumi.Input[builtins.str] role: The role that will automatically be granted the specified privileges on new objects created by the owner.
        :param pulumi.Input[builtins.str] schema: The database schema to set default privileges for this role.
        :param pulumi.Input[builtins.bool] with_grant_option: Permit the grant recipient to grant it to others
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if object_type is not None:
            pulumi.set(__self__, "object_type", object_type)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if privileges is not None:
            pulumi.set(__self__, "privileges", privileges)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The database to grant default privileges for this role.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        """
        return pulumi.get(self, "object_type")

    @object_type.setter
    def object_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "object_type", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the role that creates objects for which the default privileges will be applied.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role that will automatically be granted the specified privileges on new objects created by the owner.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The database schema to set default privileges for this role.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Permit the grant recipient to grant it to others
        """
        return pulumi.get(self, "with_grant_option")

    @with_grant_option.setter
    def with_grant_option(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "with_grant_option", value)


warnings.warn("""postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges""", DeprecationWarning)


class DefaultPrivileg(pulumi.CustomResource):
    warnings.warn("""postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[builtins.str]] = None,
                 object_type: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[pulumi.Input[builtins.str]] = None,
                 with_grant_option: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        The ``DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.

        > **Note:** This resource needs Postgresql version 9 or above.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        read_only_tables = postgresql.DefaultPrivileges("read_only_tables",
            role="test_role",
            database="test_db",
            schema="public",
            owner="db_owner",
            object_type="table",
            privileges=["SELECT"])
        ```

        ## Examples

        ### Grant default privileges for tables to "current_role" role:

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        grant_table_privileges = postgresql.DefaultPrivileges("grant_table_privileges",
            database=example_db["name"],
            role="current_role",
            owner="owner_role",
            schema="public",
            object_type="table",
            privileges=[
                "SELECT",
                "INSERT",
                "UPDATE",
            ])
        ```
        Whenever the `owner_role` creates a new table in the `public` schema, the `current_role` is automatically granted SELECT, INSERT, and UPDATE privileges on that table.

        ### Revoke default privileges for functions for "public" role:

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        revoke_public = postgresql.DefaultPrivileges("revoke_public",
            database=example_db["name"],
            role="public",
            owner="object_owner",
            object_type="function",
            privileges=[])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database: The database to grant default privileges for this role.
        :param pulumi.Input[builtins.str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        :param pulumi.Input[builtins.str] owner: Specifies the role that creates objects for which the default privileges will be applied.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] privileges: List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        :param pulumi.Input[builtins.str] role: The role that will automatically be granted the specified privileges on new objects created by the owner.
        :param pulumi.Input[builtins.str] schema: The database schema to set default privileges for this role.
        :param pulumi.Input[builtins.bool] with_grant_option: Permit the grant recipient to grant it to others
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultPrivilegArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.

        > **Note:** This resource needs Postgresql version 9 or above.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        read_only_tables = postgresql.DefaultPrivileges("read_only_tables",
            role="test_role",
            database="test_db",
            schema="public",
            owner="db_owner",
            object_type="table",
            privileges=["SELECT"])
        ```

        ## Examples

        ### Grant default privileges for tables to "current_role" role:

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        grant_table_privileges = postgresql.DefaultPrivileges("grant_table_privileges",
            database=example_db["name"],
            role="current_role",
            owner="owner_role",
            schema="public",
            object_type="table",
            privileges=[
                "SELECT",
                "INSERT",
                "UPDATE",
            ])
        ```
        Whenever the `owner_role` creates a new table in the `public` schema, the `current_role` is automatically granted SELECT, INSERT, and UPDATE privileges on that table.

        ### Revoke default privileges for functions for "public" role:

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        revoke_public = postgresql.DefaultPrivileges("revoke_public",
            database=example_db["name"],
            role="public",
            owner="object_owner",
            object_type="function",
            privileges=[])
        ```

        :param str resource_name: The name of the resource.
        :param DefaultPrivilegArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultPrivilegArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[builtins.str]] = None,
                 object_type: Optional[pulumi.Input[builtins.str]] = None,
                 owner: Optional[pulumi.Input[builtins.str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[pulumi.Input[builtins.str]] = None,
                 with_grant_option: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        pulumi.log.warn("""DefaultPrivileg is deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultPrivilegArgs.__new__(DefaultPrivilegArgs)

            if database is None and not opts.urn:
                raise TypeError("Missing required property 'database'")
            __props__.__dict__["database"] = database
            if object_type is None and not opts.urn:
                raise TypeError("Missing required property 'object_type'")
            __props__.__dict__["object_type"] = object_type
            if owner is None and not opts.urn:
                raise TypeError("Missing required property 'owner'")
            __props__.__dict__["owner"] = owner
            if privileges is None and not opts.urn:
                raise TypeError("Missing required property 'privileges'")
            __props__.__dict__["privileges"] = privileges
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["schema"] = schema
            __props__.__dict__["with_grant_option"] = with_grant_option
        super(DefaultPrivileg, __self__).__init__(
            'postgresql:index/defaultPrivileg:DefaultPrivileg',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[builtins.str]] = None,
            object_type: Optional[pulumi.Input[builtins.str]] = None,
            owner: Optional[pulumi.Input[builtins.str]] = None,
            privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            role: Optional[pulumi.Input[builtins.str]] = None,
            schema: Optional[pulumi.Input[builtins.str]] = None,
            with_grant_option: Optional[pulumi.Input[builtins.bool]] = None) -> 'DefaultPrivileg':
        """
        Get an existing DefaultPrivileg resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database: The database to grant default privileges for this role.
        :param pulumi.Input[builtins.str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        :param pulumi.Input[builtins.str] owner: Specifies the role that creates objects for which the default privileges will be applied.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] privileges: List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        :param pulumi.Input[builtins.str] role: The role that will automatically be granted the specified privileges on new objects created by the owner.
        :param pulumi.Input[builtins.str] schema: The database schema to set default privileges for this role.
        :param pulumi.Input[builtins.bool] with_grant_option: Permit the grant recipient to grant it to others
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultPrivilegState.__new__(_DefaultPrivilegState)

        __props__.__dict__["database"] = database
        __props__.__dict__["object_type"] = object_type
        __props__.__dict__["owner"] = owner
        __props__.__dict__["privileges"] = privileges
        __props__.__dict__["role"] = role
        __props__.__dict__["schema"] = schema
        __props__.__dict__["with_grant_option"] = with_grant_option
        return DefaultPrivileg(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[builtins.str]:
        """
        The database to grant default privileges for this role.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> pulumi.Output[builtins.str]:
        """
        The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
        """
        return pulumi.get(self, "object_type")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the role that creates objects for which the default privileges will be applied.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of privileges (e.g., SELECT, INSERT, UPDATE, DELETE) to grant on new objects created by the owner. An empty list could be provided to revoke all default privileges for this role.
        """
        return pulumi.get(self, "privileges")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        The role that will automatically be granted the specified privileges on new objects created by the owner.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The database schema to set default privileges for this role.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Permit the grant recipient to grant it to others
        """
        return pulumi.get(self, "with_grant_option")

