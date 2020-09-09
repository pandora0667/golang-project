package main

import (
	book "./book"
	"github.com/gofiber/fiber"
)

const PORT = 3000

func main()  {

	app := fiber.New()
	SetupRouter(app)

	app.Listen(PORT)

}

func HelloWorld(c *fiber.Ctx)  {

	c.Send("Hello, Seongwon")

}

func SetupRouter(app *fiber.App)  {

	app.Get("/", HelloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.Getbook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}
