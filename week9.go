package main

import "fmt"

func main() {
	var n int
	var sum int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	var arr [100]int

	fmt.Println("Enter the elements:")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sum = 0

	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}

	average := float64(sum) / float64(n)

	fmt.Println("Sum =", sum)
	fmt.Println("Average =", average)
}
