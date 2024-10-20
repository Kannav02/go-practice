// needs to start with a package name, it is special because it states that the file bleongs to an executable
package main

// import statement
import "fmt"

func main() {
	fmt.Println("Hello From Kannav!!")

	// data types are the same, so string,int,floats and the other common ones

	var task1 string = "something random"

	fmt.Println(task1)
	// shorthand for initialising and declaring a variable
	age := 10

	fmt.Println(age)

	// way to declare vars is

	var name string = "que pasa"

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 5, 4}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// this is used to make a sepearete thread or a go-routine
	go add(3, 5)

	// to make a channel, channels are used to share data between routines
	//  here we are making a channel that would store string data
	c := make(chan string)

	go func() {
		c <- "Hello from go-routine"
	}()

	message := <-c
	fmt.Println(message)

	c1 := make(chan string,2)
    c2 := make(chan string,2)

    go func() {
        c1 <- "message from c1"
    }()

    go func() {
        c2 <- "message from c2"
    }()

    select {
    case msg1 := <-c1:
        fmt.Println("Received:", msg1)
    case msg2 := <-c2:
        fmt.Println("Received:", msg2)
    }

}

func add(a int, b int) int {
	return a + b
}


//  unbuffered channles reuqire both the reciever and the sender to communicatie simultaneously, so it blocks after one input
//  buffered allows async as well and blocks if the reciever never recieves anything or if the sender tries to put more messages than required
