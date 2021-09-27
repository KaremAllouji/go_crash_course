package main

import "fmt"

// func <nameofFunc>(param) return type {}
func greeting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(greeting("5ra"))
}
