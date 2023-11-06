# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'FunctionArg',
    'SchemaPolicy',
    'GetSequencesSequenceResult',
    'GetTablesTableResult',
]

@pulumi.output_type
class FunctionArg(dict):
    def __init__(__self__, *,
                 type: str,
                 default: Optional[str] = None,
                 mode: Optional[str] = None,
                 name: Optional[str] = None):
        """
        :param str type: The type of the argument.
        :param str default: An expression to be used as default value if the parameter is not specified.
        :param str mode: Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        :param str name: The name of the argument.
        """
        FunctionArg._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            type=type,
            default=default,
            mode=mode,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             type: Optional[str] = None,
             default: Optional[str] = None,
             mode: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if type is None:
            raise TypeError("Missing 'type' argument")

        _setter("type", type)
        if default is not None:
            _setter("default", default)
        if mode is not None:
            _setter("mode", mode)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the argument.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def default(self) -> Optional[str]:
        """
        An expression to be used as default value if the parameter is not specified.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        """
        Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class SchemaPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createWithGrant":
            suggest = "create_with_grant"
        elif key == "usageWithGrant":
            suggest = "usage_with_grant"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SchemaPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SchemaPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SchemaPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 create: Optional[bool] = None,
                 create_with_grant: Optional[bool] = None,
                 role: Optional[str] = None,
                 usage: Optional[bool] = None,
                 usage_with_grant: Optional[bool] = None):
        """
        :param bool create: Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        :param bool create_with_grant: Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        :param str role: The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        :param bool usage: Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        :param bool usage_with_grant: Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
               
               > **NOTE on `policy`:** The permissions of a role specified in multiple policy blocks is cumulative.  For example, if the same role is specified in two different `policy` each with different permissions (e.g. `create` and `usage_with_grant`, respectively), then the specified role with have both `create` and `usage_with_grant` privileges.
        """
        SchemaPolicy._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            create=create,
            create_with_grant=create_with_grant,
            role=role,
            usage=usage,
            usage_with_grant=usage_with_grant,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             create: Optional[bool] = None,
             create_with_grant: Optional[bool] = None,
             role: Optional[str] = None,
             usage: Optional[bool] = None,
             usage_with_grant: Optional[bool] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if create_with_grant is None and 'createWithGrant' in kwargs:
            create_with_grant = kwargs['createWithGrant']
        if usage_with_grant is None and 'usageWithGrant' in kwargs:
            usage_with_grant = kwargs['usageWithGrant']

        if create is not None:
            _setter("create", create)
        if create_with_grant is not None:
            _setter("create_with_grant", create_with_grant)
        if role is not None:
            _setter("role", role)
        if usage is not None:
            _setter("usage", usage)
        if usage_with_grant is not None:
            _setter("usage_with_grant", usage_with_grant)

    @property
    @pulumi.getter
    def create(self) -> Optional[bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "create")

    @property
    @pulumi.getter(name="createWithGrant")
    def create_with_grant(self) -> Optional[bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        """
        return pulumi.get(self, "create_with_grant")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        """
        The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def usage(self) -> Optional[bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "usage")

    @property
    @pulumi.getter(name="usageWithGrant")
    def usage_with_grant(self) -> Optional[bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.

        > **NOTE on `policy`:** The permissions of a role specified in multiple policy blocks is cumulative.  For example, if the same role is specified in two different `policy` each with different permissions (e.g. `create` and `usage_with_grant`, respectively), then the specified role with have both `create` and `usage_with_grant` privileges.
        """
        return pulumi.get(self, "usage_with_grant")


@pulumi.output_type
class GetSequencesSequenceResult(dict):
    def __init__(__self__, *,
                 data_type: str,
                 object_name: str,
                 schema_name: str):
        """
        :param str data_type: The sequence's data type as defined in ``information_schema.sequences``.
        :param str object_name: The sequence name.
        :param str schema_name: The parent schema.
        """
        GetSequencesSequenceResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            data_type=data_type,
            object_name=object_name,
            schema_name=schema_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             data_type: Optional[str] = None,
             object_name: Optional[str] = None,
             schema_name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if data_type is None and 'dataType' in kwargs:
            data_type = kwargs['dataType']
        if data_type is None:
            raise TypeError("Missing 'data_type' argument")
        if object_name is None and 'objectName' in kwargs:
            object_name = kwargs['objectName']
        if object_name is None:
            raise TypeError("Missing 'object_name' argument")
        if schema_name is None and 'schemaName' in kwargs:
            schema_name = kwargs['schemaName']
        if schema_name is None:
            raise TypeError("Missing 'schema_name' argument")

        _setter("data_type", data_type)
        _setter("object_name", object_name)
        _setter("schema_name", schema_name)

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> str:
        """
        The sequence's data type as defined in ``information_schema.sequences``.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> str:
        """
        The sequence name.
        """
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> str:
        """
        The parent schema.
        """
        return pulumi.get(self, "schema_name")


@pulumi.output_type
class GetTablesTableResult(dict):
    def __init__(__self__, *,
                 object_name: str,
                 schema_name: str,
                 table_type: str):
        """
        :param str object_name: The table name.
        :param str schema_name: The parent schema.
        :param str table_type: The table type as defined in ``information_schema.tables``.
        """
        GetTablesTableResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            object_name=object_name,
            schema_name=schema_name,
            table_type=table_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             object_name: Optional[str] = None,
             schema_name: Optional[str] = None,
             table_type: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if object_name is None and 'objectName' in kwargs:
            object_name = kwargs['objectName']
        if object_name is None:
            raise TypeError("Missing 'object_name' argument")
        if schema_name is None and 'schemaName' in kwargs:
            schema_name = kwargs['schemaName']
        if schema_name is None:
            raise TypeError("Missing 'schema_name' argument")
        if table_type is None and 'tableType' in kwargs:
            table_type = kwargs['tableType']
        if table_type is None:
            raise TypeError("Missing 'table_type' argument")

        _setter("object_name", object_name)
        _setter("schema_name", schema_name)
        _setter("table_type", table_type)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> str:
        """
        The table name.
        """
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> str:
        """
        The parent schema.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter(name="tableType")
    def table_type(self) -> str:
        """
        The table type as defined in ``information_schema.tables``.
        """
        return pulumi.get(self, "table_type")


