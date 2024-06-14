package main

import (
	"log"

	"github.com/rayfanaqbil/ws-rayfan2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/rayfanaqbil/ws-rayfan2024/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/rayfanaqbil/ws-rayfan2024/docs"

	// @title TES SWAGGER ULBI
	// @version 1.0
	// @description This is a sample swagger for Fiber

	// @contact.name API Support
	// @contact.url https://github.com/rayfanaqbil
	// @contact.email rayfana09@gmail.com

	// @host ws-rayfan2024-7c90fe3029b2.herokuapp.com
	// @BasePath /
	// @schemes https http
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
