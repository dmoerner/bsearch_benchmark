package bsearch_benchmark

import "math/rand"

type Tree struct {
	v int
	l *Tree
	r *Tree
}

func BSearchOneComp(t *Tree, n int) bool {
	var candidate *Tree
	for t != nil {
		if n < t.v {
			t = t.l
		} else {
			candidate = t
			t = t.r
		}
	}
	if candidate != nil && candidate.v == n {
		return true
	}

	return false
}

func BSearchTwoComp(t *Tree, n int) bool {
	for t != nil {
		if n < t.v {
			t = t.l
		} else if n > t.v {
			t = t.r
		} else {
			return true
		}

	}
	return false
}

func Insert(t *Tree, n int) {
	for t != nil {
		if n < t.v {
			if t.l == nil {
				t.l = &Tree{n, nil, nil}
				return
			}
			t = t.l
		} else if n > t.v {
			if t.r == nil {
				t.r = &Tree{n, nil, nil}
				return
			}
			t = t.r
		} else {
			return
		}
	}

	panic("can't insert into nil tree")

}

func BuildTree(n int) *Tree {
	t := &Tree{rand.Intn(n), nil, nil}
	for i := 0; i < n; i++ {
		Insert(t, rand.Intn(n))

	}
	return t
}
