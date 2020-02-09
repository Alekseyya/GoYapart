package main

import (
	"fmt"
	"strconv"
)

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
	ch <- result
}

// func factorial2(n int) {
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	fmt.Println(n, "-", result)
// }

func WriteText(i int, ch chan int) {
	factorial(500, ch)
	fmt.Println("Hello World" + strconv.Itoa(i))
	ch <- i
}

func main() {
	var intChan = make(chan int, 10)

	for index := 0; index < 10; index++ {
		factorial(index+20, intChan)
	}

	fmt.Println(<-intChan)
	fmt.Println("The End")

	// for i := 0; i < 10; i++ {
	// 	go WriteText(i, intChan)

	// }
	// fmt.Println(<-intChan)
	// fmt.Println("The End")

}
