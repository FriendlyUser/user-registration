package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/",  http.FileServer(http.Dir("dist/")))
	// Root
	mux.Handle("/templates",  http.FileServer(http.Dir("templates/")))

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	// User Registration and Authentication
	mux.HandleFunc("/auth/login", userLogin)
	mux.HandleFunc("/auth/register", userRegister)

	// Get all users
	mux.HandleFunc("/users", GetAllUsers)
	// mux.HandleFunc("/getUser", GetUsers)
	// test JWT token
	mux.HandleFunc("/jwt/test", tokenParse)
	return mux
}
