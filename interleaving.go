/*
Explanation:

The race condition occurs because we can not predict the execution order of increment and ispositive functions. Even though we placed ispositiv() function call after the increment() call.
It will print that the number is negative in some cases.

*/

package main

import (
	"fmt"
	"time"
)

func increment(num *int, step int) {
	*num += step
}

func ispositive(num int) {
	if num >= 0 {
		fmt.Println("The num is positive")
	} else {
		fmt.Println("The num is negative")
	}
}

func main() {
	x := -1
	go increment(&x, 2)
	go ispositive(x)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Value of x:", x)
}
