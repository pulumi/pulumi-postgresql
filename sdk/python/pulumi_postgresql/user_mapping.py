# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserMappingArgs', 'UserMapping']

@pulumi.input_type
class UserMappingArgs:
    def __init__(__self__, *,
                 server_name: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a UserMapping resource.
        :param pulumi.Input[str] server_name: The name of an existing server for which the user mapping is to be created.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        :param pulumi.Input[str] user_name: The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        """
        pulumi.set(__self__, "server_name", server_name)
        pulumi.set(__self__, "user_name", user_name)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Input[str]:
        """
        The name of an existing server for which the user mapping is to be created.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)


@pulumi.input_type
class _UserMappingState:
    def __init__(__self__, *,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserMapping resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        :param pulumi.Input[str] server_name: The name of an existing server for which the user mapping is to be created.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        :param pulumi.Input[str] user_name: The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        """
        if options is not None:
            pulumi.set(__self__, "options", options)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an existing server for which the user mapping is to be created.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class UserMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The ``UserMapping`` resource creates and manages a user mapping on a PostgreSQL server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        ext_postgres_fdw = postgresql.Extension("extPostgresFdw")
        myserver_postgres = postgresql.Server("myserverPostgres",
            server_name="myserver_postgres",
            fdw_name="postgres_fdw",
            options={
                "host": "foo",
                "dbname": "foodb",
                "port": "5432",
            },
            opts=pulumi.ResourceOptions(depends_on=[ext_postgres_fdw]))
        remote_role = postgresql.Role("remoteRole")
        remote_user_mapping = postgresql.UserMapping("remoteUserMapping",
            server_name=myserver_postgres.server_name,
            user_name=remote_role.name,
            options={
                "user": "admin",
                "password": "pass",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        :param pulumi.Input[str] server_name: The name of an existing server for which the user mapping is to be created.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        :param pulumi.Input[str] user_name: The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``UserMapping`` resource creates and manages a user mapping on a PostgreSQL server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        ext_postgres_fdw = postgresql.Extension("extPostgresFdw")
        myserver_postgres = postgresql.Server("myserverPostgres",
            server_name="myserver_postgres",
            fdw_name="postgres_fdw",
            options={
                "host": "foo",
                "dbname": "foodb",
                "port": "5432",
            },
            opts=pulumi.ResourceOptions(depends_on=[ext_postgres_fdw]))
        remote_role = postgresql.Role("remoteRole")
        remote_user_mapping = postgresql.UserMapping("remoteUserMapping",
            server_name=myserver_postgres.server_name,
            user_name=remote_role.name,
            options={
                "user": "admin",
                "password": "pass",
            })
        ```

        :param str resource_name: The name of the resource.
        :param UserMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserMappingArgs.__new__(UserMappingArgs)

            __props__.__dict__["options"] = options
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__.__dict__["server_name"] = server_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(UserMapping, __self__).__init__(
            'postgresql:index/userMapping:UserMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'UserMapping':
        """
        Get an existing UserMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        :param pulumi.Input[str] server_name: The name of an existing server for which the user mapping is to be created.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        :param pulumi.Input[str] user_name: The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
               Changing this value
               will force the creation of a new resource as this value can only be set
               when the user mapping is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserMappingState.__new__(_UserMappingState)

        __props__.__dict__["options"] = options
        __props__.__dict__["server_name"] = server_name
        __props__.__dict__["user_name"] = user_name
        return UserMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        The name of an existing server for which the user mapping is to be created.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
        Changing this value
        will force the creation of a new resource as this value can only be set
        when the user mapping is created.
        """
        return pulumi.get(self, "user_name")

