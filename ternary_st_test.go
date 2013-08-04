package trie

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	key := "hello"

	want := simpleType{2, "hahaha value of 2"}

	trie := NewTernaryST()

	trie.Put(key, want)

	tempGot, ok := trie.Get(key)
	if !ok {
		t.Errorf("Put %#v but got '!ok'", want)
	}
	got := tempGot.(simpleType)

	if want != got {
		t.Errorf("Want %#v got %#v", want, got)
	}

}

// func TestDeletes(t *testing.T) {
// 	key := "hello"

// 	want := simpleType{2, "jdhbvjhj"}

// 	trie := NewTernaryST()

// 	trie.Put(key, want)
// 	got, _ := trie.Get(key)
// 	if want != got {
// 		t.Errorf("Want %#v, got %#v", want, got)
// 	}

// 	trie.Delete(key)

// 	got, ok := trie.Get(key)
// 	if ok {
// 		t.Errorf("Want nil, got %#v from trie %#v", got, trie)
// 	}
// }
