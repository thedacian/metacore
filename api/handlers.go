package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//LoginHandler is a function that handles authentication requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is "POST"
	if r.Method == "POST" {

		//Get the json request body and save it into a instance of a user struct
		decoder := json.NewDecoder(r.Body)
		user := User{}
		err := decoder.Decode(&user)
		handleError(err)

		//Get a session if exists
		session, err := Store.Get(r, "session")
		handleError(err)

		//Check if the user/password combination received in the body request matches any db records. If so, set the session cookie.
		if DbCheckUser(user.Username, user.Password) == true {
			session.Values["authenticated"] = true
			session.Values["username"] = user.Username
			session.Save(r, w)
		}

	} else {
		//If request method is NOT "POST" respond with 405 Not allowed
		http.Error(w, http.StatusText(405), 405)
	}
}

//LogoutHandler handles user log out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Get a session if exists
		session, err := Store.Get(r, "session")
		handleError(err)
		//Delete the session values
		session.Values["authenticated"] = false
		session.Values["username"] = nil
		session.Save(r, w)

	} else {
		//If request method is other than "POST" respond with 405 Not allowed
		http.Error(w, http.StatusText(405), 405)
	}

}

//SecretHandler will handle a secret area to demonstrate how to protect API routes
func SecretAreaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a secret area")
}
