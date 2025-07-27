package sorting

// InsertionSort is a stable sorting algorithm that is efficient for small data sets.
// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort[S ~[]E, E any](x S, cmp func(a, b E) int) {
	for i := 1; i < len(x); i++ {
		j := i
		for j > 0 && cmp(x[j-1], x[j]) > 0 {
			x[j-1], x[j] = x[j], x[j-1]
			j -= 1
		}
	}
}
