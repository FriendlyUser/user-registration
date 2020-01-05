package handlers

import (
	"net/http"
	db "github.com/FriendlyUser/user-registration/db"
	"fmt"
	"encoding/json"
)

// probably handled by the login function
func createUser(w http.ResponseWriter, r *http.Request) {

	err := db.CreateUser("test", "test")
	if err != nil {
		// if user exists already send an error
		badReq(w, true, err.Error())
		// return
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	allUsers,err := db.GetAllUsers()
	if err != nil {
		// if user exists already send an error
		badReq(w, true, err.Error())
		return
	}
	b, err := json.Marshal(allUsers)
	if err != nil {
		badReq(w, true, err.Error())
		return
	}
	fmt.Fprintf(w,string(b))
}
