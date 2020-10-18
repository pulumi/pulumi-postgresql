# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'Clientcert',
]

@pulumi.output_type
class Clientcert(dict):
    def __init__(__self__, *,
                 cert: str,
                 key: str):
        pulumi.set(__self__, "cert", cert)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def cert(self) -> str:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


