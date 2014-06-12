Trie
====

Implements:

* A trie, use it like a `map[string]interface{}`.
* A ternary search tree, like a `map[string]interface{}` from which you cannot delete.
* A trieset, use it like a `map[string]struct{}`.

Uses package `unsafe` to avoid allocating memory.

Performance
===========

This benchmarks shows lookups only (`Get`):

| keys | `map[string]bool` | `trie.TernaryST` | `trie.Trie` |
|:----:|:-----------------:|:----------------:|:-----------:|
| 4    |        7.36 ns/op |       24.1 ns/op | 28.4 ns/op  |
| 8    |        7.37 ns/op |       24.3 ns/op | 28.3 ns/op  |
| 16   |        28.6 ns/op |       33.4 ns/op | 35.0 ns/op  |
| 32   |        24.6 ns/op |       33.7 ns/op | 35.3 ns/op  |
| 64   |        25.1 ns/op |       33.2 ns/op | 35.2 ns/op  |
| 512  |        24.4 ns/op |       38.5 ns/op | 42.0 ns/op  |
| 1024 |        33.9 ns/op |       43.5 ns/op | 49.5 ns/op  |
| 1M   |        32.1 ns/op |       58.5 ns/op | 72.2 ns/op  |


Known bugs
==========

* Panics on rune that have a value greater than 126.


License
=======
MIT, see ./LICENSE
