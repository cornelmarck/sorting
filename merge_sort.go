package sorting

import "fmt"

// MergeSort is a stable sorting algorithm based on divide-and-conquer.
// This function implements the top-down variant.
// https://en.wikipedia.org/wiki/Merge_sort
func MergeSort[S ~[]E, E any](x S, cmp func(a, b E) int) {
	// Allocate a working array containing a single copy
	buf := make(S, len(x))
	copy(buf, x)
	topDownMergeSort(buf, x, cmp)
}

// mergeSort sorts the data from src and copies it into dst
func topDownMergeSort[S ~[]E, E any](src S, dst S, cmp func(a, b E) int) {
	if len(src) != len(dst) {
		panic(fmt.Sprintf("mergeSort arrays do not have same length: %d != %d", len(src), len(dst)))
	}
	n := len(src)
	if n <= 1 {
		return
	}

	// Sort the two halfs around the mid-point.
	mid := n / 2
	topDownMergeSort(dst[:mid], src[:mid], cmp)
	topDownMergeSort(dst[mid:], src[mid:], cmp)

	// Merge the two halfs in-order, copying from src into dst.
	i := 0
	j := mid
	for k := range n {
		if i < mid && (j >= n || cmp(src[i], src[j]) <= 0) { // src[i] <= src[j]
			dst[k] = src[i]
			i += 1
		} else {
			dst[k] = src[j]
			j += 1
		}
	}
}
