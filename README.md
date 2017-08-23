# goa example

An example use of [Goa](https://goa.design/).

## Development

Prepare your environment by:
- installing [Go](https://golang.org/)
- installing [Docker](https://www.docker.com/)
- running `make devtools`

Build, test, and lint with

    make

Run with

    make run

Integration test by first starting the Docker containers

    make docker-run

then running integration tests

    make integration-test

then stopping the Docker containers

    make docker-stop

We use [gvt](https://github.com/FiloSottile/gvt) for handling dependencies.
Add a dependency with

    gvt fetch path/to/dependency
