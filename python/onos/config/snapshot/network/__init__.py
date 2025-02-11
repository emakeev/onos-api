# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/snapshot/network/types.proto
# plugin: python-betterproto
from dataclasses import dataclass
from datetime import datetime
from typing import List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase


@dataclass(eq=False, repr=False)
class NetworkSnapshot(betterproto.Message):
    """NetworkSnapshot is a snapshot of all network changes"""

    # 'id' is the unique snapshot identifier
    id: str = betterproto.string_field(1)
    # 'index' is a monotonically increasing, globally unique snapshot index
    index: int = betterproto.uint64_field(2)
    # 'revision' is the request revision number
    revision: int = betterproto.uint64_field(3)
    # 'status' is the snapshot status
    status: "__snapshot__.Status" = betterproto.message_field(4)
    # 'retention' specifies the duration for which to retain changes
    retention: "__snapshot__.RetentionOptions" = betterproto.message_field(6)
    # 'created' is the time at which the configuration was created
    created: datetime = betterproto.message_field(8)
    # 'updated' is the time at which the configuration was last updated
    updated: datetime = betterproto.message_field(9)
    # 'refs' is a set of references to stored device snapshots
    refs: List["DeviceSnapshotRef"] = betterproto.message_field(10)


@dataclass(eq=False, repr=False)
class DeviceSnapshotRef(betterproto.Message):
    """DeviceSnapshotRef is a reference to a device snapshot"""

    # 'device_snapshot_id' is the unique identifier of the device snapshot
    device_snapshot_id: str = betterproto.string_field(1)


from ... import snapshot as __snapshot__
