package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

//Store - in memory session database
var Store = sessions.NewCookieStore([]byte("changeme"))

//IsAuthenticated checks the session id and returns the user if the session is valid
func IsAuthenticated(r *http.Request) bool {
	session, err := Store.Get(r, "session")
	handleError(err)
	if session.Values["authenticated"] == true {
		return true
	}
	return false // Not sure if this is safe
}
