package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dimfeld/httptreemux"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
	//och      *ochttp.Handler
	shutdown chan os.Signal
	log      *log.Logger
	//mw       []Middleware
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, log *log.Logger) *App {
	app := App{
		TreeMux:  httptreemux.New(),
		shutdown: shutdown,
		log:      log,
	}

	return &app
}

func (a *App) Handle(verb, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request, param map[string]string) {
		if err := handler(r.Context(), w, r, param); err != nil {

		}
	}

}
