// Connect to MySQL database with user=root and password=password
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)
func main() {
	db, err := sql.Open("mysql", "root:password@/test?charset=utf8")
	checkErr(err)

	// insert to database
	// User =? to pass argument from Exec()
	stmt, err := db.Prepare("INSERT userinfo SET username=?, department=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("Bunchhieng", "Computer Science", "2016-04-20")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("Bunchhieng Soth", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// Query from the database
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete database
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
