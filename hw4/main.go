package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	books []Book
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) RemoveBook(title string) {
	for i, book := range l.books {
		if book.Title == title {
			l.books = append(l.books[:i], l.books[i+1:]...)
			return
		}
	}
}

func (l *Library) DisplayBooks() {
	fmt.Println("Список книг в библиотеке:")
	for _, book := range l.books {
		fmt.Printf("Название: %s, Автор: %s, Год издания: %d\n", book.Title, book.Author, book.Year)
	}
}

type Readable interface {
	Read()
}

func (b Book) Read() {
	fmt.Printf("Читатель читает книгу '%s' автора %s.\n", b.Title, b.Author)
}

func ReadBook(r Readable) {
	r.Read()
}

func main() {
	library := Library{}

	library.AddBook(Book{Title: "Война и мир", Author: "Лев Толстой", Year: 1869})
	library.AddBook(Book{Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Year: 1967})
	library.AddBook(Book{Title: "1984", Author: "Джордж Оруэлл", Year: 1949})

	library.DisplayBooks()

	ReadBook(library.books[0])
}
