package fenwicktree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFenwickTree_OutofIndex(t *testing.T) {
	n := 10000
	fw := NewFenwickTree(n)

	assert.Panics(t, func() { fw.Add(-1, 100) })
	assert.Panics(t, func() { fw.Add(10000, 100) })

	assert.Panics(t, func() { fw.Sum(-1, -1) })
	assert.Panics(t, func() { fw.Sum(3, 2) })
	assert.Panics(t, func() { fw.Sum(-1, 2) })
	assert.Panics(t, func() { fw.Sum(100, -1) })
}

func TestFenwickTree_ValidIndex(t *testing.T) {
	n := 10000
	fw := NewFenwickTree(n)

	fw.Add(0, 100)
	fw.Add(2222, 1010)
	fw.Add(9999, 123)

	fw.Sum(0, 0)
	fw.Sum(100, 100)
	fw.Sum(10000, 10000)

	fw.Sum(0, 10000)
	fw.Sum(100, 10000)
	fw.Sum(0, 100)
	fw.Sum(100, 1000)
}

func TestFenwickTree_TestAdd(t *testing.T) {
	n := 4
	fw := NewFenwickTree(n)

	fw.Add(0, 1)
	fw.Add(1, 10)
	fw.Add(2, 100)
	fw.Add(3, 1000)

	want := [5]int{0, 1, 11, 100, 1111}

	for i := 0; i <= n; i++ {
		assert.Equal(t, want[i], fw.d[i])
	}
}

func TestFenwickTree_TestPrefixSum(t *testing.T) {
	n := 10000
	fw := NewFenwickTree(n)

	fw.Add(0, 100)
	fw.Add(1000, 2000)
	fw.Add(9999, 30000)

	assert.Equal(t, 32100, fw.Sum(0, 10000))
	assert.Equal(t, 100, fw.Sum(0, 1000))
	assert.Equal(t, 2100, fw.Sum(0, 9999))
	assert.Equal(t, 32000, fw.Sum(1000, 10000))
	assert.Equal(t, 30000, fw.Sum(1001, 10000))
}

func TestFenwickTree_TestLsb(t *testing.T) {
	n := 10000
	fw := NewFenwickTree(n)

	assert.Equal(t, 1, fw.lsb(3))
	assert.Equal(t, 8, fw.lsb(8))
	assert.Equal(t, 1, fw.lsb(15))
	assert.Equal(t, 4, fw.lsb(28))
}
