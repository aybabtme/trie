package trie

import (
	"fmt"
	"runtime"
	// "strings"
	"testing"
	"unicode/utf8"
)

// func Benchmark_map_BigStr_0(b *testing.B)    { benchmark_map_BigStr(b, 0, false) }
// func Benchmark_map_BigStr_4(b *testing.B)    { benchmark_map_BigStr(b, 4, false) }
// func Benchmark_map_BigStr_8(b *testing.B)    { benchmark_map_BigStr(b, 8, false) }
// func Benchmark_map_BigStr_16(b *testing.B)   { benchmark_map_BigStr(b, 16, false) }
// func Benchmark_map_BigStr_32(b *testing.B)   { benchmark_map_BigStr(b, 32, false) }
// func Benchmark_map_BigStr_64(b *testing.B)   { benchmark_map_BigStr(b, 64, false) }
// func Benchmark_map_BigStr_512(b *testing.B)  { benchmark_map_BigStr(b, 512, false) }
// func Benchmark_map_BigStr2_0(b *testing.B)   { benchmark_map_BigStr(b, 0, true) }
// func Benchmark_map_BigStr2_4(b *testing.B)   { benchmark_map_BigStr(b, 4, true) }
// func Benchmark_map_BigStr2_8(b *testing.B)   { benchmark_map_BigStr(b, 8, true) }
// func Benchmark_map_BigStr2_16(b *testing.B)  { benchmark_map_BigStr(b, 16, true) }
// func Benchmark_map_BigStr2_32(b *testing.B)  { benchmark_map_BigStr(b, 32, true) }
// func Benchmark_map_BigStr2_64(b *testing.B)  { benchmark_map_BigStr(b, 64, true) }
// func Benchmark_map_BigStr2_512(b *testing.B) { benchmark_map_BigStr(b, 512, true) }

// func benchmark_map_BigStr(b *testing.B, keys int, two bool) {
//  runtime.GC()
// 	m := make(map[string]bool)
// 	for i := 0; i < keys; i++ {
// 		suffix := fmt.Sprint(i)
// 		key := strings.Repeat("X", 1<<20-len(suffix)) + suffix
// 		m[key] = true
// 	}
// 	key := strings.Repeat("X", 1<<20-1) + "k"
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		if two {
// 			_, _ = m[key]
// 		} else {
// 			_ = m[key]
// 		}
// 	}
// }

func Benchmark_map_SmallStr_0(b *testing.B)     { benchmark_map_SmallStr(b, 0, false) }
func Benchmark_map_SmallStr_4(b *testing.B)     { benchmark_map_SmallStr(b, 4, false) }
func Benchmark_map_SmallStr_8(b *testing.B)     { benchmark_map_SmallStr(b, 8, false) }
func Benchmark_map_SmallStr_16(b *testing.B)    { benchmark_map_SmallStr(b, 16, false) }
func Benchmark_map_SmallStr_32(b *testing.B)    { benchmark_map_SmallStr(b, 32, false) }
func Benchmark_map_SmallStr_64(b *testing.B)    { benchmark_map_SmallStr(b, 64, false) }
func Benchmark_map_SmallStr_512(b *testing.B)   { benchmark_map_SmallStr(b, 512, false) }
func Benchmark_map_SmallStr_1024(b *testing.B)  { benchmark_map_SmallStr(b, 1024, false) }
func Benchmark_map_SmallStr_1M(b *testing.B)    { benchmark_map_SmallStr(b, 1<<20, false) }
func Benchmark_map_SmallStr2_0(b *testing.B)    { benchmark_map_SmallStr(b, 0, true) }
func Benchmark_map_SmallStr2_4(b *testing.B)    { benchmark_map_SmallStr(b, 4, true) }
func Benchmark_map_SmallStr2_8(b *testing.B)    { benchmark_map_SmallStr(b, 8, true) }
func Benchmark_map_SmallStr2_16(b *testing.B)   { benchmark_map_SmallStr(b, 16, true) }
func Benchmark_map_SmallStr2_32(b *testing.B)   { benchmark_map_SmallStr(b, 32, true) }
func Benchmark_map_SmallStr2_64(b *testing.B)   { benchmark_map_SmallStr(b, 64, true) }
func Benchmark_map_SmallStr2_512(b *testing.B)  { benchmark_map_SmallStr(b, 512, true) }
func Benchmark_map_SmallStr2_1024(b *testing.B) { benchmark_map_SmallStr(b, 1024, true) }
func Benchmark_map_SmallStr2_1M(b *testing.B)   { benchmark_map_SmallStr(b, 1<<20, true) }

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

// func Benchmark_trie256_BigStr_0(b *testing.B)   { benchmark_trie_BigStr(b, 0, 256) }
// func Benchmark_trie256_BigStr_4(b *testing.B)   { benchmark_trie_BigStr(b, 4, 256) }
// func Benchmark_trie256_BigStr_8(b *testing.B)   { benchmark_trie_BigStr(b, 8, 256) }
// func Benchmark_trie256_BigStr_16(b *testing.B)  { benchmark_trie_BigStr(b, 16, 256) }
// func Benchmark_trie256_BigStr_32(b *testing.B)  { benchmark_trie_BigStr(b, 32, 256) }
// func Benchmark_trie256_BigStr_64(b *testing.B)  { benchmark_trie_BigStr(b, 64, 256) }
// func Benchmark_trie256_BigStr_512(b *testing.B) { benchmark_trie_BigStr(b, 512, 256) }

// func Benchmark_trieUTF8_BigStr_0(b *testing.B)   { benchmark_trie_BigStr(b, 0, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_4(b *testing.B)   { benchmark_trie_BigStr(b, 4, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_8(b *testing.B)   { benchmark_trie_BigStr(b, 8, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_16(b *testing.B)  { benchmark_trie_BigStr(b, 16, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_32(b *testing.B)  { benchmark_trie_BigStr(b, 32, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_64(b *testing.B)  { benchmark_trie_BigStr(b, 64, utf8.MaxRune) }
// func Benchmark_trieUTF8_BigStr_512(b *testing.B) { benchmark_trie_BigStr(b, 512, utf8.MaxRune) }

// func benchmark_trie_BigStr(b *testing.B, keys int, alphaSize int) {
//  runtime.GC()
// 	trie := NewTrie(alphaSize)
// 	for i := 0; i < keys; i++ {
// 		suffix := fmt.Sprint(i)
// 		key := strings.Repeat("X", 1<<20-len(suffix)) + suffix
// 		trie.Put(key, true)
// 	}
// 	key := strings.Repeat("X", 1<<20-1) + "k"
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, _ = trie.Get(key)
// 	}
// }
func Benchmark_trie128_SmallStr_0(b *testing.B)    { benchmark_trie_SmallStr(b, 0, 128) }
func Benchmark_trie128_SmallStr_4(b *testing.B)    { benchmark_trie_SmallStr(b, 4, 128) }
func Benchmark_trie128_SmallStr_8(b *testing.B)    { benchmark_trie_SmallStr(b, 8, 128) }
func Benchmark_trie128_SmallStr_16(b *testing.B)   { benchmark_trie_SmallStr(b, 16, 128) }
func Benchmark_trie128_SmallStr_32(b *testing.B)   { benchmark_trie_SmallStr(b, 32, 128) }
func Benchmark_trie128_SmallStr_64(b *testing.B)   { benchmark_trie_SmallStr(b, 64, 128) }
func Benchmark_trie128_SmallStr_512(b *testing.B)  { benchmark_trie_SmallStr(b, 512, 128) }
func Benchmark_trie128_SmallStr_1024(b *testing.B) { benchmark_trie_SmallStr(b, 1024, 128) }
func Benchmark_trie128_SmallStr_1M(b *testing.B)   { benchmark_trie_SmallStr(b, 1<<20, 256) }

func Benchmark_trie256_SmallStr_0(b *testing.B)    { benchmark_trie_SmallStr(b, 0, 256) }
func Benchmark_trie256_SmallStr_4(b *testing.B)    { benchmark_trie_SmallStr(b, 4, 256) }
func Benchmark_trie256_SmallStr_8(b *testing.B)    { benchmark_trie_SmallStr(b, 8, 256) }
func Benchmark_trie256_SmallStr_16(b *testing.B)   { benchmark_trie_SmallStr(b, 16, 256) }
func Benchmark_trie256_SmallStr_32(b *testing.B)   { benchmark_trie_SmallStr(b, 32, 256) }
func Benchmark_trie256_SmallStr_64(b *testing.B)   { benchmark_trie_SmallStr(b, 64, 256) }
func Benchmark_trie256_SmallStr_512(b *testing.B)  { benchmark_trie_SmallStr(b, 512, 256) }
func Benchmark_trie256_SmallStr_1024(b *testing.B) { benchmark_trie_SmallStr(b, 1024, 256) }
func Benchmark_trie256_SmallStr_1M(b *testing.B)   { benchmark_trie_SmallStr(b, 1<<20, 256) }

func Benchmark_trieUTF8_SmallStr_0(b *testing.B)    { benchmark_trie_SmallStr(b, 0, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_4(b *testing.B)    { benchmark_trie_SmallStr(b, 4, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_8(b *testing.B)    { benchmark_trie_SmallStr(b, 8, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_16(b *testing.B)   { benchmark_trie_SmallStr(b, 16, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_32(b *testing.B)   { benchmark_trie_SmallStr(b, 32, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_64(b *testing.B)   { benchmark_trie_SmallStr(b, 64, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_512(b *testing.B)  { benchmark_trie_SmallStr(b, 512, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_1024(b *testing.B) { benchmark_trie_SmallStr(b, 1024, utf8.MaxRune) }
func Benchmark_trieUTF8_SmallStr_1M(b *testing.B)   { benchmark_trie_SmallStr(b, 1<<20, utf8.MaxRune) }

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
