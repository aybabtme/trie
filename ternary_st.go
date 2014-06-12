package trie

import (
	"reflect"
	"unsafe"
)

type ternNode struct {
	Value *interface{}
	Code  rune

	left  *ternNode
	child *ternNode
	right *ternNode
}

// TernaryST is a symbol table specifically for string indexed keys.
type TernaryST struct {
	root  *ternNode
	count int
}

// NewTernaryST creates a trie.
func NewTernaryST() *TernaryST {
	return &TernaryST{nil, 0}
}

// Put puts the value `val` into the trie at key `key`.
func (t *TernaryST) Put(key string, val interface{}) {
	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	t.root = t.put(t.root, data, val, 0)
}

func (t *TernaryST) put(x *ternNode, key []uint8, val interface{}, d int) *ternNode {
	c := key[d]
	if x == nil {
		x = &ternNode{Code: rune(c), Value: &val}
	}

	if c < uint8(x.Code) {
		x.left = t.put(x.left, key, val, d)
	} else if c > uint8(x.Code) {
		x.right = t.put(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.child = t.put(x.child, key, val, d+1)
	} else {
		if x.Value == nil {
			t.count++
		}
		x.Value = &val
	}
	return x
}

// Get returns the value found at this location, if it exists
func (t *TernaryST) Get(key string) (interface{}, bool) {

	if key == "" {
		return nil, false
	}

	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	ternNode := t.get(t.root, data, 0)

	if ternNode == nil {
		return nil, false
	}

	return *ternNode.Value, true
}

func (t *TernaryST) get(x *ternNode, key []uint8, d int) *ternNode {
	if x == nil {
		return nil
	}

	c := key[d]

	if c < uint8(x.Code) {
		return t.get(x.left, key, d)
	} else if c > uint8(x.Code) {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.child, key, d+1)
	}

	return x
}

// Delete removes the value found at this location, if it exists
func (t *TernaryST) Delete(key string) {
	if key == "" {
		return
	}
	t.delete(t.root, []rune(key), 0)
}

func (t *TernaryST) delete(*ternNode, []rune, int) *ternNode {
	panic("Not implemented yet!")
	return nil
}

// Len returns the count of elements in this trie
func (t *TernaryST) Len() int {
	return t.count
}

// Keys returns all the keys known to this trie
func (t *TernaryST) Keys() []string {
	var outCollection []string
	collect(t.root, []uint8{}, outCollection)
	return outCollection
}

// KeysWithPrefix returns all the keys starting with prefix `key`
func (t *TernaryST) KeysWithPrefix(key string) []string {
	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	var outCollection []string
	x := t.get(t.root, data, 0)
	if x == nil {
		return outCollection
	}
	if x.Value != nil {
		outCollection = append(outCollection, key)
	}
	collect(x.child, data, outCollection)
	return outCollection
}

// LongestPrefix returns the longest string in this trie that has `key` for prefix
func (t *TernaryST) LongestPrefix(key string) string {
	if key == "" {
		return ""
	}

	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	lenght := 0

	x := t.root
	i := 0
	var c uint8
	for x != nil && i < len(data) {
		c = data[i]
		if c < uint8(x.Code) {
			x = x.left
		} else if c > uint8(x.Code) {
			x = x.right
		} else {
			i++
			if x.Value != nil {
				lenght = i
			}
			x = x.child
		}
	}
	return string(data[0:lenght])
}

// KeysMatching returns all the keys that share prefix `key`, where `key` can
// contain a wildcard character `.`.
func (t *TernaryST) KeysMatching(key string) []string {
	header := *(*reflect.StringHeader)(unsafe.Pointer(&key))
	data := *(*[]uint8)(unsafe.Pointer(&header))
	var outCollection []string
	patternCollect(t.root, []uint8{}, 0, data, outCollection)
	return outCollection
}

// Helpers

// `collect` collects a set of string that share a prefix
func collect(x *ternNode, key []uint8, outCollection []string) {
	if x == nil {
		return
	}
	collect(x.left, key, outCollection)
	newKey := append(key, uint8(x.Code))
	if x.Value != nil {
		outCollection = append(outCollection, string(newKey))
	}
	collect(x.child, newKey, outCollection)
	collect(x.right, key, outCollection)
}

// 'patternCollect' collects a set of string that matches the pattern
func patternCollect(
	x *ternNode,
	prefix []uint8,
	i int,
	pat []uint8,
	outCollection []string,
) {
	if x == nil {
		return
	}

	c := pat[i]

	if c == uint8('.') || c < uint8(x.Code) {
		patternCollect(x.left, prefix, i, pat, outCollection)
	}

	if c == uint8('.') || c == uint8(x.Code) {
		newKey := append(prefix, c)
		if i == len(pat)-1 && x.Value != nil {
			outCollection = append(outCollection, string(newKey))
		}
		if i < len(pat)-1 {
			patternCollect(x.child, newKey, i+1, pat, outCollection)
		}
	}

	if c == uint8('.') || c > uint8(x.Code) {
		patternCollect(x.right, prefix, i, pat, outCollection)
	}
}
