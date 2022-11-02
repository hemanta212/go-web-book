package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func sqlLogin() {
	db, err := sql.Open("mysql", "pykancha:testpass@/go_web_test?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, department=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("pykancha", "nepali", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("UPDATE userinfo SET username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("hemu", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// Query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// Delete
	stmt, err = db.Prepare("DELETE from userinfo WHERE uid=?")
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
