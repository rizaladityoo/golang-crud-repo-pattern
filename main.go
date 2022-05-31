package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	//books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	router := gin.Default()

	router.GET("/", bookHandler.RootHandler)
	router.GET("/books", bookHandler.GetBooks)
	router.GET("/book/:id", bookHandler.GetBook)
	router.POST("/books", bookHandler.PostBooksHandler)
	// router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run(":8080")
}
