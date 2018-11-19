package main

import (
	"log"
	"net/http"
)

//Middleware provides a convenient mechanism for filtering HTTP rqeuests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler

//WithLogging logs the remote address of the request
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

//WithTracing traces the URL path from where the request was sent
func WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request to %s", r.URL)
		next.ServeHTTP(w, r)
	}
}

//WithAllowCORS - by default, most browsers block cross origin requests if these headers are not set. Use for development only
func WithAllowCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		w.Header().Set("Access-Control-Allow-Origin", "*") // <--- THIS IS UNSAFE AND SHALL ONLY BE USED FOR DEVELOPMENT. SET UP CORS HEADERS WHEN GOING TO PRODUCTION.
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	}

}
