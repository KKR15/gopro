package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	reverse := ""

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}

	if str == reverse {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome")
	}
}
