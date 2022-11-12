package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name  string `json:"name"`
	Age   uint16 `json:"age"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("подключение субд")

	db, err := sql.Open("mysql", "root:croftsky1@/mydbtest")
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	res, err := db.Query("SELECT name, age, email FROM users")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age, &user.Email)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d has email: %s", user.Name, user.Age, user.Email))
	}

	//insert, err := db.Query("INSERT INTO users (name, age, email) VALUES('zerg', 12000, 'forthequeen@list.com')")
	//if err != nil {
	//	panic(err)
	//}
	//defer func(insert *sql.Rows) {
	//	err := insert.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(insert)

	fmt.Println("Connected")
}
