package langPractice

import "fmt"

// when we close a channel it puts a default value in the channel, ,lets this hypothesis
// with int, string, and bool

// conclusion : better use comma,ok syntax to avoid using the extra sentinal value put to channel.
func channelWorkGotcha() {
	fmt.Println("TESTING CHANNEL GOTCHA")
	testWithInt()
	testWithBool()
	testWithString()
}

func testWithString() {
	fmt.Println("TEST WITH STRING")
	intChan := make(chan string)
	go func() {
		intChan <- "Hello!"
		close(intChan)
	}()
	data1, ok1 := <-intChan
	fmt.Println("["+data1+"]", ok1)
	data2, ok2 := <-intChan
	fmt.Println("["+data2+"]", ok2) // default empty string!
}

func testWithBool() {
	fmt.Println("TEST WITH BOOL")

	intChan := make(chan bool)
	go func() {
		intChan <- true
		close(intChan)
	}()
	data1, ok1 := <-intChan
	fmt.Println(data1, ok1)
	data2, ok2 := <-intChan
	fmt.Println(data2, ok2)

}

func testWithInt() {
	fmt.Println("TEST WITH INT")
	intChan := make(chan int)
	go func() {
		intChan <- 500
		close(intChan)
	}()
	data1, ok1 := <-intChan
	// we put 500 , we receive 500 with ok1=true
	fmt.Println(data1, ok1)
	// but then why do we receive 0 with ok1=false
	data2, ok2 := <-intChan
	fmt.Println(data2, ok2)
	// because channel put default value of int which is 0, then how do we distinguish?
	// well we use the ok flag, if ok is true, we put that value in our channel, else go put it!
}
