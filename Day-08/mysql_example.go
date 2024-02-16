package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var dbUrl = os.Getenv("MYSQL_DB_URI")

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new user
	createUser(db, "user1", "sample@email.in")

	// Read users
	users := readUsers(db)
	fmt.Println("Users:", users)

	// Update a user
	updateUser(db, 1, "user2", "sample2@email.in")

	// Delete a user
	deleteUser(db, 1)
}

func createUser(db *sql.DB, name string, email string) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created successfully")
}

func readUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users
}

func updateUser(db *sql.DB, id int, name string, email string) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User updated successfully")
}

func deleteUser(db *sql.DB, id int) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully")
}
