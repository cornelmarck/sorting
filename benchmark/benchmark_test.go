package benchmark

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/cornelmarck/sorting"
)

func randomIntFunc(n int) []int {
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = rand.Intn(n)
	}
	return x
}

func BenchmarkSort(b *testing.B) {
	minValue := 8
	numPoints := 16

	for algoName, sortFunc := range sorting.SortFuncs {
		for i := range numPoints {
			size := minValue * int(math.Pow(2., float64(i)))
			label := algoName + "-" + strconv.Itoa(size)
			data := randomIntFunc(size)
			b.Run(label, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sortFunc(data, sorting.CmpInt)
				}
			})
		}
	}
}
