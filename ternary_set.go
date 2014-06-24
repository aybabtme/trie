package trie

type ternSetNode struct {
	exists bool
	Code   rune

	left  *ternSetNode
	child *ternSetNode
	right *ternSetNode
}

// TernarySet is a set specifically for string indexed keys.
type TernarySet struct {
	root  *ternSetNode
	count int
}

var (
	TernarySetIsList ListSet = NewTernarySet()
)

// NewTernarySet creates a trie.
func NewTernarySet() *TernarySet {
	return &TernarySet{nil, 0}
}

// Add the key to the set.
func (t *TernarySet) Add(key string) {

	t.root = t.put(t.root, key, 0)
}

func (t *TernarySet) put(x *ternSetNode, key string, d int) *ternSetNode {
	c := key[d]
	if x == nil {
		x = &ternSetNode{Code: rune(c)}
	}

	if c < uint8(x.Code) {
		x.left = t.put(x.left, key, d)
	} else if c > uint8(x.Code) {
		x.right = t.put(x.right, key, d)
	} else if d < len(key)-1 {
		x.child = t.put(x.child, key, d+1)
	} else {
		x.exists = true

	}
	return x
}

// Contains tells if key exists.
func (t *TernarySet) Contains(key string) bool {

	if key == "" {
		return false
	}

	ternSetNode := t.get(t.root, key, 0)

	if ternSetNode == nil {
		return false
	}

	return ternSetNode.exists
}

func (t *TernarySet) get(x *ternSetNode, key string, d int) *ternSetNode {
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

// Len returns the count of elements in this set.
func (t *TernarySet) Len() int { return t.count }

// IsEmpty tells if this set contains any elements.
func (t *TernarySet) IsEmpty() bool { return t.Len() == 0 }

// Keys returns all the keys known to this trie
func (t *TernarySet) Keys() []string {
	var outCollection []string
	collectSetKeys(t.root, []uint8{}, outCollection)
	return outCollection
}

// KeysWithPrefix returns all the keys starting with prefix `key`
func (t *TernarySet) KeysWithPrefix(key string) []string {

	var outCollection []string
	x := t.get(t.root, key, 0)
	if x == nil {
		return outCollection
	}
	if x.exists {
		outCollection = append(outCollection, key)
	}
	collectSetKeys(x.child, []byte(key), outCollection)
	return outCollection
}

// LongestPrefix returns the longest string in this trie that has `key` for prefix
func (t *TernarySet) LongestPrefix(key string) string {
	if key == "" {
		return ""
	}

	lenght := 0

	x := t.root
	i := 0
	var c uint8
	for x != nil && i < len(key) {
		c = key[i]
		if c < uint8(x.Code) {
			x = x.left
		} else if c > uint8(x.Code) {
			x = x.right
		} else {
			i++
			if x.exists {
				lenght = i
			}
			x = x.child
		}
	}
	return string(key[0:lenght])
}

// KeysMatching returns all the keys that share prefix `key`, where `key` can
// contain a wildcard character `.`.
func (t *TernarySet) KeysMatching(key string) []string {
	var outCollection []string
	patternCollectSetKeys(t.root, []uint8{}, 0, []byte(key), outCollection)
	return outCollection
}

// Helpers

// `collectSetKeys` collectSetKeyss a set of string that share a prefix
func collectSetKeys(x *ternSetNode, key []byte, outCollection []string) {
	if x == nil {
		return
	}
	collectSetKeys(x.left, key, outCollection)
	newKey := append(key, uint8(x.Code))
	if x.exists {
		outCollection = append(outCollection, string(newKey))
	}
	collectSetKeys(x.child, newKey, outCollection)
	collectSetKeys(x.right, key, outCollection)
}

// 'patternCollectSetKeys' collectSetKeyss a set of string that matches the pattern
func patternCollectSetKeys(
	x *ternSetNode,
	prefix []uint8,
	i int,
	pat []byte,
	outCollection []string,
) {
	if x == nil {
		return
	}

	c := pat[i]

	if c == uint8('.') || c < uint8(x.Code) {
		patternCollectSetKeys(x.left, prefix, i, pat, outCollection)
	}

	if c == uint8('.') || c == uint8(x.Code) {
		newKey := append(prefix, c)
		if i == len(pat)-1 && x.exists {
			outCollection = append(outCollection, string(newKey))
		}
		if i < len(pat)-1 {
			patternCollectSetKeys(x.child, newKey, i+1, pat, outCollection)
		}
	}

	if c == uint8('.') || c > uint8(x.Code) {
		patternCollectSetKeys(x.right, prefix, i, pat, outCollection)
	}
}
