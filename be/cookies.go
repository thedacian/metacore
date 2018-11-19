package main

import (
	"github.com/gorilla/sessions"
)

//Cookie store
var Store = sessions.NewCookieStore([]byte("changeme"))
