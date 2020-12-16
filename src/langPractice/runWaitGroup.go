package langPractice

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func runWaitGroup() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("OS\t", runtime.GOARCH)
	fmt.Println("OS\t", runtime.NumCPU())
	fmt.Println("OS\t", runtime.NumGoroutine())

	// right before go func(), we are telling that we are gonna launch one go routine
	wg.Add(1)
	go print1()
	print2()
	fmt.Println("OS\t", runtime.NumGoroutine())
	wg.Wait()

}

func print2() {
	fmt.Println(2)
}

func print1() {
	fmt.Println(1)
	wg.Done()
}
