package dijkstra

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra_Heap(t *testing.T) {
	const N = 4
	q := &Heap{[2]int{0, 3}, [2]int{1, 2}}
	heap.Init(q)
	heap.Push(q, [2]int{4, 1})
	heap.Push(q, [2]int{5, 10})
	fmt.Println(q)

	want := [N][2]int{{4, 1}, {1, 2}, {0, 3}, {5, 10}}

	for i := 0; i < N; i++ {
		top := heap.Pop(q).([2]int)
		assert.Equal(t, want[i], top)
	}
}

func TestDijkstra_SimpleDist(t *testing.T) {
	const N = 5
	g := make([][]Edge, N)
	g[0] = append(g[0], Edge{1, 3})
	g[1] = append(g[1], Edge{2, 4})
	g[1] = append(g[1], Edge{3, 4})
	g[2] = append(g[2], Edge{3, 5})
	g[3] = append(g[3], Edge{4, 6})

	want := [N]int{0, 3, 7, 7, 13}

	dist := Dijkstra(g, N, 0)

	for i := 0; i < N; i++ {
		assert.Equal(t, want[i], dist[i])
	}
}
