package models

import (
	"arnov17/echo-test/db"
	"arnov17/echo-test/helper"
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"usename"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCont()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
