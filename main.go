package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("подключение субд")

	db, err := sql.Open("mysql", "root:croftsky1@/mydbtest")
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	insert, err := db.Query("INSERT INTO users (name, age, email) VALUES('golang', 12, 'go@list.com')")
	if err != nil {
		panic(err)
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			panic(err)
		}
	}(insert)

	fmt.Println("Connected")
}
