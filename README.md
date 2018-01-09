# Blockcypher API Showdown

Showcases usage of the Blockcypher API.

**Note:** You need to have Go 1.8+ installed locally in order to build/test.

## How to run

First, build the project:

    $ go build -o blockcypher-test

Then, set your Blockcypher API Token:

    $ export BLOCKCYPHER_API_TOKEN=<your-token-here>

Lastly, execute the binary:

    $ ./blockcypher-test

## How to test

Simply run:

    $ go test
