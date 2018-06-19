// +build appengine

package main

import (
	"net/http"

	"github.com/prokosna/gcp-sample-go/lib"
)

func init() {
	e := lib.CreateMux()
	// note: we don't need to provide the middleware or static handlers, that's taken care of by the platform
	// app engine has it's own "main" wrapper - we just need to hook echo into the default handler
	http.Handle("/", e)
}
