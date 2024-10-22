// needs to start with a package name, it is special because it states that the file bleongs to an executable
package main

// import statement
import (
	"fmt"
	"slices"
)

const x int64 = 255

const (
	id   string = "radnom"
	pass string = "hello"
)

func main() {
	fmt.Println("Hello From Kannav!!")

	// data types are the same, so string,int,floats and the other common ones

	var task1 string = "something random"

	// declaring multiple variables at once

	var (
		mssg1 string = "hello"

		randomInt int = 10

		randomByte byte = 244

		random int8 = 0
	)
	// this is not legal outside of function definitions
	x := 10

	fmt.Println(task1)
	// shorthand for initialising and declaring a variable
	age := 10

	fmt.Println(age)

	// way to declare vars is

	// var name string = "que pasa"

	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// arr2 := [5]int{1, 2, 3, 5, 4}

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

	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

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

	var x0 [3]int
	var x1 = [3]int{1, 2, 3}
	var x2 [2][3]int
	fmt.Println(len(x0))

	// slices

	var slice1 = []int{1, 2, 3}
	var slice2 []int
	var slice3 [][]int
	// can be used for comparing slices
	slices.Equal(slice1, slice2)

	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20, 30, 40)
	slice2 = append(slice2, slice1...)

	// length of 5 and capacity of 10
	makeSlice := make([]int, 5, 10)

	x10 := []int{1, 2, 3, 4}
	tempSlice := make([]int, 0, 10)
	lenOfTempSlice := copy(tempSlice, x10)
}

func add(a int, b int) int {
	return a + b
}

//  unbuffered channles reuqire both the reciever and the sender to communicatie simultaneously, so it blocks after one input
//  buffered allows async as well and blocks if the reciever never recieves anything or if the sender tries to put more messages than required
