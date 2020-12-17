package langPractice

import "fmt"

func channelWork() {
	way1()
	way2()
	twoBuffer()
	directionalChannel()
	channelRange()
	selectOnChannel()
}

func selectOnChannel() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	// send value to channel
	go send(even, odd, quit)
	receive(even, odd, quit)
	channelWorkGotcha()
}

func receive(even, odd, quit <-chan int) {
	for {
		fmt.Println("FORING")
		select {
		case v := <-even:
			fmt.Printf("Received value from even channel %v\n", v)
		case v := <-odd:
			fmt.Printf("Received value from odd channel %v\n", v)
		case <-quit:
			fmt.Println("Time to quit")
			return
		}

	}

}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if (i & 1) == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- 0
}

func channelRange() {
	c := make(chan int)
	go pubMany(c)
	for i := range c {
		fmt.Println(i)
	}
}

func pubMany(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // if we don't close the program will fall in deadlock because range waits for channel until its closed.
}

func directionalChannel() {
	c := make(chan int)
	go pub(c)
	sub(c)
}

func sub(c <-chan int) {
	fmt.Println(<-c)
}

func pub(c chan<- int) {
	c <- 55
}

func twoBuffer() {
	c := make(chan int, 2)
	c <- 10
	c <- 11
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func way2() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func way1() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
