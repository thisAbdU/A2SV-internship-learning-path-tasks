package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type Library struct {
	Books          map[int]models.Book
	BorrowedBooks  map[int]int // bookID -> memberID
	MemberBorrowed map[int][]models.Book // memberID -> list of borrowed bookIDs
}


func NewLibrary() *Library {
	return &Library{
		Books:          make(map[int]models.Book),
		BorrowedBooks:  make(map[int]int),
		MemberBorrowed: make(map[int][]models.Book),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = models.Available
	l.Books[book.ID] = book
	fmt.Println("Book added successfully.")
}

func (l *Library) RemoveBook(bookID int) {
	if _, exists := l.Books[bookID]; !exists {
		fmt.Printf("Book ID %d does not exist.\n", bookID)
		return
	}

	delete(l.BorrowedBooks, bookID)
	
	for memberID, books := range l.MemberBorrowed {
		for i, book := range books {
			if book.ID == bookID {
				l.MemberBorrowed[memberID] = append(books[:i], books[i+1:]...)
				break
			}
		}
	}

	delete(l.Books, bookID)
	fmt.Println("Book removed successfully.")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status == models.Borrowed {
		return errors.New("book already borrowed")
	}

	book.Status = models.Borrowed
	l.Books[bookID] = book
	l.BorrowedBooks[bookID] = memberID
	l.MemberBorrowed[memberID] = append(l.MemberBorrowed[memberID], book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status == models.Available {
		return errors.New("book is not borrowed")
	}
	if borrower, exists := l.BorrowedBooks[bookID]; !exists || borrower != memberID {
		return errors.New("book was not borrowed by this member")
	}

	book.Status = models.Available
	l.Books[bookID] = book
	delete(l.BorrowedBooks, bookID)

	for i, book := range l.MemberBorrowed[memberID] {
		if book.ID == bookID {
			l.MemberBorrowed[memberID] = append(l.MemberBorrowed[memberID][:i], l.MemberBorrowed[memberID][i+1:]...)
			break
		}
	}

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range l.Books {
		if book.Status == models.Available {
			availableBooks = append(availableBooks, book)
		}
	}
	
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	return l.MemberBorrowed[memberID]
}