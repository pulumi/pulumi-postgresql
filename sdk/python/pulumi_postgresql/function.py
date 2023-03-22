# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FunctionArgs', 'Function']

@pulumi.input_type
class FunctionArgs:
    def __init__(__self__, *,
                 body: pulumi.Input[str],
                 args: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 returns: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Function resource.
        :param pulumi.Input[str] body: Function body.
               This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        :param pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]] args: List of arguments for the function.
        :param pulumi.Input[str] database: The database where the function is located.
               If not specified, the function is created in the current database.
        :param pulumi.Input[bool] drop_cascade: True to automatically drop objects that depend on the function (such as 
               operators or triggers), and in turn all objects that depend on those objects. Default is false.
        :param pulumi.Input[str] language: The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        :param pulumi.Input[str] name: The name of the argument.
        :param pulumi.Input[str] returns: Type that the function returns. It can be computed from the OUT arguments. Default is void.
        :param pulumi.Input[str] schema: The schema where the function is located.
               If not specified, the function is created in the current schema.
        """
        pulumi.set(__self__, "body", body)
        if args is not None:
            pulumi.set(__self__, "args", args)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if returns is not None:
            pulumi.set(__self__, "returns", returns)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter
    def body(self) -> pulumi.Input[str]:
        """
        Function body.
        This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: pulumi.Input[str]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]]:
        """
        List of arguments for the function.
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The database where the function is located.
        If not specified, the function is created in the current database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        True to automatically drop objects that depend on the function (such as 
        operators or triggers), and in turn all objects that depend on those objects. Default is false.
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[str]]:
        """
        The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def returns(self) -> Optional[pulumi.Input[str]]:
        """
        Type that the function returns. It can be computed from the OUT arguments. Default is void.
        """
        return pulumi.get(self, "returns")

    @returns.setter
    def returns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "returns", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        The schema where the function is located.
        If not specified, the function is created in the current schema.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)


@pulumi.input_type
class _FunctionState:
    def __init__(__self__, *,
                 args: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 returns: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Function resources.
        :param pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]] args: List of arguments for the function.
        :param pulumi.Input[str] body: Function body.
               This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        :param pulumi.Input[str] database: The database where the function is located.
               If not specified, the function is created in the current database.
        :param pulumi.Input[bool] drop_cascade: True to automatically drop objects that depend on the function (such as 
               operators or triggers), and in turn all objects that depend on those objects. Default is false.
        :param pulumi.Input[str] language: The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        :param pulumi.Input[str] name: The name of the argument.
        :param pulumi.Input[str] returns: Type that the function returns. It can be computed from the OUT arguments. Default is void.
        :param pulumi.Input[str] schema: The schema where the function is located.
               If not specified, the function is created in the current schema.
        """
        if args is not None:
            pulumi.set(__self__, "args", args)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if returns is not None:
            pulumi.set(__self__, "returns", returns)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]]:
        """
        List of arguments for the function.
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionArgArgs']]]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        Function body.
        This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The database where the function is located.
        If not specified, the function is created in the current database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        True to automatically drop objects that depend on the function (such as 
        operators or triggers), and in turn all objects that depend on those objects. Default is false.
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[str]]:
        """
        The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def returns(self) -> Optional[pulumi.Input[str]]:
        """
        Type that the function returns. It can be computed from the OUT arguments. Default is void.
        """
        return pulumi.get(self, "returns")

    @returns.setter
    def returns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "returns", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        The schema where the function is located.
        If not specified, the function is created in the current schema.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)


class Function(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionArgArgs']]]]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 returns: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The ``Function`` resource creates and manages a function on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        increment = postgresql.Function("increment",
            args=[postgresql.FunctionArgArgs(
                name="i",
                type="integer",
            )],
            body=\"\"\"    BEGIN
                RETURN i + 1;
            END;

        \"\"\",
            language="plpgsql",
            returns="integer")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionArgArgs']]]] args: List of arguments for the function.
        :param pulumi.Input[str] body: Function body.
               This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        :param pulumi.Input[str] database: The database where the function is located.
               If not specified, the function is created in the current database.
        :param pulumi.Input[bool] drop_cascade: True to automatically drop objects that depend on the function (such as 
               operators or triggers), and in turn all objects that depend on those objects. Default is false.
        :param pulumi.Input[str] language: The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        :param pulumi.Input[str] name: The name of the argument.
        :param pulumi.Input[str] returns: Type that the function returns. It can be computed from the OUT arguments. Default is void.
        :param pulumi.Input[str] schema: The schema where the function is located.
               If not specified, the function is created in the current schema.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``Function`` resource creates and manages a function on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        increment = postgresql.Function("increment",
            args=[postgresql.FunctionArgArgs(
                name="i",
                type="integer",
            )],
            body=\"\"\"    BEGIN
                RETURN i + 1;
            END;

        \"\"\",
            language="plpgsql",
            returns="integer")
        ```

        :param str resource_name: The name of the resource.
        :param FunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionArgArgs']]]]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 returns: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionArgs.__new__(FunctionArgs)

            __props__.__dict__["args"] = args
            if body is None and not opts.urn:
                raise TypeError("Missing required property 'body'")
            __props__.__dict__["body"] = body
            __props__.__dict__["database"] = database
            __props__.__dict__["drop_cascade"] = drop_cascade
            __props__.__dict__["language"] = language
            __props__.__dict__["name"] = name
            __props__.__dict__["returns"] = returns
            __props__.__dict__["schema"] = schema
        super(Function, __self__).__init__(
            'postgresql:index/function:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            args: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionArgArgs']]]]] = None,
            body: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            drop_cascade: Optional[pulumi.Input[bool]] = None,
            language: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            returns: Optional[pulumi.Input[str]] = None,
            schema: Optional[pulumi.Input[str]] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionArgArgs']]]] args: List of arguments for the function.
        :param pulumi.Input[str] body: Function body.
               This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        :param pulumi.Input[str] database: The database where the function is located.
               If not specified, the function is created in the current database.
        :param pulumi.Input[bool] drop_cascade: True to automatically drop objects that depend on the function (such as 
               operators or triggers), and in turn all objects that depend on those objects. Default is false.
        :param pulumi.Input[str] language: The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        :param pulumi.Input[str] name: The name of the argument.
        :param pulumi.Input[str] returns: Type that the function returns. It can be computed from the OUT arguments. Default is void.
        :param pulumi.Input[str] schema: The schema where the function is located.
               If not specified, the function is created in the current schema.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionState.__new__(_FunctionState)

        __props__.__dict__["args"] = args
        __props__.__dict__["body"] = body
        __props__.__dict__["database"] = database
        __props__.__dict__["drop_cascade"] = drop_cascade
        __props__.__dict__["language"] = language
        __props__.__dict__["name"] = name
        __props__.__dict__["returns"] = returns
        __props__.__dict__["schema"] = schema
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output[Optional[Sequence['outputs.FunctionArg']]]:
        """
        List of arguments for the function.
        """
        return pulumi.get(self, "args")

    @property
    @pulumi.getter
    def body(self) -> pulumi.Output[str]:
        """
        Function body.
        This should be the body content withing the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        The database where the function is located.
        If not specified, the function is created in the current database.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> pulumi.Output[Optional[bool]]:
        """
        True to automatically drop objects that depend on the function (such as 
        operators or triggers), and in turn all objects that depend on those objects. Default is false.
        """
        return pulumi.get(self, "drop_cascade")

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[Optional[str]]:
        """
        The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def returns(self) -> pulumi.Output[str]:
        """
        Type that the function returns. It can be computed from the OUT arguments. Default is void.
        """
        return pulumi.get(self, "returns")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        """
        The schema where the function is located.
        If not specified, the function is created in the current schema.
        """
        return pulumi.get(self, "schema")

