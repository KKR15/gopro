package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello Go Programming"

	fmt.Println("Contains:", strings.Contains(str, "Go"))

	fmt.Println("ContainsAny:", strings.ContainsAny(str, "xyzG"))

	fmt.Println("Count:", strings.Count(str, "o"))

	fmt.Println("EqualFold:", strings.EqualFold("Hello", "hello"))
}
