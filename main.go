package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/FriendlyUser/user-registration/handlers"
	db "github.com/FriendlyUser/user-registration/db"
	util "github.com/FriendlyUser/user-registration/util"
)

func main() {
	// check to ensure env variables are set
	util.GetEnv("GOOGLE_OAUTH_CLIENT_ID", "fail")
	util.GetEnv("GOOGLE_OAUTH_CLIENT_SECRET", "fail")
	util.GetEnv("JWT_SECRET", "fail")
	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr: fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	db.InitDB()
	util.JWTInit()
	// defer db.Close()

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
