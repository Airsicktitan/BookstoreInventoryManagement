package main

import (
	"errors"
	"fmt"
)

type typeOfBook string

const (
	eBook        typeOfBook = "eBook"
	physicalCopy typeOfBook = "Physical"
	rental       typeOfBook = "Rental"
)

type Author struct {
	name                 string
	numberOfBooksWritten int
}

type Book struct {
	name          string
	bookType      typeOfBook
	numPages      int
	price         float64
	isInStock     bool
	isBestSelling bool
	author        Author
}

type Bookstore struct {
	authors []Author
	books   []Book
}

func (b *Bookstore) AddBookToInventory(bookName string, bookType typeOfBook, numPages int, price float64, isInStock, isBestSelling bool, auth Author) error {
	if bookName == "" {
		return errors.New("Cannot have a blank name for this book. Please enter a book title.")
	}

	if numPages < 0 {
		return errors.New("Pages cannot be negative. Check your entry and try again.")
	}

	if price < 0 {
		return errors.New("Price cannot be less than zero. Check your entry and try again.")
	}

	newBook := Book{
		name:          bookName,
		bookType:      bookType,
		numPages:      numPages,
		price:         price,
		isInStock:     isInStock,
		isBestSelling: isBestSelling,
		author:        auth,
	}

	if newBook.bookType == rental {
		newBook.price = 0.00
	}

	b.books = append(b.books, newBook)
	return nil
}

func (b *Bookstore) AddAuthorToList(authName string, numOfBooksWritten int) {
	newAuth := Author{
		name:                 authName,
		numberOfBooksWritten: numOfBooksWritten,
	}

	b.authors = append(b.authors, newAuth)
}

func (b Bookstore) CountBooksAvailable() (int, int) {
	countInStock, countOutStock := 0, 0
	for _, book := range b.books {
		if book.isInStock {
			countInStock += 1
		} else {
			countOutStock += 1
		}
	}
	return countInStock, countOutStock
}

func (b *Bookstore) UpdateBookInInventory(bookName string, bookType typeOfBook, numPages int, price float64, isInStock, isBestSelling bool, auth Author) error {
	if bookName == "" {
		return errors.New("Cannot have a blank name for this book. Please enter a book title.")
	}

	if numPages < 0 {
		return errors.New("Pages cannot be negative. Check your entry and try again.")
	}

	if price < 0 {
		return errors.New("Price cannot be less than zero. Check your entry and try again.")
	}

	updateBook := Book{
		name:          bookName,
		bookType:      bookType,
		numPages:      numPages,
		price:         price,
		isInStock:     isInStock,
		isBestSelling: isBestSelling,
		author:        auth,
	}

	if updateBook.bookType == rental {
		updateBook.price = 0.00
	}

	for i, book := range b.books {
		if book.name == bookName {
			b.books[i] = updateBook
			return nil
		}
	}

	return errors.New("Book not found in inventory, try again.")
}

func (b *Bookstore) UpdateAuthorInList(auth string, author Author) error {
	updatedAuth := Author{
		name:                 author.name,
		numberOfBooksWritten: author.numberOfBooksWritten,
	}

	for i, author := range b.authors {
		if author.name == auth {
			b.authors[i] = updatedAuth
			return nil
		}
	}

	return errors.New("Author not found, please try again.")
}

func (b *Bookstore) DeleteBookInInventory(bookName string) error {
	if bookName == "" {
		return errors.New("Cannot have a blank name for this book. Please enter a book title.")
	}

	for i, book := range b.books {
		if book.name == bookName {
			b.books = append(b.books[:i], b.books[i+1:]...)
			return nil
		}
	}

	return errors.New("Book not found in inventory, try again.")
}

func (b *Bookstore) DeleteAuthorInInventory(authName string) error {
	if authName == "" {
		return errors.New("Cannot have a blank name for this author. Please enter an author name.")
	}

	for i, auth := range b.authors {
		if auth.name == authName {
			b.authors = append(b.authors[:i], b.authors[i+1:]...)
			return nil
		}
	}

	return errors.New("Author not found in inventory, try again.")
}

func main() {
	fmt.Print("\n")
	title := "Bookstore Inventory management"
	fmt.Println(title)

	for i := 0; i < len(title); i++ {
		fmt.Print("-")
	}
	fmt.Print("\n\n")

	author := Author{
		name:                 "JRR Tolkien",
		numberOfBooksWritten: 12,
	}
	book := Book{"Lord of The Rings", eBook, 1_000, 19.99, true, true, author}
	bookStore := Bookstore{}

	bookStore.AddAuthorToList(author.name, author.numberOfBooksWritten)
	err := bookStore.AddBookToInventory(book.name, book.bookType, book.numPages, book.price, book.isInStock, book.isBestSelling, book.author)
	if err != nil {
		fmt.Println(err)
	}

	book = Book{"The Hobbit", physicalCopy, 1_000, 39.99, false, true, author}
	err = bookStore.AddBookToInventory(book.name, book.bookType, book.numPages, book.price, book.isInStock, book.isBestSelling, book.author)
	if err != nil {
		fmt.Println(err)
	}

	book = Book{
		name:          "The Hobbit: Rental",
		bookType:      rental,
		isInStock:     false,
		isBestSelling: false,
		price:         19.99,
		numPages:      1_000,
		author:        author,
	}

	err = bookStore.AddBookToInventory(book.name, book.bookType, book.numPages, book.price, book.isInStock, book.isBestSelling, book.author)
	if err != nil {
		fmt.Println(err)
	}

	book = Book{
		name:          "The Hobbit: Rental",
		bookType:      rental,
		isInStock:     true,
		isBestSelling: false,
		price:         -1,
		numPages:      1_000,
		author:        author,
	}

	err = bookStore.AddBookToInventory(book.name, book.bookType, book.numPages, book.price, book.isInStock, book.isBestSelling, book.author)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bookStore.books)

	countBooks, countBooksUn := bookStore.CountBooksAvailable()
	fmt.Print("\n")
	fmt.Printf("Total books Available: %d\nTotal books unavailable: %d\n", countBooks, countBooksUn)
	fmt.Print("\n")

	err = bookStore.UpdateBookInInventory(book.name, physicalCopy, book.numPages, 20.19, true, false, Author{"Dr Suess", 60})
	if err != nil {
		fmt.Println(err)
	}

	err = bookStore.DeleteBookInInventory("Lord of The Rings")

	fmt.Println(bookStore.books)

	countBooks, countBooksUn = bookStore.CountBooksAvailable()
	fmt.Print("\n")
	fmt.Printf("Total books Available: %d\nTotal books unavailable: %d\n", countBooks, countBooksUn)
	fmt.Print("\n")

}
