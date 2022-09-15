package dmap_test

import (
	"fmt"
	"testing"

	dmap "github.com/cool-develope/deterministic-go-map"
)

func TestDMap(t *testing.T) {
	m := dmap.NewMap[string, int]()
	m.Set("haha", 4)
	m.Set("hello", 5)

	v, found := m.Get("haha")
	t.Log(v, found)
	t.Log(m.Range())
}

const count = 1000000

func BenchmarkDMapSet(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	m := dmap.NewMap[string, int]()
	for i := 0; i < count; i++ {
		m.Set(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkMapSet(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	m := make(map[string]int)
	for i := 0; i < count; i++ {
		m[fmt.Sprintf("key%d", i)] = i
	}
}
