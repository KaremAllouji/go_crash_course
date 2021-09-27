package main

import "fmt"

func main() {
	//Arrays
	var arr [2]string

	arr[0] = "arr0"
	arr[1] = "arr1"

	fmt.Println(arr)
	fmt.Println(arr[0:1])

	x := 5
	y := 10

	// if else
	if x < y {
		fmt.Println("x < y")
	} else {
		fmt.Println("else")
	}

	// switch

	switch x {
	case 0:
		fmt.Printf("switch case 0")
	}

	// loops
	//Long method

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
