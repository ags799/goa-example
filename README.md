# goa example

An example use of [Goa](https://goa.design/).

## Development

Prepare your environment by:
- installing [Go](https://golang.org/)
- installing [Goa](https://goa.design/). Note that the project uses Goa v1.
- running `make devtools`

Test and lint with

    make

Run with

    make run

We use [gvt](https://github.com/FiloSottile/gvt) for handling dependencies.
Add a dependency with

    gvt fetch path/to/dependency
