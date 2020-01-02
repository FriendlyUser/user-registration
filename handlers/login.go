package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	// "database/sql"
	util "github.com/FriendlyUser/user-registration/util"
	db "github.com/FriendlyUser/user-registration/db"
)

func badReq(w http.ResponseWriter, isJson bool, err string) {
	if !isJson {
			http.Error(w, err, http.StatusFound)
			return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	fmt.Fprintf(w, `{"error":%q}`, err)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Login Endpoint")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var s util.User
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
			badReq(w, true, "no valid var")
			fmt.Printf("ErrorCode is -10 : There is empty data.")
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
	}
	var User util.User
	User, err = db.GetUser(s.Email)
	dbPwd := User.Password  // DB simulation
	dbEmail := User.Email // DB simulation
	fmt.Printf(User.Password)
	fmt.Printf(User.Email)
	if s.Email == dbEmail && s.Password == dbPwd {
			// fmt.Fprintln(w, "Login succesfull!")
			// fmt.Printf("Login failed!")
			// generate JWT
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			jwtToken := util.CreateJWT(s.Email)
			fmt.Fprintf(w, `{"token":%q}`, jwtToken)
	} else {
			badReq(w, true, "Login Failed")
	}
}

func userRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Register Endpoint")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var s util.User
	err = json.Unmarshal(b, &s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Empty data checking
	emailCheck := util.IsEmpty(s.Email)
	pwdCheck := util.IsEmpty(s.Password)
	if emailCheck || pwdCheck {
			// badReq(w, true, "no valid var")
			fmt.Printf("ErrorCode is -10 : There is empty data.")
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
	}
	// insert to database	
	err = db.CreateUser(s.Email, s.Password)
	if err != nil {
		// if user exists already send an error
		// http.Error(w, err.Error(), 500)
		badReq(w, true, err.Error())
		// return
	}
	fmt.Fprintf(w, `{"email":%q,"password":%q}`, s.Email, s.Password)
}
