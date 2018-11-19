package main

import (
	"encoding/json"
	"net/http"
)

//LoginHandler is a function that handles authentication requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//Get the json request body and save it into a instance of a user struct
		decoder := json.NewDecoder(r.Body)
		user := User{}
		err := decoder.Decode(&user)
		if err != nil {
			handleError(err)
		}

		//Get a session if exists
		session, err := Store.Get(r, "session")

		if DbCheckUser(user.Username, user.Password) == true {
			session.Values["authenticated"] = true
			session.Values["username"] = user.Username
			session.Save(r, w)
		}

	} else {
		//If request method is other than "POST" respond with 405 Not allowed
		http.Error(w, http.StatusText(405), 405)
	}

}
