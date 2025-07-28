package segtree

type SegTree[T any] struct {
	n    int
	size int
	op   func(a, b T) T
	e    func() T
	d    []T
}

func NewSegTree[T any](n int, op func(a, b T) T, e func() T) SegTree[T] {
	bitLen := func(x int) int {
		cnt := 0
		for x > 0 {
			x /= 2
			cnt += 1
		}
		return cnt
	}

	size := 1 << bitLen(n-1)

	seg := SegTree[T]{n, size, op, e, make([]T, size*2)}

	for i := 0; i < seg.n; i++ {
		seg.Set(i, seg.e())
	}

	return seg
}

func (s *SegTree[T]) Set(p int, x T) {
	if p < 0 || s.n <= p {
		panic("")
	}
	p += s.size
	s.d[p] = x
	for p > 0 {
		p >>= 1
		s.d[p] = s.op(s.d[p*2], s.d[p*2+1])
	}
}

func (s *SegTree[T]) Get(p int) T {
	if p < 0 || s.n <= p {
		panic("")
	}

	return s.d[p+s.size]
}

func (s *SegTree[T]) Prod(l, r int) T {
	if l < 0 || s.n < l || r < 0 || s.n < r || l > r {
		panic("")
	}
	l += s.size
	r += s.size
	lres, rres := s.e(), s.e()

	for l < r {
		if l&1 == 1 {
			lres = s.op(lres, s.d[l])
			l += 1
		}
		if r&1 == 1 {
			r -= 1
			rres = s.op(s.d[r], rres)
		}
		l >>= 1
		r >>= 1
	}
	return s.op(lres, rres)
}
