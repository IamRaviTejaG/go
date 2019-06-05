package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "me"
	password = "password"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s
		sslmode=disable`, host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i + 1)
		sqlQuery := fmt.Sprintf(`
			INSERT INTO users(name, email) VALUES ('Testing %d', 'test%d@sharklasers.com'),
			('Testing %d', 'test%d@sharklasers.com');
			`, i, i, i+1, i+1)
		_, err = db.Exec(sqlQuery)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
}
