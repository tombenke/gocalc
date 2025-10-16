# gocalc

[![Actions Status](https://github.com/tombenke/gocalc/workflows/ci/badge.svg)](https://github.com/tombenke/gocalc)

## About

gocalc is a built-in calculator package for Golang.

This project was made to demonstrate how to use the [parc](https://github.com/tombenke/parc) parser combinator package
to implement a language for a calculator that can parse simple arithmetic formulas,
then compile to an executable array of instructions of a stack machine.

The parser, the compiler and the runner are all included in the package, and can be embedded into any go application.

See the API docs by starting the docs generator and server using the `task docs` task after installation.

## Development

Clone the repository, then install the dependencies and the development tools:

```bash
task install
```

List the tasks:

```bash
task list
```
