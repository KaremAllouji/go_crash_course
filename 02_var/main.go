package main

import "fmt"

// global var are also allowed

func main() {
	//var <name of var> <type of var>
	// strings dont need keyword "string"
	var name string = "Go"
	var Name = "World"

	fmt.Println(name, Name)
	fmt.Println(name + Name)

	// get type of the var
	fmt.Printf("%T\n", name)

	// constant
	const isCons = true

	// Shorthand method
	nameofvar := "whatalanguage"
	fmt.Printf("%T\n", nameofvar)

	// multiple
	var01, var02 := "var01", "var02"
	fmt.Println(var01, var02)

}

// go run <name.ext>
// go build <name.ext>
// to create bin :
// go mod init <name.ext>
// go install
