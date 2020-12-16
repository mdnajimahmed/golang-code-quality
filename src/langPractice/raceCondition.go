package langPractice

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func raceCondition() {
	testMutex()
	testAtomic()
}

func testAtomic() {
	var v int32
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&v, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("testAtomic", v)
}

func testMutex() {
	v := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("testMutex", v)
}
