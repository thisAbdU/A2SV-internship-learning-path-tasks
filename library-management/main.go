package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/services"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	for {
		fmt.Println("\nWelcom Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.BorrowBook()
		case 4:
			controller.ReturnBook()
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks()
		case 7:
			fmt.Println("Exiting...")
			fmt.Println()
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}