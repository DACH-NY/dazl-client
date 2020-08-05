# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

import json
import logging
from threading import Thread
from time import sleep

from dazl import create, LOG, setup_default_logger, Network, SimplePartyClient
from dazl.model.reading import ReadyEvent
from .dars import PostOffice


def test_simple_flask_integration(sandbox):
    setup_default_logger(logging.INFO)

    network = Network()
    network.set_config(url=sandbox)
    client = network.simple_new_party()

    # seed the ledger with some initial state
    client.add_ledger_ready(create_initial_state)

    network.start_in_background()

    LOG.info('Waiting for the client to be ready...')
    client.ready()

    # TODO: This is currently necessary because there is a bug that prevents ready() from waiting
    #  on ledger_ready events even when those callbacks are added at the appropriate time.
    sleep(10)

    LOG.info('Client is ready.')

    # now start a Flask app
    LOG.info('Starting up the flask app...')
    main_thread = Thread(target=run_flask_app, args=(client, 9999))
    main_thread.start()

    returned_data = run_flask_test(9999)
    assert returned_data == {'postman': client.party}

    network.shutdown()
    network.join(30)

    main_thread.join(30)
    if main_thread.is_alive():
        raise Exception('The Flask thread should have terminated, but did not.')


def create_initial_state(event: 'ReadyEvent'):
    try:
        LOG.info('Uploading our DAR...')
        event.client.ensure_dar(PostOffice)
        while not event.package_store.resolve_template('Main.PostmanRole'):
            logging.info("Waiting for our DAR to be uploaded...")
            sleep(1)

        LOG.info('DAR uploaded. Creating the initial postman role contract...')
        return create('Main.PostmanRole', dict(postman=event.party))
    except:
        LOG.exception('Failed to set up our initial state!')


def run_flask_app(client: SimplePartyClient, port: int) -> None:
    from flask import Flask, jsonify
    app = Flask('TestFlaskIntegration')

    @app.route('/')
    def hello():
        from flask import request

        # shut down Flask as soon as this request is serviced
        func = request.environ.get('werkzeug.server.shutdown')
        func()

        cid, cdata = client.find_one('Main.PostmanRole')
        return jsonify(cdata)

    app.run(host='localhost', port=port)


def run_flask_test(port: int):
    from urllib.request import urlopen
    from time import sleep
    sleep(2)
    with urlopen(f'http://localhost:{port}') as f:
        return json.loads(f.read().decode('utf-8'))
