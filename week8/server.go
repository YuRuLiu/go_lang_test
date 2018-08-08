package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Book struct {
	Author string `json:"author"`
	Name   string `json:"name"`
}

var books = []Book{}

func main() {
	e := echo.New()
	e.GET("/booklist", getBookList)
	e.POST("/booking", postBook)
	e.Logger.Fatal(e.Start(":1323"))
}

func getBookList(c echo.Context) error {
	// book1 := Book{
	// 	Author: "東野圭吾",
	// 	Name:   "嫌疑犯X的獻身",
	// }
	// book2 := Book{
	// 	Author: "東野圭吾",
	// 	Name:   "怪人們",
	// }
	books = append(books)

	return c.JSONPretty(http.StatusOK, books, " ")
}

func postBook(c echo.Context) (err error) {
	name := c.FormValue("name")
	bookName := c.FormValue("bookName")
	author := c.FormValue("author")

	book := Book{author, bookName}

	books = append(books, book)

	return c.String(http.StatusOK, name+"訂購成功")
}
