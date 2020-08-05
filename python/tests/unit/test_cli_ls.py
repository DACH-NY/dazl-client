# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

"""
Tests to ensure that CLI commands work properly.
"""
from dazl.cli import _main


def test_simple_ls(sandbox):
    exit_code = _main(f'dazl ls --url {sandbox} --parties=Alice'.split(' '))
    assert exit_code == 0


def test_simple_ls_two_parties(sandbox):
    exit_code = _main(f'dazl ls --url {sandbox} --parties=Alice,Bob'.split(' '))

    assert exit_code == 0


def test_env_ls(sandbox):
    import os
    os.environ['DAML_LEDGER_URL'] = sandbox
    os.environ['DAML_LEDGER_PARTY'] = 'Alice'
    exit_code = _main('dazl ls'.split(' '))

    assert exit_code == 0
