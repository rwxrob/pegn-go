# 🌑 PEGN library in Go

This repo contains high-level functions and types for working with Parsing Expression Grammar Notation (PEGN) as defined in the [2023-04 specification](https://github.com/rwxrob/pegn-spec). 

## Difference from previous versions

The 2023-04 version of PEGN specification differs significantly from previous versions primarily by introducing the following changes which greatly simplify the specification and implementations of it:

* `.` - default rune set rule
* `..` - "to" rules
* `p{}` - dynamic unicode support (no need to static class generation)
* `<-` - all rules are named (`<--` operator dropped)
* no formal AST format requirements
* only 14 tokens
* only `Unicode` and `Hexadec` (drop `Binary`)

Related

* <https://pegn.dev>
* <https://github.com/rwxrob/pegn-spec>
