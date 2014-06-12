package trie

import (
	"testing"
)

type simpleType struct {
	n    int
	name string
}

func TestPutGetWith128CharAlphabet(t *testing.T) {
	key := "hello"

	want := simpleType{2, "hahaha value of 2"}

	trie := NewTrie('A', 128)

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

func TestDeletesWith128CharAlphabet(t *testing.T) {
	key := "hello"

	want := simpleType{2, "jdhbvjhj"}

	trie := NewTrie('A', 128)

	trie.Put(key, want)
	got, _ := trie.Get(key)
	if want != got {
		t.Errorf("Want %#v, got %#v", want, got)
	}

	trie.Delete(key)

	got, ok := trie.Get(key)
	if ok {
		t.Errorf("Want nil, got %#v from trie %#v", got, trie)
	}
}
