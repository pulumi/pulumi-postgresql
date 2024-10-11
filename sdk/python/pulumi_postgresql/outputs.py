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
        pulumi.set(__self__, "type", type)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
        pulumi.set(__self__, "data_type", data_type)
        pulumi.set(__self__, "object_name", object_name)
        pulumi.set(__self__, "schema_name", schema_name)

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
        pulumi.set(__self__, "object_name", object_name)
        pulumi.set(__self__, "schema_name", schema_name)
        pulumi.set(__self__, "table_type", table_type)

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


