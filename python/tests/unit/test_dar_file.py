# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

from dazl.util.dar import DarFile
from .dars import Pending


def test_get_sdk_version():
    with DarFile(Pending) as dar:
        print(dar.get_manifest())
        assert '1.3.0' == dar.get_sdk_version()
