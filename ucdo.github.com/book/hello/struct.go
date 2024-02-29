package main

import "fmt"

// Book def a struct
type Book struct {
	bookName string
	author   string
	classify string
	price    float32
}

func main() {
	book := Book{"go语言圣经", "author", "计算机类/编程", 99.99}
	fmt.Println(book.bookName)
	test(&book)
	fmt.Println(book.bookName)
}

func test(s *Book) {
	s.bookName = "effective go"
}
