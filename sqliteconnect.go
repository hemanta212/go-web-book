package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func sqliteConnect() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// Insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("username", "nepali", "2022-11-02")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// Update
	stmt, err = db.Prepare("UPDATE userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("newuser", id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affected)

	// Query
	rows, err := db.Query("SELECT * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// Delete
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affected, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)
	db.Close()
}
