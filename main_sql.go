package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main_sql() {
	db, err := sql.Open("mysql", "root:Aesaj2025@/golang")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into golang.users (id, name, age) values (?, ?, ?)",
		1, "SERJ", 18)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println("Connected to db...")
}
