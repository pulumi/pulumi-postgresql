# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .database import *
from .default_privileg import *
from .default_privileges import *
from .extension import *
from .grant import *
from .grant_role import *
from .physical_replication_slot import *
from .provider import *
from .replication_slot import *
from .role import *
from .schema import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_postgresql.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_postgresql.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "postgresql",
  "mod": "index/database",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/database:Database": "Database"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/defaultPrivileg",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/defaultPrivileg:DefaultPrivileg": "DefaultPrivileg"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/defaultPrivileges",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/defaultPrivileges:DefaultPrivileges": "DefaultPrivileges"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/extension",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/extension:Extension": "Extension"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/grant",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/grant:Grant": "Grant"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/grantRole",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/grantRole:GrantRole": "GrantRole"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/physicalReplicationSlot",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot": "PhysicalReplicationSlot"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/replicationSlot",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/replicationSlot:ReplicationSlot": "ReplicationSlot"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/role",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/role:Role": "Role"
  }
 },
 {
  "pkg": "postgresql",
  "mod": "index/schema",
  "fqn": "pulumi_postgresql",
  "classes": {
   "postgresql:index/schema:Schema": "Schema"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "postgresql",
  "token": "pulumi:providers:postgresql",
  "fqn": "pulumi_postgresql",
  "class": "Provider"
 }
]
"""
)
