package sorting

import (
	"reflect"
	"runtime"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testSortingAlgorithm(t, SelectSort)
}

func TestBubbleSort(t *testing.T) {
	testSortingAlgorithm(t, BubbleSort)
}

func TestInserSort(t *testing.T) {
	testSortingAlgorithm(t, InsertionSort)
}

func testSortingAlgorithm(t *testing.T, sa func(s []int) []int) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{3, 1, 2},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{4, 10, 3, 5, 1},
			expected: []int{1, 3, 4, 5, 10},
		},
	}

	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		result := sa(test.input)
		if !isEqual(result, test.expected) {
			t.Errorf("%s(%v) = %v, wants = %v",
				runtime.FuncForPC(reflect.ValueOf(sa).Pointer()).Name(),
				input,
				result,
				test.expected)
		}
	}
}

func isEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
