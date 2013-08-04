package trie

import (
	"fmt"
	"runtime"
	"testing"
)

func Benchmark_map_SmallStr_0(b *testing.B)    { benchmark_map_SmallStr(b, 0, false) }
func Benchmark_map_SmallStr_4(b *testing.B)    { benchmark_map_SmallStr(b, 4, false) }
func Benchmark_map_SmallStr_8(b *testing.B)    { benchmark_map_SmallStr(b, 8, false) }
func Benchmark_map_SmallStr_16(b *testing.B)   { benchmark_map_SmallStr(b, 16, false) }
func Benchmark_map_SmallStr_32(b *testing.B)   { benchmark_map_SmallStr(b, 32, false) }
func Benchmark_map_SmallStr_64(b *testing.B)   { benchmark_map_SmallStr(b, 64, false) }
func Benchmark_map_SmallStr_512(b *testing.B)  { benchmark_map_SmallStr(b, 512, false) }
func Benchmark_map_SmallStr_1024(b *testing.B) { benchmark_map_SmallStr(b, 1024, false) }
func Benchmark_map_SmallStr_1M(b *testing.B)   { benchmark_map_SmallStr(b, 1<<20, false) }

func benchmark_map_SmallStr(b *testing.B, keys int, two bool) {
	runtime.GC()
	m := make(map[string]bool)
	for i := 0; i < keys; i++ {
		m[fmt.Sprint(i)] = true
	}
	b.ResetTimer()
	key := fmt.Sprint(keys + 1)
	for i := 0; i < b.N; i++ {
		if two {
			_, _ = m[key]
		} else {
			_ = m[key]
		}
	}
}

func Benchmark_ternaryST_SmallStr_0(b *testing.B)    { benchmark_ternaryST_SmallStr(b, 0) }
func Benchmark_ternaryST_SmallStr_4(b *testing.B)    { benchmark_ternaryST_SmallStr(b, 4) }
func Benchmark_ternaryST_SmallStr_8(b *testing.B)    { benchmark_ternaryST_SmallStr(b, 8) }
func Benchmark_ternaryST_SmallStr_16(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 16) }
func Benchmark_ternaryST_SmallStr_32(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 32) }
func Benchmark_ternaryST_SmallStr_64(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 64) }
func Benchmark_ternaryST_SmallStr_512(b *testing.B)  { benchmark_ternaryST_SmallStr(b, 512) }
func Benchmark_ternaryST_SmallStr_1024(b *testing.B) { benchmark_ternaryST_SmallStr(b, 1024) }
func Benchmark_ternaryST_SmallStr_1M(b *testing.B)   { benchmark_ternaryST_SmallStr(b, 1<<20) }

func benchmark_ternaryST_SmallStr(b *testing.B, keys int) {
	runtime.GC()
	trie := NewTernaryST()
	for i := 0; i < keys; i++ {
		trie.Put(fmt.Sprint(i), true)
	}
	b.ResetTimer()
	key := fmt.Sprint(keys + 1)
	for i := 0; i < b.N; i++ {
		_, _ = trie.Get(key)
	}
}

func Benchmark_trie128_SmallStr_0(b *testing.B)    { benchmark_trie_SmallStr(b, 0, 128) }
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
	trie := NewTrie(alphaSize)
	for i := 0; i < keys; i++ {
		trie.Put(fmt.Sprint(i), true)
	}
	b.ResetTimer()
	key := fmt.Sprint(keys + 1)
	for i := 0; i < b.N; i++ {
		_, _ = trie.Get(key)
	}
}
