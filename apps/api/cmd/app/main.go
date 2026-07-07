package main

import (
	"net/http"

	"github.com/dtt4h/devodoro/internal/server"
)

func main() {
	r := server.NewRouter()
	http.ListenAndServe(":8080", r)

}
