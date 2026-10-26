package main

import "fmt"

func main() {
	var n int
	var arr [100]int

	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	i := 0

	for i < n {
		j := i + 1

		for j < n {
			if arr[i] == arr[j] {

				// Shift elements to the left
				for k := j; k < n-1; k++ {
					arr[k] = arr[k+1]
				}

				n--
			} else {
				j++
			}
		}

		i++
	}

	fmt.Println("Array after deleting duplicates:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
