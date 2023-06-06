package api

import (
	"net/http"

	"github.com/dwdarm/go-url-shortener/cmd"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := cmd.Init()

	app.ServeHTTP(w, r)
}
