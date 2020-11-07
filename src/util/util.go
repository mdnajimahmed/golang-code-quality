package util

import "math/rand"

func MakeRandInt(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100000)
	}
	return numbers
}
