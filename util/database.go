package util

import (
	"database/sql"
	"fmt"
)

const (
	Name = "yukon"
	User = "root"
	Password = ""
)

func StartDatabase() {
	db, err := sql.Open("mysql", User + ":" + Password + "@/" + Name)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	_ = db
	fmt.Println("Opened Database connection")
}