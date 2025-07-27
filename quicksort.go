package sorting

import "math/rand"

// Quicksort is an efficient, unstable sorting algorithm.
// Pivot selection is randomized, so the worst case is avoided.
func QuickSort[S ~[]E, E any](x S, cmp func(a, b E) int) {
	if len(x) <= 1 {
		return
	}
	pivotIdx := partition(x, cmp, true)
	QuickSort(x[:pivotIdx], cmp)
	QuickSort(x[pivotIdx+1:], cmp)
}

func partition[S ~[]E, E any](x S, cmp func(a, b E) int, random bool) int {
	pivot := len(x) - 1
	if random {
		randomIdx := rand.Intn(len(x))
		x[randomIdx], x[pivot] = x[pivot], x[randomIdx]
	}

	i := 0
	for j := 0; j < len(x)-1; j++ {
		if cmp(x[j], x[len(x)-1]) <= 0 {
			x[i], x[j] = x[j], x[i]
			i += 1
		}
	}

	x[i], x[pivot] = x[pivot], x[i]
	return i
}

// QuickSortNaive uses the Lomuto partitioning scheme. The pivot is always the
// last item in the subarray so it is non-random. This algorithm degrades to
// O(N^2) time complexity for the worst case of an array that is already (nearly)
// sorted.
func QuickSortNaive[S ~[]E, E any](x S, cmp func(a, b E) int) {
	if len(x) <= 1 {
		return
	}
	pivotIdx := partition(x, cmp, false)
	QuickSortNaive(x[:pivotIdx], cmp)
	QuickSortNaive(x[pivotIdx+1:], cmp)
}
