package bsearch_benchmark

import (
	"math/rand"
	"testing"
)

var size int = 100000
var tree *Tree = BuildTree(size)
var blackhole bool

func BenchmarkOneComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := BSearchOneComp(tree, rand.Intn(size))
		blackhole = result
	}
}

func BenchmarkTwoComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := BSearchTwoComp(tree, rand.Intn(size))
		blackhole = result
	}
}

func TestKeyPresent(t *testing.T) {
	tree := &Tree{0, nil, nil}
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		Insert(tree, i)
	}

	data := []struct {
		name     string
		target   int
		expected bool
	}{
		{"in tree", 2, true},
		{"too big", 11, false},
		{"too small", -5, false},
		{"not in middle", 5, false},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := BSearchOneComp(tree, d.target)
			if result != d.expected {
				t.Errorf("OneComp expected %v, got %v", d.expected, result)
			}
		})
		t.Run(d.name, func(t *testing.T) {
			result := BSearchTwoComp(tree, d.target)
			if result != d.expected {
				t.Errorf("TwoComp expected %v, got %v", d.expected, result)
			}
		})
	}

	resultOne := BSearchOneComp(tree, 2)
	resultTwo := BSearchTwoComp(tree, 2)

	expectedOne := true
	expectedTwo := true

	if expectedOne != resultOne {
		t.Error("OneComp expected true, got", expectedOne)
	}
	if expectedTwo != resultTwo {
		t.Error("TwoComp expected true, got", expectedTwo)
	}
}
