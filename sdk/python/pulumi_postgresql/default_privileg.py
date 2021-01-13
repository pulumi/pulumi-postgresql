# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DefaultPrivileg']

warnings.warn("""postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges""", DeprecationWarning)


class DefaultPrivileg(pulumi.CustomResource):
    warnings.warn("""postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 object_type: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a DefaultPrivileg resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant default privileges for this role
        :param pulumi.Input[str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
        :param pulumi.Input[str] owner: Target role for which to alter default privileges.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The list of privileges to apply as default privileges
        :param pulumi.Input[str] role: The name of the role to which grant default privileges on
        :param pulumi.Input[str] schema: The database schema to set default privileges for this role
        """
        pulumi.log.warn("DefaultPrivileg is deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if database is None and not opts.urn:
                raise TypeError("Missing required property 'database'")
            __props__['database'] = database
            if object_type is None and not opts.urn:
                raise TypeError("Missing required property 'object_type'")
            __props__['object_type'] = object_type
            if owner is None and not opts.urn:
                raise TypeError("Missing required property 'owner'")
            __props__['owner'] = owner
            if privileges is None and not opts.urn:
                raise TypeError("Missing required property 'privileges'")
            __props__['privileges'] = privileges
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if schema is None and not opts.urn:
                raise TypeError("Missing required property 'schema'")
            __props__['schema'] = schema
        super(DefaultPrivileg, __self__).__init__(
            'postgresql:index/defaultPrivileg:DefaultPrivileg',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            object_type: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None,
            schema: Optional[pulumi.Input[str]] = None) -> 'DefaultPrivileg':
        """
        Get an existing DefaultPrivileg resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant default privileges for this role
        :param pulumi.Input[str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
        :param pulumi.Input[str] owner: Target role for which to alter default privileges.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The list of privileges to apply as default privileges
        :param pulumi.Input[str] role: The name of the role to which grant default privileges on
        :param pulumi.Input[str] schema: The database schema to set default privileges for this role
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database"] = database
        __props__["object_type"] = object_type
        __props__["owner"] = owner
        __props__["privileges"] = privileges
        __props__["role"] = role
        __props__["schema"] = schema
        return DefaultPrivileg(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        The database to grant default privileges for this role
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> pulumi.Output[str]:
        """
        The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
        """
        return pulumi.get(self, "object_type")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Target role for which to alter default privileges.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of privileges to apply as default privileges
        """
        return pulumi.get(self, "privileges")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name of the role to which grant default privileges on
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        """
        The database schema to set default privileges for this role
        """
        return pulumi.get(self, "schema")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

