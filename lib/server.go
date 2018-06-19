package lib

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/prokosna/gcp-sample-go/lib/user"
)

func CreateMux() *echo.Echo {
	e := echo.New()
	user.InitRoute(e)
	// note: we don't need to provide the middleware or static handlers, that's taken care of by the platform
	// app engine has it's own "main" wrapper - we just need to hook echo into the default handler
	http.Handle("/", e)
	return e
}
