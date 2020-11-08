package service

import (
	"fmt"
	"github.com/mandatorySuicide/golang-code-quality/src/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var CompAsc = func(a, b int) int {
	return a - b
}

func TestMain(m *testing.M) {
	fmt.Print("WOW")
	os.Exit(m.Run())
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

func TestBubbleSortTimer(t *testing.T) {
	numbers := util.MakeRandInt(1000)
	timeoutChannel := make(chan bool, 1)
	defer close(timeoutChannel)
	millisecond := 700

	go func() {
		BubbleSort(numbers, CompAsc)
		timeoutChannel <- false
	}()

	go func() {
		time.Sleep(time.Duration(millisecond) * time.Millisecond)
		timeoutChannel <- true
	}()

	if <-timeoutChannel {
		assert.Fail(t, fmt.Sprintf("Bubble sort took more than %v ms", millisecond))
	}
}
