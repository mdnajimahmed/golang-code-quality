package service

import "testing"

func TestMergeSort(t *testing.T) {
	numbers := []int{3, 1, 2, 4, 10, 9}
	MergeSort(numbers, CompAsc)
	isAsc := true
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			isAsc = false
			break
		}
	}
	if !isAsc {
		t.Errorf("Numbers are not sorted in asc order, %v", numbers)
	} else {
		t.Logf("Sorted asc success %v", numbers)
	}
}
