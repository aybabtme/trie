// Package trie is a simple trie implementation for UTF 8 strings or any subset
// alphabet of UTF 8.
//
// A valid alphabet with size `n` is a subset of UTF 8 with `[0, n)` characters
// in it.  That means an alphabet of size 128 will effectively be the first 128
// characters of UTF8 - from `0x00` to `0x7F`.
//
// You are responsible for ensuring the strings you pass respect the range
// permissible by the alphabet size you provided. Failing to do so will be
// rewarded with a runtime panic.
package trie

import (
	"unicode/utf8"
)

type node struct {
	Value    interface{}
	Children []*node
}

func newNode(n int, val interface{}) *node {
	return &node{val, make([]*node, n)}
}

func (n *node) noChild() bool {
	for _, n := range n.Children {
		if n != nil {
			return false
		}
	}
	return true
}

// Trie is a symbol table specifically for string indexed keys.
type Trie struct {
	root      *node
	alphaSize int
}

// NewTrie creates a trie supporting alphabets of size `alphaSize`.
func NewTrie(alphaSize int) *Trie {
	return &Trie{
		alphaSize: alphaSize,
	}
}

// NewUTF8 creates a trie supporting a UTF8 alphabet.
func NewUTF8() *Trie {
	return NewTrie(utf8.MaxRune)
}

// Put puts the value `val` into the trie at key `key`.
func (t *Trie) Put(key string, val interface{}) {

	var recurPut func(*node, string, interface{}, int) *node

	recurPut = func(x *node, key string, val interface{}, d int) *node {
		if x == nil {
			x = newNode(t.alphaSize, val)
		}
		if d == len(key) {
			x.Value = val
			return x
		}
		c := key[d]
		x.Children[c] = recurPut(x.Children[c], key, val, d+1)
		return x
	}

	t.root = recurPut(t.root, key, val, 0)
}

// Get returns the value found at this location, if it exists
func (t *Trie) Get(key string) (interface{}, bool) {

	var recurGet func(*node, string, int) *node

	recurGet = func(x *node, key string, d int) *node {
		if x == nil {
			return nil
		}
		if d == len(key) {
			return x
		}
		c := key[d]
		return recurGet(x.Children[c], key, d+1)
	}

	node := recurGet(t.root, key, 0)

	if node == nil {
		return new(interface{}), false
	}

	return node.Value, true
}

// Delete removes the value found at this location, if it exists
func (t *Trie) Delete(key string) {
	var recurDel func(*node, string, int) *node

	recurDel = func(x *node, key string, d int) *node {
		if x == nil {
			// was not a key in this Trie
			return nil
		}
		if d == len(key) {
			x.Value = nil
		} else {
			c := key[d]
			x.Children[c] = recurDel(x.Children[c], key, d+1)
		}

		if x.Value != nil {
			return x
		}
		if x.noChild() {
			return nil
		}
		return x
	}

	recurDel(t.root, key, 0)
}
