package main

import (
	"net/http"
)

func (b *AppBackend) HealthHandler(w http.ResponseWriter, r *http.Request) {
	HandleSuccessJson(w, 200, "Ok!")
}
