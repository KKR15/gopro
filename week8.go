package main

import "fmt"

type Contact struct {
	Name        string
	Age         int
	PhoneNumber string
	Email       string
}

func main() {
	var contact Contact

	// Read values
	fmt.Print("Enter Name: ")
	fmt.Scanln(&contact.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&contact.Age)

	fmt.Print("Enter Phone Number: ")
	fmt.Scanln(&contact.PhoneNumber)

	fmt.Print("Enter Email: ")
	fmt.Scanln(&contact.Email)

	// Display values
	fmt.Println("\n--- Contact Details ---")
	fmt.Println("Name:", contact.Name)
	fmt.Println("Age:", contact.Age)
	fmt.Println("Phone Number:", contact.PhoneNumber)
	fmt.Println("Email:", contact.Email)
}
