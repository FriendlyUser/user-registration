package handlers

import (
	"fmt"
	"net/http"
	"strings"
	util "github.com/FriendlyUser/user-registration/util"
)

func tokenParse(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	reqToken = splitToken[1]
	isValid := util.ParseJWT(reqToken)
	if (isValid == true) {
		fmt.Println(isValid)
	} else {
		fmt.Println("Mission Failed")
	}
}
