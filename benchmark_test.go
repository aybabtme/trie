package trie

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"testing"
)

func Benchmark_map_Lookup_4(b *testing.B)    { benchmark_map_Lookup(b, 4, 64) }
func Benchmark_map_Lookup_8(b *testing.B)    { benchmark_map_Lookup(b, 8, 64) }
func Benchmark_map_Lookup_16(b *testing.B)   { benchmark_map_Lookup(b, 16, 64) }
func Benchmark_map_Lookup_32(b *testing.B)   { benchmark_map_Lookup(b, 32, 64) }
func Benchmark_map_Lookup_64(b *testing.B)   { benchmark_map_Lookup(b, 64, 64) }
func Benchmark_map_Lookup_512(b *testing.B)  { benchmark_map_Lookup(b, 512, 64) }
func Benchmark_map_Lookup_1024(b *testing.B) { benchmark_map_Lookup(b, 1024, 64) }
func Benchmark_map_Lookup_1M(b *testing.B)   { benchmark_map_Lookup(b, 1<<20, 64) }

func benchmark_map_Lookup(b *testing.B, keys int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	m := make(map[string]bool)
	var longest string
	for i := 0; i < keys; i++ {
		str := fmt.Sprint(i)
		m[str] = true
		if len(longest) < len(str) {
			longest = str
		}
	}
	// b.Logf("longest=%q", longest)
	b.ResetTimer()

	allTrue := true
	var resp, ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = m[longest]
		allTrue = allTrue && resp && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func Benchmark_ternaryST_Lookup_4(b *testing.B)    { benchmark_ternaryST_Lookup(b, 4, 64) }
func Benchmark_ternaryST_Lookup_8(b *testing.B)    { benchmark_ternaryST_Lookup(b, 8, 64) }
func Benchmark_ternaryST_Lookup_16(b *testing.B)   { benchmark_ternaryST_Lookup(b, 16, 64) }
func Benchmark_ternaryST_Lookup_32(b *testing.B)   { benchmark_ternaryST_Lookup(b, 32, 64) }
func Benchmark_ternaryST_Lookup_64(b *testing.B)   { benchmark_ternaryST_Lookup(b, 64, 64) }
func Benchmark_ternaryST_Lookup_512(b *testing.B)  { benchmark_ternaryST_Lookup(b, 512, 64) }
func Benchmark_ternaryST_Lookup_1024(b *testing.B) { benchmark_ternaryST_Lookup(b, 1024, 64) }
func Benchmark_ternaryST_Lookup_1M(b *testing.B)   { benchmark_ternaryST_Lookup(b, 1<<20, 64) }

func benchmark_ternaryST_Lookup(b *testing.B, keys int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	trie := NewTernaryST()
	var longest string
	for i := 0; i < keys; i++ {
		str := fmt.Sprint(i)
		trie.Put(str, true)
		if len(longest) < len(str) {
			longest = str
		}
	}
	// b.Logf("longest=%q", longest)
	b.ResetTimer()
	allTrue := true
	var resp interface{}
	var ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = trie.Get(longest)
		allTrue = allTrue && resp.(bool) && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func Benchmark_trie64_Lookup_4(b *testing.B)    { benchmark_trie_Lookup(b, 4, 64) }
func Benchmark_trie64_Lookup_8(b *testing.B)    { benchmark_trie_Lookup(b, 8, 64) }
func Benchmark_trie64_Lookup_16(b *testing.B)   { benchmark_trie_Lookup(b, 16, 64) }
func Benchmark_trie64_Lookup_32(b *testing.B)   { benchmark_trie_Lookup(b, 32, 64) }
func Benchmark_trie64_Lookup_64(b *testing.B)   { benchmark_trie_Lookup(b, 64, 64) }
func Benchmark_trie64_Lookup_512(b *testing.B)  { benchmark_trie_Lookup(b, 512, 64) }
func Benchmark_trie64_Lookup_1024(b *testing.B) { benchmark_trie_Lookup(b, 1024, 64) }
func Benchmark_trie64_Lookup_1M(b *testing.B)   { benchmark_trie_Lookup(b, 1<<20, 64) }

func benchmark_trie_Lookup(b *testing.B, keys int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	trie := NewTrie('0', alphaSize)
	var longest string
	for i := 0; i < keys; i++ {
		str := fmt.Sprint(i)
		trie.Put(str, true)
		if len(longest) < len(str) {
			longest = str
		}
	}
	// b.Logf("longest=%q", longest)
	b.ResetTimer()
	allTrue := true
	var resp interface{}
	var ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = trie.Get(longest)
		allTrue = allTrue && resp.(bool) && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func Benchmark_map_commonPfx_2(b *testing.B)  { benchmark_map_commonPfx(b, 100000, 100, 2) }
func Benchmark_map_commonPfx_4(b *testing.B)  { benchmark_map_commonPfx(b, 100000, 100, 4) }
func Benchmark_map_commonPfx_8(b *testing.B)  { benchmark_map_commonPfx(b, 100000, 100, 8) }
func Benchmark_map_commonPfx_16(b *testing.B) { benchmark_map_commonPfx(b, 100000, 100, 16) }
func Benchmark_map_commonPfx_32(b *testing.B) { benchmark_map_commonPfx(b, 100000, 100, 32) }
func Benchmark_map_commonPfx_64(b *testing.B) { benchmark_map_commonPfx(b, 100000, 100, 64) }

func benchmark_map_commonPfx(b *testing.B, keys, strlen int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	m := make(map[string]bool)
	var longest string
	for _, key := range genAlphaWithPrefixes(keys, strlen, alphaSize) {
		m[key] = true
		if len(longest) < len(key) {
			longest = key
		}
	}
	b.ResetTimer()

	allTrue := true
	var resp, ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = m[longest]
		allTrue = allTrue && resp && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func Benchmark_ternaryST_commonPfx_2(b *testing.B)  { benchmark_ternaryST_commonPfx(b, 100000, 100, 2) }
func Benchmark_ternaryST_commonPfx_4(b *testing.B)  { benchmark_ternaryST_commonPfx(b, 100000, 100, 4) }
func Benchmark_ternaryST_commonPfx_8(b *testing.B)  { benchmark_ternaryST_commonPfx(b, 100000, 100, 8) }
func Benchmark_ternaryST_commonPfx_16(b *testing.B) { benchmark_ternaryST_commonPfx(b, 100000, 100, 16) }
func Benchmark_ternaryST_commonPfx_32(b *testing.B) { benchmark_ternaryST_commonPfx(b, 100000, 100, 32) }
func Benchmark_ternaryST_commonPfx_64(b *testing.B) { benchmark_ternaryST_commonPfx(b, 100000, 100, 64) }

func benchmark_ternaryST_commonPfx(b *testing.B, keys, strlen int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	trie := NewTernaryST()
	var longest string
	for _, key := range genAlphaWithPrefixes(keys, strlen, alphaSize) {
		trie.Put(key, true)
		if len(longest) < len(key) {
			longest = key
		}
	}
	b.ResetTimer()
	allTrue := true
	var resp interface{}
	var ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = trie.Get(longest)
		allTrue = allTrue && resp.(bool) && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func Benchmark_trie128_commonPfx_2(b *testing.B)  { benchmark_trie_commonPfx(b, 100000, 100, 2) }
func Benchmark_trie128_commonPfx_4(b *testing.B)  { benchmark_trie_commonPfx(b, 100000, 100, 4) }
func Benchmark_trie128_commonPfx_8(b *testing.B)  { benchmark_trie_commonPfx(b, 100000, 100, 8) }
func Benchmark_trie128_commonPfx_16(b *testing.B) { benchmark_trie_commonPfx(b, 100000, 100, 16) }
func Benchmark_trie128_commonPfx_32(b *testing.B) { benchmark_trie_commonPfx(b, 100000, 100, 32) }
func Benchmark_trie128_commonPfx_64(b *testing.B) { benchmark_trie_commonPfx(b, 100000, 100, 64) }

func benchmark_trie_commonPfx(b *testing.B, keys, strlen int, alphaSize uint8) {
	runtime.GC()
	b.ReportAllocs()
	trie := NewTrie('a', alphaSize)
	var longest string
	for _, key := range genAlphaWithPrefixes(keys, strlen, alphaSize) {
		trie.Put(key, true)
		if len(longest) < len(key) {
			longest = key
		}
	}
	b.ResetTimer()
	allTrue := true
	var resp interface{}
	var ok bool
	for i := 0; i < b.N; i++ {
		resp, ok = trie.Get(longest)
		allTrue = allTrue && resp.(bool) && ok
	}
	if !allTrue {
		b.Fatal("allTrue should always be true")
	}
}

func genAlphaWithPrefixes(maxkey, strlen int, alphaSize uint8) []string {
	// for each letter in str, alphaSize choices of letter
	// num key = alpha^(strlen)
	keycount := math.Pow(float64(alphaSize), float64(strlen))
	if float64(maxkey) >= keycount {
		panic(fmt.Sprintf("can't generate %d keys, can only do %.g", maxkey, keycount))
	}

	keys := make([]string, maxkey)

	alphabet := make([]rune, alphaSize)
	for i := rune(0); i < rune(alphaSize); i++ {
		alphabet[i] = 'a' + rune(i)

	}

	buf := bytes.NewBuffer(nil)
	for i := range keys {
		for j := 0; j < strlen; j++ {
			buf.WriteRune(alphabet[rand.Intn(len(alphabet))])
		}
		keys[i] = buf.String()
		buf.Reset()
	}

	return keys
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
