package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.Equal(t, isAsc, true, "Numbers should be sorted in ascending order")
}
