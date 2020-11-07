package service

func MergeSort(numbers []int, comp func(int, int) int) {
	size := len(numbers)
	recursiveSort(numbers, 0, size-1, comp)
}

func recursiveSort(numbers []int, st, ed int, comp func(int, int) int) {
	// if we have more than one item in the range
	if ed-st+1 > 1 {
		mid := (st + ed) >> 1
		recursiveSort(numbers, st, mid, comp)
		recursiveSort(numbers, mid+1, ed, comp)
		merge(numbers, st, mid, ed, comp)
	}
}

func merge(numbers []int, st, mid, ed int, comp func(int, int) int) {
	tmp := make([]int, ed-st+1)
	pos := st
	idx := 0
	i, j := st, mid+1
	for i <= mid && j <= ed {
		if comp(numbers[i], numbers[j]) <= 0 {
			tmp[idx] = numbers[i]
			i++
			idx++
		} else {
			tmp[idx] = numbers[j]
			j++
			idx++
		}
	}
	for i <= mid {
		tmp[idx] = numbers[i]
		i++
		idx++
	}
	for j <= ed {
		tmp[idx] = numbers[j]
		j++
		idx++
	}
	for i := 0; i < len(tmp); i++ {
		numbers[pos+i] = tmp[i]
	}

}
