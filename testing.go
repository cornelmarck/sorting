package sorting

import (
	"math/rand"
)

// CmpInt will be inlined by the compiler, so performance is the same as the non-generic version.
// You can verify this by running `go test -gcflags="-m" .`.
// This will show a message like "can inline cmpInt"
func CmpInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

type SortFunc func(data []int, cmp func(a, b int) int)

var SortFuncs = map[string]SortFunc{
	"insertion_sort": InsertionSort[[]int],
	"merge_sort":     MergeSort[[]int],
	"quicksort":      QuickSort[[]int],
}

type TestCase struct {
	Name     string
	TestFunc func(n int) []int
}

var TestCases = []TestCase{
	{
		Name: "all elements equal",
		TestFunc: func(n int) []int {
			return make([]int, n)
		},
	},
	{
		Name: "sorted elements",
		TestFunc: func(n int) []int {
			x := make([]int, n)
			for i := 0; i < n; i++ {
				x[i] = i
			}
			return x
		},
	},
	{
		Name: "reverse sorted elements",
		TestFunc: func(n int) []int {
			x := make([]int, n)
			for i := 0; i < n; i++ {
				x[i] = n - i - 1
			}
			return x
		},
	},
	{
		Name: "random elements",
		TestFunc: func(n int) []int {
			x := make([]int, n)
			for i := 0; i < n; i++ {
				x[i] = rand.Intn(50)
			}
			return x
		},
	},
}
