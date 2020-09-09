package book

import (
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Books struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx)  {

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != err {
		log.Println(err.Error())
	}

	var books []Books

	db.Find(&books)

	c.JSON(books)

}

func GetBook(c *fiber.Ctx)  {

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != err {
		log.Println(err.Error())
	}

	var book Books

	id := c.Params("id")
	db.Find(&book, id)

	c.JSON(book)

}

func NewBook(c *fiber.Ctx)  {

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != err {
		log.Println(err.Error())
	}

	book := new(Books)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)

	c.JSON(book)

}

func DeleteBook(c *fiber.Ctx) {

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != err {
		log.Println(err.Error())
	}

	var book Books
	id := c.Params("id")

	db.First(&book, id)
	if book.Title != "" {
		c.Status(500).Send("No Book Found With ID")
		return
	}

	db.Delete(&book)

	c.Send("Book Successfully deleted")

}
