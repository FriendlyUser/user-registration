package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	util "github.com/FriendlyUser/user-registration/util"
)

// the database is free tier elephantsql, don't intend to ever upgrade
// access if you want
const (
  host     = "rajje.db.elephantsql.com"
  port     = 5432
  user     = "lwdwvtyj"
  password = "6ny2svkLvXP3CmtSF15bm01DCSlpLiJB"
  dbname   = "lwdwvtyj"
)

var db *sql.DB

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

// TODO HASH Password
func CreateUser(username string, password string) error {
	if _, err := db.Query("insert into users (username, password) values ($1, $2)", username, password); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		return err
	}
	return nil
}

// TODO HASH Password
func GetUser(username string) (util.User, error) {
	storedCreds := util.User{}
	storedCreds.Email = username
	result := db.QueryRow("select password from users where username=$1", username);
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		return storedCreds, err
	}
	return storedCreds, err
}

func CreateOrGetUser(username string) (util.User, error) {
	var User util.User
	User, err := GetUser(username)
	if err != nil {
		// if no rows in set create user
		if err == sql.ErrNoRows {
			// create new user without password?
			if _, err := db.Query("insert into users (username) values ($1)", username); err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				var emptyUser util.User
				return emptyUser, err
			}
		}
	}
	return User, err
}

func UpdateUser(username string, description string, phone string, job_title string) (int64, error){
	results, err := db.Exec("update users set description=($2),phone=($3),job_title=($4) where username=($1)",
		username, description, phone, job_title)
	rowsImpacted, err := results.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	return rowsImpacted,err
}

func GetAllUsers() ([]util.User, error) {
	rows, err := db.Query(`select username, 
		COALESCE(description, 'BLANK'),
		COALESCE(phone, 'BLANK'),
		COALESCE(job_title, 'BLANK') 
	from users`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	AllUsers := []util.User{}
	for rows.Next() {
		user := util.User{}
		err := rows.Scan(&user.Email, &user.Description, &user.Phone, &user.JobTitle)
		AllUsers = append(AllUsers,user)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return AllUsers, err
}

