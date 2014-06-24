// Package trie holds implementations of a simple trie and of a ternary search trie.
package trie

import (
	"unicode/utf8"
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
	elems     int
}

var (
	// TrieSetIsMutable asserts that a TrieSet can be mutated.
	TrieSetIsMutable MutableSet = new(TrieSet)
)

// NewTrieSet creates a set supporting alphabets of size `alphaSize`.
func NewTrieSet(offset rune, alphaSize uint8) *TrieSet {
	return &TrieSet{
		offset:    uint8(offset),
		alphaSize: alphaSize,
	}
}

// IsEmpty tells whether this TrieSet contains any elements.
func (t *TrieSet) IsEmpty() bool { return t.elems == 0 }

// Len tells how many elements are in this TrieSet.
func (t *TrieSet) Len() int { return t.elems }

// Add puts the key into the set.
func (t *TrieSet) Add(key string) {

	var recurAdd func(*setnode, string, int) *setnode

	recurAdd = func(x *setnode, key string, d int) *setnode {
		if x == nil {
			if d == len(key) {
				t.elems++
			}
			x = newSetNode(t.alphaSize)
		}
		if d == len(key) {
			return x
		}
		r, sz := utf8.DecodeRuneInString(key[d:])
		c := r - rune(t.offset)
		x.Children[c] = recurAdd(x.Children[c], key, d+sz)
		return x
	}

	t.root = recurAdd(t.root, key, 0)
}

// Contains tells if this key is in the set.
func (t *TrieSet) Contains(key string) bool {

	var recurGet func(*setnode, string, int) *setnode

	recurGet = func(x *setnode, key string, d int) *setnode {
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

	setnode := recurGet(t.root, key, 0)

	if setnode == nil {
		return false
	}

	return true
}

// Delete removes the key from the set.
func (t *TrieSet) Delete(key string) {

	var recurDel func(*setnode, string, int) *setnode

	recurDel = func(x *setnode, key string, d int) *setnode {
		if x == nil {
			// was not a key in this TrieSet
			return nil
		}
		if d != len(key) {
			r, sz := utf8.DecodeRuneInString(key[d:])
			c := r - rune(t.offset)
			x.Children[c] = recurDel(x.Children[c], key, d+sz)
			t.elems--
		}

		if x.noChild() {
			return nil
		}
		return x
	}

	recurDel(t.root, key, 0)
}
