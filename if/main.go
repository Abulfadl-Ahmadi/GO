package main

import "fmt"

func main() {

	var user_input int

	fmt.Print("Enter a number:")
	fmt.Scan(&user_input)

	if user_input == 0 {

		fmt.Print("Zero")

	} else if user_input == 1 {

		fmt.Print("One")

	} else {
		fmt.Print("Bad number")
	}

}
