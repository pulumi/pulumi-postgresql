# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PublicationArgs', 'Publication']

@pulumi.input_type
class PublicationArgs:
    def __init__(__self__, *,
                 all_tables: Optional[pulumi.Input[bool]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 publish_params: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 publish_via_partition_root_param: Optional[pulumi.Input[bool]] = None,
                 tables: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Publication resource.
        :param pulumi.Input[bool] all_tables: Should be ALL TABLES added to the publication. Defaults to 'false'
        :param pulumi.Input[str] database: Which database to create the publication on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: Should all subsequent resources of the publication be dropped. Defaults to 'false'
        :param pulumi.Input[str] name: The name of the publication.
        :param pulumi.Input[str] owner: Who owns the publication. Defaults to provider user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] publish_params: Which 'publish' options should be turned on. Default to 'insert','update','delete'
        :param pulumi.Input[bool] publish_via_partition_root_param: Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tables: Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        if all_tables is not None:
            pulumi.set(__self__, "all_tables", all_tables)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if publish_params is not None:
            pulumi.set(__self__, "publish_params", publish_params)
        if publish_via_partition_root_param is not None:
            pulumi.set(__self__, "publish_via_partition_root_param", publish_via_partition_root_param)
        if tables is not None:
            pulumi.set(__self__, "tables", tables)

    @property
    @pulumi.getter(name="allTables")
    def all_tables(self) -> Optional[pulumi.Input[bool]]:
        """
        Should be ALL TABLES added to the publication. Defaults to 'false'
        """
        return pulumi.get(self, "all_tables")

    @all_tables.setter
    def all_tables(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "all_tables", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Which database to create the publication on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        Should all subsequent resources of the publication be dropped. Defaults to 'false'
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the publication.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Who owns the publication. Defaults to provider user.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="publishParams")
    def publish_params(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Which 'publish' options should be turned on. Default to 'insert','update','delete'
        """
        return pulumi.get(self, "publish_params")

    @publish_params.setter
    def publish_params(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "publish_params", value)

    @property
    @pulumi.getter(name="publishViaPartitionRootParam")
    def publish_via_partition_root_param(self) -> Optional[pulumi.Input[bool]]:
        """
        Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        """
        return pulumi.get(self, "publish_via_partition_root_param")

    @publish_via_partition_root_param.setter
    def publish_via_partition_root_param(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_via_partition_root_param", value)

    @property
    @pulumi.getter
    def tables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        return pulumi.get(self, "tables")

    @tables.setter
    def tables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tables", value)


@pulumi.input_type
class _PublicationState:
    def __init__(__self__, *,
                 all_tables: Optional[pulumi.Input[bool]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 publish_params: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 publish_via_partition_root_param: Optional[pulumi.Input[bool]] = None,
                 tables: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Publication resources.
        :param pulumi.Input[bool] all_tables: Should be ALL TABLES added to the publication. Defaults to 'false'
        :param pulumi.Input[str] database: Which database to create the publication on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: Should all subsequent resources of the publication be dropped. Defaults to 'false'
        :param pulumi.Input[str] name: The name of the publication.
        :param pulumi.Input[str] owner: Who owns the publication. Defaults to provider user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] publish_params: Which 'publish' options should be turned on. Default to 'insert','update','delete'
        :param pulumi.Input[bool] publish_via_partition_root_param: Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tables: Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        if all_tables is not None:
            pulumi.set(__self__, "all_tables", all_tables)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if publish_params is not None:
            pulumi.set(__self__, "publish_params", publish_params)
        if publish_via_partition_root_param is not None:
            pulumi.set(__self__, "publish_via_partition_root_param", publish_via_partition_root_param)
        if tables is not None:
            pulumi.set(__self__, "tables", tables)

    @property
    @pulumi.getter(name="allTables")
    def all_tables(self) -> Optional[pulumi.Input[bool]]:
        """
        Should be ALL TABLES added to the publication. Defaults to 'false'
        """
        return pulumi.get(self, "all_tables")

    @all_tables.setter
    def all_tables(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "all_tables", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Which database to create the publication on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        Should all subsequent resources of the publication be dropped. Defaults to 'false'
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the publication.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Who owns the publication. Defaults to provider user.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="publishParams")
    def publish_params(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Which 'publish' options should be turned on. Default to 'insert','update','delete'
        """
        return pulumi.get(self, "publish_params")

    @publish_params.setter
    def publish_params(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "publish_params", value)

    @property
    @pulumi.getter(name="publishViaPartitionRootParam")
    def publish_via_partition_root_param(self) -> Optional[pulumi.Input[bool]]:
        """
        Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        """
        return pulumi.get(self, "publish_via_partition_root_param")

    @publish_via_partition_root_param.setter
    def publish_via_partition_root_param(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_via_partition_root_param", value)

    @property
    @pulumi.getter
    def tables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        return pulumi.get(self, "tables")

    @tables.setter
    def tables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tables", value)


class Publication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_tables: Optional[pulumi.Input[bool]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 publish_params: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 publish_via_partition_root_param: Optional[pulumi.Input[bool]] = None,
                 tables: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The `Publication` resource creates and manages a publication on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        publication = postgresql.Publication("publication", tables=[
            "public.test",
            "another_schema.test",
        ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_tables: Should be ALL TABLES added to the publication. Defaults to 'false'
        :param pulumi.Input[str] database: Which database to create the publication on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: Should all subsequent resources of the publication be dropped. Defaults to 'false'
        :param pulumi.Input[str] name: The name of the publication.
        :param pulumi.Input[str] owner: Who owns the publication. Defaults to provider user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] publish_params: Which 'publish' options should be turned on. Default to 'insert','update','delete'
        :param pulumi.Input[bool] publish_via_partition_root_param: Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tables: Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PublicationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `Publication` resource creates and manages a publication on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        publication = postgresql.Publication("publication", tables=[
            "public.test",
            "another_schema.test",
        ])
        ```

        :param str resource_name: The name of the resource.
        :param PublicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PublicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_tables: Optional[pulumi.Input[bool]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 publish_params: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 publish_via_partition_root_param: Optional[pulumi.Input[bool]] = None,
                 tables: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PublicationArgs.__new__(PublicationArgs)

            __props__.__dict__["all_tables"] = all_tables
            __props__.__dict__["database"] = database
            __props__.__dict__["drop_cascade"] = drop_cascade
            __props__.__dict__["name"] = name
            __props__.__dict__["owner"] = owner
            __props__.__dict__["publish_params"] = publish_params
            __props__.__dict__["publish_via_partition_root_param"] = publish_via_partition_root_param
            __props__.__dict__["tables"] = tables
        super(Publication, __self__).__init__(
            'postgresql:index/publication:Publication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_tables: Optional[pulumi.Input[bool]] = None,
            database: Optional[pulumi.Input[str]] = None,
            drop_cascade: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            publish_params: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            publish_via_partition_root_param: Optional[pulumi.Input[bool]] = None,
            tables: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Publication':
        """
        Get an existing Publication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_tables: Should be ALL TABLES added to the publication. Defaults to 'false'
        :param pulumi.Input[str] database: Which database to create the publication on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: Should all subsequent resources of the publication be dropped. Defaults to 'false'
        :param pulumi.Input[str] name: The name of the publication.
        :param pulumi.Input[str] owner: Who owns the publication. Defaults to provider user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] publish_params: Which 'publish' options should be turned on. Default to 'insert','update','delete'
        :param pulumi.Input[bool] publish_via_partition_root_param: Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tables: Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PublicationState.__new__(_PublicationState)

        __props__.__dict__["all_tables"] = all_tables
        __props__.__dict__["database"] = database
        __props__.__dict__["drop_cascade"] = drop_cascade
        __props__.__dict__["name"] = name
        __props__.__dict__["owner"] = owner
        __props__.__dict__["publish_params"] = publish_params
        __props__.__dict__["publish_via_partition_root_param"] = publish_via_partition_root_param
        __props__.__dict__["tables"] = tables
        return Publication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allTables")
    def all_tables(self) -> pulumi.Output[bool]:
        """
        Should be ALL TABLES added to the publication. Defaults to 'false'
        """
        return pulumi.get(self, "all_tables")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Which database to create the publication on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> pulumi.Output[Optional[bool]]:
        """
        Should all subsequent resources of the publication be dropped. Defaults to 'false'
        """
        return pulumi.get(self, "drop_cascade")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the publication.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Who owns the publication. Defaults to provider user.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="publishParams")
    def publish_params(self) -> pulumi.Output[Sequence[str]]:
        """
        Which 'publish' options should be turned on. Default to 'insert','update','delete'
        """
        return pulumi.get(self, "publish_params")

    @property
    @pulumi.getter(name="publishViaPartitionRootParam")
    def publish_via_partition_root_param(self) -> pulumi.Output[Optional[bool]]:
        """
        Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        """
        return pulumi.get(self, "publish_via_partition_root_param")

    @property
    @pulumi.getter
    def tables(self) -> pulumi.Output[Sequence[str]]:
        """
        Which tables add to the publication. By defaults no tables added. Format of table is `<schema_name>.<table_name>`. If `<schema_name>` is not specified - default database schema will be used.  Table string must be listed in alphabetical order.
        """
        return pulumi.get(self, "tables")

