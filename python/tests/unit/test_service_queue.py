# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

from asyncio import new_event_loop, gather, sleep, ensure_future, set_event_loop
from unittest import TestCase

from dazl.util.asyncio_util import ServiceQueue


class TestServiceQueue(TestCase):
    def test_simple_queue(self):
        loop = new_event_loop()
        set_event_loop(loop)

        sq = ServiceQueue()
        expected = list(range(5))
        actual = []

        async def main_test():
            async for item in sq:
                print(f'received {item}')
                actual.append(item)

        async def populate():
            await sleep(0.5)
            for i in range(5):
                print(f'putting {i}')
                await sleep(0)
                ensure_future(sq.put(i))
                print(f'putted {i}')
            sq.stop()

        sq.start()
        loop.run_until_complete(gather(main_test(), populate()))

        assert expected == actual
