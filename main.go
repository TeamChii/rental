package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/// test ping check connect db

	//err = db.Ping()
	// if err != nil {
	// 	// do something here
	// 	log.Fatal(err)
	// }

	var (
		id   int
		name string
		age  int
	)
	fmt.Scanf("%d", &id)
	rows, err := db.Query("select id, name, age from usertbl where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
