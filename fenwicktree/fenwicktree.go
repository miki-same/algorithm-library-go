package fenwicktree

type FenwickTree struct {
	n int
	d []int
}

func NewFenwickTree(n int) FenwickTree {
	fw := FenwickTree{n: n, d: make([]int, n+1)}
	return fw
}

func (f *FenwickTree) Add(p, x int) {
	if p < 0 || f.n <= p {
		panic("")
	}
	for i := p + 1; i <= f.n; i += f.lsb(i) {
		f.d[i] += x
	}
}

func (f *FenwickTree) Sum(l, r int) int {
	if l < 0 || f.n < l || r < 0 || f.n < r || l > r {
		panic("")
	}

	return f.prefixSum(r) - f.prefixSum(l)
}

func (f *FenwickTree) prefixSum(i int) int {
	if i < 0 || f.n < i {
		panic("")
	}
	res := 0
	for i := i; i > 0; i -= f.lsb(i) {
		res += f.d[i]
	}
	return res
}

func (f *FenwickTree) lsb(x int) int {
	return x & (-x)
}
