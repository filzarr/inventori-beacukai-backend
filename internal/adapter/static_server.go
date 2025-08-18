package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

// SetupStaticFiles configures static file serving for the web assets
func SetupStaticFiles(app *fiber.App) {
	// Serve static CSS files
	app.Use("/static/css", filesystem.New(filesystem.Config{
		Root:       http.Dir("./web/static/css"),
		PathPrefix: "css",
		Browse:     false,
	}))

	// Serve static JS files
	app.Use("/static/js", filesystem.New(filesystem.Config{
		Root:       http.Dir("./web/static/js"),
		PathPrefix: "js",
		Browse:     false,
	}))

	// Serve other static assets if needed
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   http.Dir("./web/static"),
		Browse: false,
	}))
}
