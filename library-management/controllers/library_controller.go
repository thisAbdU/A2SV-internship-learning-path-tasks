package controllers

import (
	"fmt"
	"library_management/models"
)

type LibraryController struct {
	library models.LibraryManager
}

func NewLibraryController(library models.LibraryManager) *LibraryController {
	return &LibraryController{library: library}
}

func (c *LibraryController) AddBook() {
	var id int
	var title, author string

	fmt.Print("Enter Book ID: ")
	fmt.Scanf("%d\n", &id)

	fmt.Print("Enter Book Title: ")
	fmt.Scanf("%s\n", &title)

	fmt.Print("Enter Book Author: ")
	fmt.Scanf("%s\n", &author)

	book := models.Book{ID: id, Title: title, Author: author}
	c.library.AddBook(book)
}

func (c *LibraryController) RemoveBook() {
	var id int
	fmt.Print("Enter Book ID to remove: ")
	fmt.Scanf("%d\n", &id)

	c.library.RemoveBook(id)
}

func (c *LibraryController) BorrowBook() {
	var bookID, memberID int
	fmt.Print("Enter Book ID to borrow: ")
	fmt.Scanf("%d\n", &bookID)

	fmt.Print("Enter Member ID: ")
	fmt.Scanf("%d\n", &memberID)

	err := c.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (c *LibraryController) ReturnBook() {
	var bookID, memberID int
	fmt.Print("Enter Book ID to return: ")
	fmt.Scanf("%d\n", &bookID)

	fmt.Print("Enter Member ID: ")
	fmt.Scanf("%d\n", &memberID)

	err := c.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (c *LibraryController) ListAvailableBooks() {
	books := c.library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) ListBorrowedBooks() {
	var memberID int
	fmt.Print("Enter Member ID: ")
	fmt.Scanf("%d\n", &memberID)

	books := c.library.ListBorrowedBooks(memberID)
	fmt.Printf("Borrowed Books for Member ID %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
