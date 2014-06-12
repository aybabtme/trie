// Package trie holds implementations of a simple trie and of a ternary search trie.
package trie

import (
	"reflect"
	"unsafe"
)

type setnode struct {
	Children []*setnode
}

func newSetNode(n uint8) *setnode {
	return &setnode{make([]*setnode, n)}
}

func (n *setnode) noChild() bool {
	for _, n := range n.Children {
		if n != nil {
			return false
		}
	}
	return true
}

// TrieSet is a set of string backed by a trie structure.
type TrieSet struct {
	offset    uint8
	root      *setnode
	alphaSize uint8
}

// NewTrieSet creates a set supporting alphabets of size `alphaSize`.
func NewTrieSet(offset rune, alphaSize uint8) *TrieSet {
	return &TrieSet{
		offset:    uint8(offset),
		alphaSize: alphaSize,
	}
}

// Add puts the key into the set.
func (t *TrieSet) Add(key string) {

	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	var recurAdd func(*setnode, []uint8, int) *setnode

	recurAdd = func(x *setnode, key []uint8, d int) *setnode {
		if x == nil {
			x = newSetNode(t.alphaSize)
		}
		if d == len(key) {
			return x
		}
		c := key[d] - t.offset
		x.Children[c] = recurAdd(x.Children[c], key, d+1)
		return x
	}

	t.root = recurAdd(t.root, data, 0)
}

// Contains tells if this key is in the set.
func (t *TrieSet) Contains(key string) bool {

	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	var recurGet func(*setnode, []uint8, int) *setnode

	recurGet = func(x *setnode, key []uint8, d int) *setnode {
		if x == nil {
			return nil
		}
		if d == len(key) {
			return x
		}
		c := key[d] - t.offset
		return recurGet(x.Children[c], key, d+1)
	}

	setnode := recurGet(t.root, data, 0)

	if setnode == nil {
		return false
	}

	return true
}

// Delete removes the key from the set.
func (t *TrieSet) Delete(key string) {

	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	var recurDel func(*setnode, []uint8, int) *setnode

	recurDel = func(x *setnode, key []uint8, d int) *setnode {
		if x == nil {
			// was not a key in this TrieSet
			return nil
		}
		if d != len(key) {
			c := key[d] - t.offset
			x.Children[c] = recurDel(x.Children[c], key, d+1)
		}

		if x.noChild() {
			return nil
		}
		return x
	}

	recurDel(t.root, data, 0)
}
