package service

func BubbleSort(numbers []int, comp func(int, int) int) {
	size := len(numbers)
	for i := 0; i < size; i++ {
		shouldContinue := false
		for j := 1; j < size; j++ {
			if comp(numbers[j-1], numbers[j]) > 0 {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
				shouldContinue = true
			}
		}
		if !shouldContinue {
			break
		}
	}
}
