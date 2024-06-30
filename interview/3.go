package main

import (
	"context"
	"fmt"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type User struct {
	Username string
	ID       int
	Email    string
}

func fetchUserInfo(db *pgx.Conn, user string, id int) []User {
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND id=%d", user, id)

	var users []User

	pgxscan.Select(context.Background(), db, &users, query)
	return users
}

func main() {
	connStr := "postgres://username:password@localhost:5432/mydb"
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	fetchUserInfo(db, "testuser", 123)
}
