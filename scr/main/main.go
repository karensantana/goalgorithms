package main

import (
	"algorithms"
	"fmt"
)

func main() {
	fmt.Println(algorithms.NumInList([]int{1, 2, 5, 9, 8, 9, 41, 8}, 8))
	fmt.Println("Sum list result: ", algorithms.SumAList([]int{4, 4, 4, 4, 4, 4}))
	fmt.Println(algorithms.ReverseString("Hello World!"))

}
