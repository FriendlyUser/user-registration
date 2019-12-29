package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
  "github.com/FriendlyUser/user-registration/handlers"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	// TODO update to exist only on critical envs not being set
	if !exists {
		value = fallback
		log.Fatal(key + " is not set")
	}
	return value
}

func main() {
	// check to ensure env variables are set
	getEnv("GOOGLE_OAUTH_CLIENT_ID", "fail")
	getEnv("GOOGLE_OAUTH_CLIENT_SECRET", "fail")
	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr: fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
