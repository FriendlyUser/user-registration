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
	if _, err := db.Query("insert into users values ($1, $2)", username, password); err != nil {
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
