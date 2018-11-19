package main

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/login", WithLogging(WithTracing(LoginHandler)))
	http.HandleFunc("/logout", WithLogging(WithTracing(RequireAuth(LogoutHandler))))
	http.HandleFunc("/secret", WithLogging(WithTracing(RequireAuth(SecretAreaHandler))))
}
