package handlers

import (
	"log"
	"os"

	"github.com/ardanlabs/service/internal/platform/web"
)

// API constructs an http.Handler with all application routes defined.
func API(shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, log)

	c := check{}
	app.Handle("GET", "/testy", c.health)

	return app
}
