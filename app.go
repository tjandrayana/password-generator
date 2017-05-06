package main

import (
	"net/http"

	"github.com/web-profile/backend"
)

func main() {
	backend.Test()
	http.ListenAndServe(":8080", nil)
}
