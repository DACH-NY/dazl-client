# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

import logging
import uuid

import pytest

from dazl import async_network, create, Party

from .dars import AllParty as AllPartyDar

PrivateContract = 'AllParty.PrivateContract'
PublicContract = 'AllParty.PublicContract'


@pytest.mark.asyncio
async def test_some_party_receives_public_contract(sandbox):
    some_party_cids = []
    publisher_cids = []

    # TODO: Switch to a Party allocation API when available.
    all_party = Party(str(uuid.uuid4()))

    async with async_network(url=sandbox, dars=AllPartyDar) as network:
        network.set_config(party_groups=[all_party])

        some_client = network.aio_new_party()
        some_client.add_ledger_ready(
            lambda _: create(PrivateContract, {'someParty': some_client.party}))

        publisher_client = network.aio_new_party()
        publisher_client.add_ledger_ready(
            lambda _: create(PublicContract, {'publisher': publisher_client.party, 'allParty': all_party}))

        some_client.add_ledger_created(PublicContract, lambda e: some_party_cids.append(e.cid))
        some_client.add_ledger_created(PrivateContract, lambda e: some_party_cids.append(e.cid))

        publisher_client.add_ledger_created(PublicContract, lambda e: publisher_cids.append(e.cid))
        publisher_client.add_ledger_created(PrivateContract, lambda e: publisher_cids.append(e.cid))

        network.start()

    logging.info(
        'got to the end with some_party contracts: %s and publisher contracts: %s',
        some_party_cids, publisher_cids)

    assert len(some_party_cids) == 2
    assert len(publisher_cids) == 1
