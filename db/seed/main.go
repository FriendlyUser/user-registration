package main
// seed db
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// the database is free tier elephantsql, don't intend to ever upgrade
// access if you want, limited to 20MB
const (
  host     = "rajje.db.elephantsql.com"
  port     = 5432
  user     = "lwdwvtyj"
  password = "6ny2svkLvXP3CmtSF15bm01DCSlpLiJB"
  dbname   = "lwdwvtyj"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("create table users (username varchar(255) primary key, password varchar(255), description text, phone varchar(255), job_title varchar(255));")
	if err != nil {
		panic(err)
	}
}
