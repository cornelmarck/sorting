package sorting

import "testing"

func TestSort(t *testing.T) {
	n := int(100)

	for name, sortFunc := range SortFuncs {
		for _, tc := range TestCases {
			name := name + "_" + tc.Name
			data := tc.TestFunc(n)
			t.Run(name, func(t *testing.T) {
				sortFunc(data, CmpInt)

				// Check if the data is sorted
				for i := 1; i < len(data); i++ {
					if data[i-1] > data[i] {
						t.Errorf("Data not sorted: %v", data)
						return
					}
				}
			})
		}
	}
}
