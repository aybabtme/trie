package trie

import (
	"testing"
)

func TestAddContainsWith128CharAlphabet(t *testing.T) {
	key := "hello"

	set := NewTrieSet('A', 128)

	set.Add(key)

	want := true
	got := set.Contains(key)

	if want != got {
		t.Errorf("Want %#v got %#v", want, got)
	}
}

func TestAddDeletesWith128CharAlphabet(t *testing.T) {
	key := "hello"

	set := NewTrieSet('A', 128)

	set.Add(key)

	want := true
	got := set.Contains(key)

	if want != got {
		t.Errorf("Want %#v got %#v", want, got)
	}

	set.Delete(key)

	want = false
	got = set.Contains(key)

	if want != got {
		t.Errorf("Want %#v got %#v", want, got)
	}
}
