# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

"""
Tests to ensure that cancelled command submissions behave correctly.
"""
from asyncio import ensure_future

import pytest

from dazl import async_network
from tests.unit.dars import Pending


@pytest.mark.asyncio
async def test_cancelled_write(sandbox):
    async with async_network(url=sandbox, dars=Pending) as network:
        client = network.aio_new_party()

        network.start()

        # Submit a command, but _immediately_ cancel it. Because there are no awaits, this code
        # cannot have possibly been interrupted by the coroutine responsible for scheduling a write
        # to the server, so the command should be cancelled.
        fut = client.create("Pending:Counter", {"owner": client.party, "value": 66})
        ensure_future(fut).cancel()

        # Immediately afterwards, schedule another command submission; this time, we wait for it.
        fut = client.create("Pending:Counter", {"owner": client.party, "value": 7})
        await fut

        data = client.find_active("Pending:Counter")

        assert len(data) == 1
        assert list(data.values())[0]["value"] == 7
