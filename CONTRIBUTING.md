## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/lint
```

This will install all the required dependencies and build the SDK.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## Modifying/Adding code

Changes to this SDK are made directly in this repository — by Cadenya's coding agents and by human
contributors — and land via pull requests or pushes. There is no code generator; edit the source
files directly.

## Adding and running examples

Files in the `examples/` directory can be freely edited or added to.

```go
# add an example to examples/<your-example>/main.go

package main

func main() {
  // ...
}
```

```sh
$ go run ./examples/<your-example>
```

## Using the repository from source

To use a local version of this library from source in another project, edit the `go.mod` with a replace
directive. This can be done through the CLI with the following:

```sh
$ go mod edit -replace go.cadenya.com/cadenya-go=/path/to/cadenya-go
```

## Running tests

```sh
$ ./scripts/test
```

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```
