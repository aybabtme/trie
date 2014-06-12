package trie

import (
	"fmt"
	"runtime"
	"testing"
)

func Benchmark_map_SmallStr_4(b *testing.B)    { benchmark_map_SmallStr(b, 4, 128) }
func Benchmark_map_SmallStr_8(b *testing.B)    { benchmark_map_SmallStr(b, 8, 128) }
func Benchmark_map_SmallStr_16(b *testing.B)   { benchmark_map_SmallStr(b, 16, 128) }
func Benchmark_map_SmallStr_32(b *testing.B)   { benchmark_map_SmallStr(b, 32, 128) }
func Benchmark_map_SmallStr_64(b *testing.B)   { benchmark_map_SmallStr(b, 64, 128) }
func Benchmark_map_SmallStr_512(b *testing.B)  { benchmark_map_SmallStr(b, 512, 128) }
func Benchmark_map_SmallStr_1024(b *testing.B) { benchmark_map_SmallStr(b, 1024, 128) }
func Benchmark_map_SmallStr_1M(b *testing.B)   { benchmark_map_SmallStr(b, 1<<20, 128) }

func benchmark_map_SmallStr(b *testing.B, keys, alphaSize int) {
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

func Benchmark_ternaryST_SmallStr_4(b *testing.B)    { benchmark_ternaryST_SmallStr(b, 4, 128) }
func Benchmark_ternaryST_SmallStr_8(b *testing.B)    { benchmark_ternaryST_SmallStr(b, 8, 128) }
func Benchmark_ternaryST_SmallStr_16(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 16, 128) }
func Benchmark_ternaryST_SmallStr_32(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 32, 128) }
func Benchmark_ternaryST_SmallStr_64(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 64, 128) }
func Benchmark_ternaryST_SmallStr_512(b *testing.B)  { benchmark_ternaryST_SmallStr(b, 512, 128) }
func Benchmark_ternaryST_SmallStr_1024(b *testing.B) { benchmark_ternaryST_SmallStr(b, 1024, 128) }
func Benchmark_ternaryST_SmallStr_1M(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 1<<20, 128) }

func benchmark_ternaryST_SmallStr(b *testing.B, keys, alphaSize int) {
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

func Benchmark_trie128_SmallStr_4(b *testing.B)    { benchmark_trie_SmallStr(b, 4, 128) }
func Benchmark_trie128_SmallStr_8(b *testing.B)    { benchmark_trie_SmallStr(b, 8, 128) }
func Benchmark_trie128_SmallStr_16(b *testing.B)   { benchmark_trie_SmallStr(b, 16, 128) }
func Benchmark_trie128_SmallStr_32(b *testing.B)   { benchmark_trie_SmallStr(b, 32, 128) }
func Benchmark_trie128_SmallStr_64(b *testing.B)   { benchmark_trie_SmallStr(b, 64, 128) }
func Benchmark_trie128_SmallStr_512(b *testing.B)  { benchmark_trie_SmallStr(b, 512, 128) }
func Benchmark_trie128_SmallStr_1024(b *testing.B) { benchmark_trie_SmallStr(b, 1024, 128) }
func Benchmark_trie128_SmallStr_1M(b *testing.B)   { benchmark_trie_SmallStr(b, 1<<20, 256) }

func benchmark_trie_SmallStr(b *testing.B, keys int, alphaSize int) {
	runtime.GC()
	b.ReportAllocs()
	trie := NewTrie(alphaSize)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
