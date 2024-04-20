package sorting

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	cases := map[string]struct {
		arr    Array
		result Array
	}{
		"empty array": {arr: Array{}, result: Array{}},
		"array-1": {
			arr:    Array{3, 2, 19, 0, 23, 4},
			result: Array{0, 2, 3, 4, 19, 23},
		},
		"array-2": {
			arr:    Array{18, 0, 3, 6},
			result: Array{0, 3, 6, 18},
		},
		"array-3": {
			arr:    Array{121, 2, 43, 1, 73, 7, 8, 23, 98},
			result: Array{1, 2, 7, 8, 23, 43, 73, 98, 121},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			sarr := Quicksort(tc.arr)

			if !reflect.DeepEqual(sarr, tc.result) {
				t.Errorf("result mismatch in test %s: want %v, got %v", name, tc.result, sarr)
			}

		})
	}

}
