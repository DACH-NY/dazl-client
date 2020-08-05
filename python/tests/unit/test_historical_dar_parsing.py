# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

from pathlib import Path

from dazl import LOG
from dazl.util.dar import DarFile

ARCHIVES: Path = Path(__file__).absolute().parent.parent.parent.parent / '_fixtures' / 'archives'


def test_dar_version_compatibility(subtests):
    dars = list(ARCHIVES.glob('**/*.dar'))
    assert len(dars) > 0, f'Could not find any DARs in {ARCHIVES}'

    for dar in ARCHIVES.glob('**/*.dar'):
        short_dar = dar.relative_to(ARCHIVES)
        with subtests.test(str(short_dar)):
            dar_file = DarFile(dar)
            metadata = dar_file.read_metadata()
            LOG.info('Successfully read %s with package IDs %r.', short_dar,
                     metadata.package_ids())

