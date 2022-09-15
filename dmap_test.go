package dmap_test

import (
	"fmt"
	"testing"

	"github.com/cool-develope/deterministic-go-map/dmap"
)

func TestDMap(t *testing.T) {
	m := dmap.NewMap[string, int]
	m.Set("haha", 4)
	m.Set("hello", 5)

	v, found := m.Get("haha")
	fmt.Print(v, found)
	fmt.Print(m.Range())
}
