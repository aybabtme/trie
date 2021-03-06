// Package trie holds implementations of a simple trie and of a ternary search trie.
package trie

import (
	"log"
	"unicode/utf8"
)

type node struct {
	Value    interface{}
	Children []*node
}

func newNode(n uint8, val interface{}) *node {
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
	offset    uint8
	root      *node
	alphaSize uint8
}

// NewTrie creates a trie supporting alphabets of size `alphaSize`.
func NewTrie(offset byte, alphaSize uint8) *Trie {
	return &Trie{
		offset:    uint8(offset),
		alphaSize: alphaSize,
	}
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
		r, sz := utf8.DecodeRuneInString(key[d:])
		c := r - rune(t.offset)
		if len(x.Children) < int(c) {
			log.Panicf("key=%q\tkey[d]=%v\td=%d\tc=%d\tt.offset=%d\tlen(x.Children)=%d", key, key[d], d, c, t.offset, len(x.Children))
		}
		x.Children[c] = recurPut(x.Children[c], key, val, d+sz)
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
		r, sz := utf8.DecodeRuneInString(key[d:])
		c := r - rune(t.offset)
		return recurGet(x.Children[c], key, d+sz)
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
			r, sz := utf8.DecodeRuneInString(key[d:])
			c := r - rune(t.offset)
			x.Children[c] = recurDel(x.Children[c], key, d+sz)
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
