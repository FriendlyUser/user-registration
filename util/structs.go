package util 

type User struct {
	Email string `json:"email", db:"username"`
	Password string `json:"password", db:"password"`
}

type Token struct {
	Token string `json:"token"`
}