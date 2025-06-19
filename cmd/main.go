package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	// New fiber app
	app := fiber.New()

	// http.Handler -> fiber.Handler
	app.Get("/", adaptor.HTTPHandler(handler(greet)))

	// http.HandlerFunc -> fiber.Handler
	app.Get("/func", adaptor.HTTPHandlerFunc(greetfunc))

	// Listen on port 3000
	app.Listen(":3000")
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func greetfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Hello!")
}
