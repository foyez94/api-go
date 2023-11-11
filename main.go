package main

import "fmt"

// type Book struct {
// 	Title     string
// 	Author    string
// 	pages     int
// 	price     float64
// 	publisher string
// 	ISBN      string
// }

func main() {
	//var b1 Book
	// b1.Title = "Golang Book"
	// b1.Author = "Abdul Foyez"
	// b1.pages = 600
	// b1.price = 300.89
	// b1.publisher = "zoomin"
	// b1.ISBN = "1234567890"

	// fmt.Println(b1)
	// fmt.Println(b1.Title)
	// fmt.Println(b1.Author)
	// fmt.Println(b1.pages)
	// fmt.Println(b1.price)
	// fmt.Println(b1.publisher)
	// fmt.Println(b1.ISBN)

	b1 := struct {
		Title     string
		Author    string
		price     float64
		pages     int
		publisher string
	}{
		Title:     "Golang Book",
		Author:    "Abdul Foyez",
		pages:     600,
		price:     300.89,
		publisher: "zoomin",
	}
	fmt.Println(b1)
}
