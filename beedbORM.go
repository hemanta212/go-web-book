package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

type Userinfo struct {
	Uid        int `PK` // if primary key is not id, you need to label your custom id var with `PK`
	Username   string
	Department string
	Created    time.Time
}

func ormConnect() {
	db, err := sql.Open("mysql", "username:testpass@/go_web_test?charset=utf8")
	checkErr(err)

	orm := beedb.New(db)
	beedb.OnDebug = true

	// Insert
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Department = "Test Dep"
	saveone.Created = time.Now()
	orm.Save(&saveone)
	fmt.Println("Id: ", saveone.Uid)

	// Insert with Map
	add := make(map[string]interface{})
	add["username"] = "My name"
	add["department"] = "My dep"
	add["created"] = "2020-10-20"
	orm.SetTable("userinfo").Insert(add)

	// Update
	saveone.Username = "Updated username"
	saveone.Department = "updated dep"
	// Since saveone has a uid, it infers from that to update not insert
	orm.Save(&saveone)

	addUpdate := make(map[string]interface{})
	addUpdate["username"] = "My updated name"
	orm.SetTable("userinfo").SetPK("uid").Where(2).Update(addUpdate)

	// Querying
	var user Userinfo
	orm.Where(1).Find(&user)
	fmt.Println("User: ", user.Username, "\nDep: ", user.Department, "\nuid: ", user.Uid, "\nCreatedAt: ", user.Created)

	var user2 Userinfo
	orm.Where("username=?", "Updated username").Find(&user2)
	fmt.Println("User: ", user.Username, "\nDep: ", user.Department, "\nuid: ", user.Uid, "\nCreatedAt: ", user.Created)

	orm.Delete(&saveone)
	orm.SetTable("userinfo").Where("uid>?", 1).DeleteRow()
}
