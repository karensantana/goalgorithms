package algorithms

import "fmt"

//FizzBuzz  will print all the numbers
//From 1 to N:  will print Fizz if N is
//divisible by 3, will print Buzz if N is
//divisible by 5 and will print both Fizz Buzz if the
//N is divisible by both 3 and 5, otherwise will print Nx
func FizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}

		if !(i == num) {
			fmt.Print(", ")
		}

	}
}
