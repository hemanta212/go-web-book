package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "pykancha"
	DB_PASSWORD = "sharmaji"
	DB_NAME     = "go_postgres"
)

func postgresConnect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# inserting values")
	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, department, created) VALUES ($1, $2, $3) returning uid;", "pykancha", "nepali", "2022-10-11").Scan(&lastInsertId)

	fmt.Println("# Updating ")
	stmt, err := db.Prepare("UPDATE userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("hemu", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("Querying")

	rows, err := db.Query("SELECT * from userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}

	fmt.Println("#Deleting")
	stmt, err = db.Prepare("DELETE from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, " Rows affected")
}
