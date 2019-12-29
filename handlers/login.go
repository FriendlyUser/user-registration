package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	util "github.com/FriendlyUser/user-registration/util"
)

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func badReq(w http.ResponseWriter, isJson bool, err string) {
	if !isJson {
			http.Error(w, err, http.StatusFound)
			return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusFound)
	fmt.Fprintf(w, `{"result":"","error":%q}`, err)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HIT THIS ENDPOINT")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(b)

	// Unmarshal
	var s UserLogin
	err = json.Unmarshal(b, &s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err != nil {
			fmt.Println(s)
			w.WriteHeader(http.StatusFound)
			fmt.Fprintf(w, `{"result":"","error":%q}`, err)
			// http.Error(w, "Failed to do anything", http.StatusFound)
			return
	}
	fmt.Println(s)

	// Empty data checking
	emailCheck := util.IsEmpty(s.Email)
	pwdCheck := util.IsEmpty(s.Password)
	if emailCheck || pwdCheck {
			// badReq(w, true, "no valid var")
			fmt.Printf("ErrorCode is -10 : There is empty data.")
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
	}

	dbPwd := "1234!*."                   // DB simulation
	dbEmail := "cihan.ozhan@hotmail.com" // DB simulation

	if s.Email == dbEmail && s.Password == dbPwd {
			fmt.Fprintln(w, "Login succesful!")
	} else {
			badReq(w, true, "no valid var")
			fmt.Fprintln(w, "Login failed!")
	}
	fmt.Printf("Login failed!")
}
