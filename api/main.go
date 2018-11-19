package main

import (
	"net/http"
)

func main() {
	Routes()
	DbInit()
	http.ListenAndServe(":8080", nil)
}
