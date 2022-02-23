# PEGN Parser Generator

This generator was built by [di-wu](https://github.com/di-wu) and is
based on the original functional AST parser created by
[rwxrob](https://github.com/rwxrob) using a unique, Go-friendly
types-as-operators approach to facilitate the early stages of parser
design and development. The plan is to rewrite this generator with two
modes, the more understandable but less efficient current functional
method, and a more performant single, recursive-descent parser upon
which future regex-like functions will be created.

## Installation

Run the following to generate the `pegn` parser subpackage from the
canonical specification sources:

```shell
go generate
```
