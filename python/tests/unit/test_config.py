# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

import http.server
import tempfile

from dazl import Network
from dazl.client.config import ConfigurationError, fetch_config, NetworkConfig
from dazl.model.core import Party
from dazl_internal.background_http_server import TestHTTPServer


# Ensure the key/value pairs that initialize a manager create an appropriate configuration.

def test_party_config_allows_single_string():
    config = NetworkConfig.parse_kwargs(parties='Bob', participant_url='http://nowhere/')
    assert config.parties[0].party == Party('Bob')
    assert config.parties[0].url == 'http://nowhere/'


def test_party_config_allows_list():
    config = NetworkConfig.parse_kwargs(parties=['Bob'], participant_url='http://nowhere/')
    assert config.parties[0].party == Party('Bob')
    assert config.parties[0].url == 'http://nowhere/'


def test_global_url_propagates_to_parties():
    config = NetworkConfig.parse_kwargs(url='http://nowhere/')
    network = Network()
    network.set_config(config)
    client = network.aio_party('Bob')
    party_config = client.resolved_config()
    assert party_config.url == 'http://nowhere/'


# Ensure that the value passed in through --config is either appropriately read from
# a remote HTTP server or the local file system, or that an appropriate error is returned.

def test_fetch_from_url_200():
    with TestHTTPServer(SimpleHandler) as server:
        actual = fetch_config(f'http://localhost:{server.port}/good_path.txt')
    assert SimpleHandler.BODY == actual


def test_fetch_from_url_404():
    url = None
    try:
        with TestHTTPServer(SimpleHandler) as server:
            url = f'http://localhost:{server.port}/bad_path.txt'
            fetch_config(url)
            assert False, 'The fetch should not have succeeded'
    except ConfigurationError as error:
        assert list(error.reasons) == [f'HTTP 404: {url}']


def test_fetch_no_server():
    url = 'http://localhost:65534/something-silly.txt'
    try:
        fetch_config(url)
        assert False, 'The fetch should not have succeeded'
    except ConfigurationError as error:
        pass
        # TODO: The error returned by this failure is currently not the string I'd like to
        #  have returned in this context:
        #  'OSError(99, 'Cannot assign requested address')'
        # self.assertListEqual(error.reasons, [f'HTTP unreachable: {url}'])


def test_fetch_from_file():
    with tempfile.NamedTemporaryFile() as buf:
        buf.file.write('TEST CONFIG'.encode('utf8'))
        buf.file.close()
        actual = fetch_config(buf.name)
    assert 'TEST CONFIG' == actual


class SimpleHandler(http.server.BaseHTTPRequestHandler):

    BODY = '{}'

    # noinspection PyPep8Naming
    def do_GET(self):
        if self.path == '/good_path.txt':
            self.send_response(200)
            self.end_headers()
            self.wfile.write(self.BODY.encode('utf8'))
        else:
            self.send_response(404)
            self.end_headers()

    def log_message(self, format_, *args):
        pass
