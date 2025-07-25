# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = [
    'FunctionArg',
    'SchemaPolicy',
    'GetSequencesSequenceResult',
    'GetTablesTableResult',
]

@pulumi.output_type
class FunctionArg(dict):
    def __init__(__self__, *,
                 type: _builtins.str,
                 default: Optional[_builtins.str] = None,
                 mode: Optional[_builtins.str] = None,
                 name: Optional[_builtins.str] = None):
        """
        :param _builtins.str type: The type of the argument.
        :param _builtins.str default: An expression to be used as default value if the parameter is not specified.
        :param _builtins.str mode: Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        :param _builtins.str name: The name of the argument.
        """
        pulumi.set(__self__, "type", type)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @_builtins.property
    @pulumi.getter
    def type(self) -> _builtins.str:
        """
        The type of the argument.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter
    def default(self) -> Optional[_builtins.str]:
        """
        An expression to be used as default value if the parameter is not specified.
        """
        return pulumi.get(self, "default")

    @_builtins.property
    @pulumi.getter
    def mode(self) -> Optional[_builtins.str]:
        """
        Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        """
        return pulumi.get(self, "mode")

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[_builtins.str]:
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
                 create: Optional[_builtins.bool] = None,
                 create_with_grant: Optional[_builtins.bool] = None,
                 role: Optional[_builtins.str] = None,
                 usage: Optional[_builtins.bool] = None,
                 usage_with_grant: Optional[_builtins.bool] = None):
        """
        :param _builtins.bool create: Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        :param _builtins.bool create_with_grant: Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        :param _builtins.str role: The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        :param _builtins.bool usage: Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        :param _builtins.bool usage_with_grant: Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
               
               > **NOTE on `policy`:** The permissions of a role specified in multiple policy blocks is cumulative.  For example, if the same role is specified in two different `policy` each with different permissions (e.g. `create` and `usage_with_grant`, respectively), then the specified role with have both `create` and `usage_with_grant` privileges.
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if create_with_grant is not None:
            pulumi.set(__self__, "create_with_grant", create_with_grant)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if usage_with_grant is not None:
            pulumi.set(__self__, "usage_with_grant", usage_with_grant)

    @_builtins.property
    @pulumi.getter
    def create(self) -> Optional[_builtins.bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "create")

    @_builtins.property
    @pulumi.getter(name="createWithGrant")
    def create_with_grant(self) -> Optional[_builtins.bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        """
        return pulumi.get(self, "create_with_grant")

    @_builtins.property
    @pulumi.getter
    def role(self) -> Optional[_builtins.str]:
        """
        The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        """
        return pulumi.get(self, "role")

    @_builtins.property
    @pulumi.getter
    def usage(self) -> Optional[_builtins.bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "usage")

    @_builtins.property
    @pulumi.getter(name="usageWithGrant")
    def usage_with_grant(self) -> Optional[_builtins.bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.

        > **NOTE on `policy`:** The permissions of a role specified in multiple policy blocks is cumulative.  For example, if the same role is specified in two different `policy` each with different permissions (e.g. `create` and `usage_with_grant`, respectively), then the specified role with have both `create` and `usage_with_grant` privileges.
        """
        return pulumi.get(self, "usage_with_grant")


@pulumi.output_type
class GetSequencesSequenceResult(dict):
    def __init__(__self__, *,
                 data_type: _builtins.str,
                 object_name: _builtins.str,
                 schema_name: _builtins.str):
        """
        :param _builtins.str data_type: The sequence's data type as defined in ``information_schema.sequences``.
        :param _builtins.str object_name: The sequence name.
        :param _builtins.str schema_name: The parent schema.
        """
        pulumi.set(__self__, "data_type", data_type)
        pulumi.set(__self__, "object_name", object_name)
        pulumi.set(__self__, "schema_name", schema_name)

    @_builtins.property
    @pulumi.getter(name="dataType")
    def data_type(self) -> _builtins.str:
        """
        The sequence's data type as defined in ``information_schema.sequences``.
        """
        return pulumi.get(self, "data_type")

    @_builtins.property
    @pulumi.getter(name="objectName")
    def object_name(self) -> _builtins.str:
        """
        The sequence name.
        """
        return pulumi.get(self, "object_name")

    @_builtins.property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> _builtins.str:
        """
        The parent schema.
        """
        return pulumi.get(self, "schema_name")


@pulumi.output_type
class GetTablesTableResult(dict):
    def __init__(__self__, *,
                 object_name: _builtins.str,
                 schema_name: _builtins.str,
                 table_type: _builtins.str):
        """
        :param _builtins.str object_name: The table name.
        :param _builtins.str schema_name: The parent schema.
        :param _builtins.str table_type: The table type as defined in ``information_schema.tables``.
        """
        pulumi.set(__self__, "object_name", object_name)
        pulumi.set(__self__, "schema_name", schema_name)
        pulumi.set(__self__, "table_type", table_type)

    @_builtins.property
    @pulumi.getter(name="objectName")
    def object_name(self) -> _builtins.str:
        """
        The table name.
        """
        return pulumi.get(self, "object_name")

    @_builtins.property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> _builtins.str:
        """
        The parent schema.
        """
        return pulumi.get(self, "schema_name")

    @_builtins.property
    @pulumi.getter(name="tableType")
    def table_type(self) -> _builtins.str:
        """
        The table type as defined in ``information_schema.tables``.
        """
        return pulumi.get(self, "table_type")


