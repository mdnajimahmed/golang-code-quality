package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var CompAsc = func(a, b int) int {
	return a - b
}

func TestBubbleSort(t *testing.T) {
	numbers := []int{3, 2, 1}
	BubbleSort(numbers, CompAsc)
	isAsc := true
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			isAsc = false
			break
		}
	}
	assert.Equal(t, isAsc, true, "Numbers should be sorted in ascending order")
}
