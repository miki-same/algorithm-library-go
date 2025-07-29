package dijkstra

import (
	"container/heap"
	"math"
)

const INF = math.MaxInt64 / 2

type Edge struct {
	to   int
	cost int
}

type Graph [][]Edge

type Heap [][2]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h *Heap) Push(e interface{}) {
	*h = append(*h, e.([2]int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func Dijkstra(g Graph, n int, s int) []int {
	if len(g) != n {
		panic("")
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	que := &Heap{[2]int{s, 0}}
	heap.Init(que)

	for len(*que) > 0 {
		state := heap.Pop(que).([2]int)
		v, d := state[0], state[1]
		if d > dist[v] {
			continue
		}

		for _, edge := range g[v] {
			if dist[v]+edge.cost < dist[edge.to] {
				dist[edge.to] = dist[v] + edge.cost
				heap.Push(que, [2]int{edge.to, dist[edge.to]})
			}
		}
	}

	return dist
}
