package main

import (
	"net/http"
)

func main() {
	Routes()
	PrepareDb()
	http.ListenAndServe(":8080", nil)
}
