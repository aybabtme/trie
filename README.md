Trie
====

Simple, unoptimized trie implementation for UTF 8 strings or any subset alphabet of UTF 8.

A valid alphabet with size `n` is a subset of UTF 8 with `[0, n)` characters in it.  That means an alphabet of size 128 will effectively be the first 128 characters of UTF8 - from `0x00` to `0x7F`.

You are responsible for ensuring the strings you pass respect the range permissible by the alphabet size you provided. Failing to do so will be rewarded with a runtime panic.

Live long and prosper.

License
=======
MIT, see ./LICENSE
