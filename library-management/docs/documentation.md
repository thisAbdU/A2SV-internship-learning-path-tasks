# Library Management System

## Overview

The Library Management System is a command-line application that allows you to manage a library's collection of books and track borrowed books by members. The application supports adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Installation

### Prerequisites

- Go (version 1.16 or later)

### Steps

1. **Clone the Repository**

   Clone the repository to your local machine:

   ```sh
   git clone https://github.com/thisAbdU/A2SV-internship-learning-path-tasks.git

   cd library-management
   ```

2. **Build the Application**

   Build the Go application:

   ```sh

   go build -o library-management

   ```

3. **Run the Application**

   Run the built application:

   ```sh

   ./library-management
   ```

## Usage

When you run the application, you will see a menu with various options:

```
Library Management System
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
7. Exit

Enter your choice:
```

### Commands

1. **Add Book**

   Adds a book to the library.

   - **Prompt**: Enter Book ID, Enter Book Title, Enter Book Author

   - **Example**:

     ```
     Enter Book ID: 1
     Enter Book Title: Go Programming
     Enter Book Author: John Doe
     Book added successfully.
     ```

2. **Remove Book**

   Removes a book from the library.

   - **Prompt**: Enter Book ID to remove
   - **Example**:

     ```
     Enter Book ID to remove: 1
     Book removed successfully.
     ```

3. **Borrow Book**

   Allows a member to borrow a book if it is available.

   - **Prompt**: Enter Book ID to borrow, Enter Member ID
   - **Example**:

     ```
     Enter Book ID to borrow: 1
     Enter Member ID: 1001
     Book borrowed successfully.
     ```

4. **Return Book**

   Allows a member to return a borrowed book.

   - **Prompt**: Enter Book ID to return, Enter Member ID
   - **Example**:

     ```
     Enter Book ID to return: 1
     Enter Member ID: 1001
     Book returned successfully.
     ```

5. **List Available Books**

   Lists all books that are currently available.

   - **Example**:

     ```
     Available Books:
     ID: 2, Title: Advanced Go, Author: Jane Doe
     ```

6. **List Borrowed Books**

   Lists all books borrowed by a specific member.

   - **Prompt**: Enter Member ID
   - **Example**:

     ```
     Enter Member ID: 1001
     Borrowed Books for Member ID 1001:
     ID: 1, Title: Go Programming, Author: John Doe
     ```

7. **Exit**

   Exits the application.

## Implementation Details

### LibraryController

The `LibraryController` struct handles operations for the library and interacts with the `LibraryManager` interface.

- **AddBook**: Prompts the user for book details and adds the book to the library.
- **RemoveBook**: Prompts the user for a book ID and removes the book from the library.
- **BorrowBook**: Prompts the user for a book ID and member ID, and borrows the book for the member.
- **ReturnBook**: Prompts the user for a book ID and member ID, and returns the book to the library.
- **ListAvailableBooks**: Lists all available books in the library.
- **ListBorrowedBooks**: Lists all books borrowed by a specific member.

### Models

- **Book**: Represents a book in the library with fields for ID, Title, Author, and Status.
- **Member**: Represents a library member with fields for ID, Name, and BorrowedBooks.
- **Library**: Represents the library and manages the collection of books and borrowed books.