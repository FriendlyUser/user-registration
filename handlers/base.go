package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/",  http.FileServer(http.Dir("front-end/dist/")))
	// Root
	mux.Handle("/templates",  http.FileServer(http.Dir("templates/")))

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	// User Registration and Authentication
	mux.HandleFunc("/auth/login", userLogin)
	mux.HandleFunc("/auth/register", userRegister)

	// test JWT token

	mux.HandleFunc("/jwt/test", tokenParse)
	return mux
}
