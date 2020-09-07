package main

import "github.com/gofiber/fiber"

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
}
