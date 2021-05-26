package web

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
)

// ctxKey ..
type ctxKey int

// KeyValues ...
const KeyValues ctxKey = 1

// Values ...
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// Handler ...
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App ...
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// SignalShutdown ...
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {

	// Specific
	handler = wrapMiddleware(mw, handler)
	// Application general
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {

		v := Values{
			TraceID: uuid.New().String(),
			Now:     time.Now(),
		}
		ctx := context.WithValue(r.Context(), KeyValues, &v)

		if err := handler(ctx, w, r); err != nil {
			// Handler error
			a.SignalShutdown()
			return
		}

		// BOILERPLATE
	}
	a.ContextMux.Handle(method, path, h)
}

// NewApp ...
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}

	return &app
}
