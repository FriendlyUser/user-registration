package util 

type User struct {
	Email string `json:"email", db:"username"`
	Password string `json:"password", db:"password"`
	Description string `json:"description", db:"description"`
	Phone string `json:"phone", db:"phone"`
	JobTitle string `json:"job_title", db:"job_title"`
}

type Token struct {
	Token string `json:"token"`
}
