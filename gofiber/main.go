package main

import (
	book "./book"
	database "./database"
	"github.com/gofiber/fiber"
)

const PORT = 3000

func main()  {

	database.InitDatabase()
	app := fiber.New()
	SetupRouter(app)
	app.Listen(PORT)

}

func HelloWorld(c *fiber.Ctx)  {

	c.Send("Hello, Seongwon")

}

func SetupRouter(app *fiber.App)  {

	app.Get("/", HelloWorld)
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)

	app.Post("/api/v1/book/:id", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}
