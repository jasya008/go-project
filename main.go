	package main

	import (
		"net/http"
		"github.com/gin-gonic/gin"
		"errors"
	)

	type book struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Author   string `json:"author"`
		Quantity int    `json:"quantity"`
	}

	var books = []book{
		{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	}

	func getBooks(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, books)
	}

	func bookByID (c * gin.Context){
		id := c.Param("id") 
		book, err := getBookById(id)

		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book Not Found."})
			return
		}

		c.IndentedJSON(http.StatusOK, book)
	}


	func deleteBook( c* gin.Context){
		id := c.Param("id") 
        for i, b := range books {
		if b.ID == id {
			  books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message":"Book delete succesful! "})
			return
		}
	 
	  }

	  c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})

	}


	func getBookById(id string) (*book, error){
	  for i, b := range books{
	  	if b.ID == id {
	  		return &books[i], nil
	  	}

	  }
	  return nil, errors.New("book not foud")
	}

	func checkoutBook (c * gin.Context){
		id, ok := c.GetQuery("id")

		if !ok {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Missing id query parametr."})
			return
		}

		book, err := getBookById(id)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book Not Found."})
			return
		}


		if book.Quantity <= 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Book not avialable"})
			return
		}

		book.Quantity -= 1 
		c.IndentedJSON(http.StatusOK,book)
	}

	func createBook(c * gin.Context){
		var newBook book

		if err := c.BindJSON(&newBook); err!= nil{
			return
		}


		books = append(books, newBook)

		c.IndentedJSON(http.StatusCreated, newBook)
	}

	func returnBook(c * gin.Context){
	    id, ok := c.GetQuery("id")

		if !ok {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Missing id query parametr."})
			return
		}

		book, err := getBookById(id)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book Not Found."})
			return
		}

		book.Quantity += 1
		c.IndentedJSON(http.StatusOK, book)
	}

	func main() {
		router := gin.Default()
		router.GET("/books", getBooks)
		router.POST("/books", createBook)
		router.PATCH("/checkout", checkoutBook)
		router.PATCH("/return", returnBook)
		router.GET("/books/:id", bookByID)
		router.DELETE("/books/:id", deleteBook)
		router.Run("localhost:8080")
	}
