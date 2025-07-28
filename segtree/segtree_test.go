package segtree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func addId() int {
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minId() int {
	return math.MaxInt64
}

func TestSegTree_OutofIndex(t *testing.T) {
	n := 10000
	seg := NewSegTree(n, add, addId)

	assert.Panics(t, func() { seg.Get(-1) })
	assert.Panics(t, func() { seg.Get(10000) })

	assert.Panics(t, func() { seg.Set(-1, 100) })
	assert.Panics(t, func() { seg.Set(10000, 100) })

	assert.Panics(t, func() { seg.Prod(-1, -1) })
	assert.Panics(t, func() { seg.Prod(3, 2) })
	assert.Panics(t, func() { seg.Prod(-1, 2) })
	assert.Panics(t, func() { seg.Prod(100, -1) })
}

func TestSegTree_ValidIndex(t *testing.T) {
	n := 10000
	seg := NewSegTree(n, add, addId)

	seg.Set(0, 100)
	seg.Set(2222, 1010)
	seg.Set(9999, 123)

	seg.Get(0)
	seg.Get(2222)
	seg.Get(9999)

	seg.Prod(0, 0)
	seg.Prod(100, 100)
	seg.Prod(10000, 10000)
}

func TestSegTree_Values(t *testing.T) {
	n := 10000
	seg := NewSegTree(n, add, addId)

	seg.Set(0, 100)
	seg.Set(1234, 2000)
	seg.Set(9999, 30000)

	assert.Equal(t, 100, seg.Get(0))
	assert.Equal(t, 2000, seg.Get(1234))
	assert.Equal(t, 30000, seg.Get(9999))

	assert.Equal(t, 32100, seg.Prod(0, 10000))
	assert.Equal(t, 2100, seg.Prod(0, 9999))
	assert.Equal(t, 2100, seg.Prod(0, 1235))
	assert.Equal(t, 100, seg.Prod(0, 1234))
	assert.Equal(t, 32000, seg.Prod(1, 10000))
	assert.Equal(t, 30000, seg.Prod(1235, 10000))
}

func TestSegTree_ValuesMin(t *testing.T) {
	n := 10000
	seg := NewSegTree(n, min, minId)

	seg.Set(0, 100)
	seg.Set(1234, 2000)
	seg.Set(9999, 30000)

	assert.Equal(t, 100, seg.Prod(0, 10000))
	assert.Equal(t, 2000, seg.Prod(1, 10000))
	assert.Equal(t, 30000, seg.Prod(1235, 10000))
	assert.Equal(t, 100, seg.Prod(0, 1234))
}
