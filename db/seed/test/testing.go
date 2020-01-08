package testing
// seed db
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	util "github.com/FriendlyUser/user-registration/util"
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
		AllUsers = append(AllUsers, user)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(AllUsers)
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	result, err := db.Exec("update users set description=($2),phone=($3),job_title=($4) where username=($1)",
	"studentdavidli@gmail.com", "Huge Gamer", "111", "Software Engineer")
	if err != nil {
		fmt.Println(err)
	}
	result_id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result_id)
	rowsImpacted, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rowsImpacted)
}

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "password=%s dbname=%s sslmode=disable",
//     host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	username := "epic"
// 	storedCreds := util.User{}
// 	storedCreds.Email = username
// 	result := db.QueryRow("select password from users where username=$1", username);
// 	err = result.Scan(&storedCreds.Password)
// 	if err != nil {
// 		// if no rows in set create user
// 		if err == sql.ErrNoRows {
// 			fmt.Printf("Trying to load files")
// 			fmt.Printf(err.Error())
// 		}
// 	}

// }
