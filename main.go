package main

import (
	// "context"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Users struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
}

var FirstName, SecondName string
var Id, Age, Num int

func LookListUser() {
	db, err := sql.Open("sqlite", "file:mydb.dbUsers.sqlite -- DO NOT DELETE!!!")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Your look all list users:")
	time.Sleep(2 * time.Second)

	rows, err := db.Query("SELECT * FROM users", FirstName, SecondName, Age)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u Users
		err = rows.Scan(&u.Id, &u.FirstName, &u.SecondName, &u.Age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("User: %s %s, AGE = %d, ID = %d\n", u.FirstName, u.SecondName, u.Age, u.Id)
		fmt.Println("-----------------------------------------------------------------------")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Press key 1 and enter for return to the main menu!")
	time.Sleep(2 * time.Second)

	fmt.Print(">>  ")

	for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
		fmt.Println("Incorret number!!!")
		fmt.Print(">>  ")
	}
	time.Sleep(2 * time.Second)
	fmt.Println()
	Menu()
}

func AddUser() {
	db, err := sql.Open("sqlite", "file:mydb.dbUsers.sqlite -- DO NOT DELETE!!!")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Your have add new user:")
	time.Sleep(2 * time.Second)

	fmt.Println("Enter name: ")
	fmt.Print(">>  ")
	fmt.Scan(&FirstName)
	fmt.Println("Enter second name: ")
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)
	fmt.Println("Enter age: ")
	fmt.Print(">>  ")
	fmt.Scan(&Age)

	_, err = db.Exec("INSERT INTO users (Firstname, Secondname, Age) VALUES (?,?,?)", FirstName, SecondName, Age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User added successfully!")
	time.Sleep(2 * time.Second)
	fmt.Println("Press key 1 and enter for return to the main menu!")
	time.Sleep(2 * time.Second)

	fmt.Print(">>  ")

	for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
		fmt.Println("Incorret number!!!")
		fmt.Print(">>  ")
	}
	time.Sleep(2 * time.Second)
	fmt.Println()
	Menu()
}

func ChangeUser() {
	db, err := sql.Open("sqlite", "file:mydb.dbUsers.sqlite -- DO NOT DELETE!!!")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Your have change user")
	time.Sleep(2 * time.Second)

	fmt.Println("Select id user:")
	fmt.Scan(&Id)

	fmt.Println("Enter name: ")
	fmt.Print(">>  ")
	fmt.Scan(&FirstName)
	fmt.Println("Enter second name: ")
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)
	fmt.Println("Enter age: ")
	fmt.Print(">>  ")
	fmt.Scan(&Age)

	_, err = db.Exec("UPDATE users SET FirstName=?, SecondName=?, Age=? WHERE Id=?", FirstName, SecondName, Age, Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User changed successfully!")
	time.Sleep(2 * time.Second)
	fmt.Println("Press key 1 and enter for return to the main menu!")
	time.Sleep(2 * time.Second)

	fmt.Print(">>  ")

	for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
		fmt.Println("Incorret number!!!")
		fmt.Print(">>  ")
	}
	time.Sleep(2 * time.Second)
	fmt.Println()
	Menu()
}

func DeleteUser() {
	fmt.Println("Your have delete user")
	fmt.Println("Select id user:")

	db, err := sql.Open("sqlite", "file:mydb.dbUsers.sqlite -- DO NOT DELETE!!!")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Scan(&Id)
	_, err = db.Exec("DELETE FROM users WHERE Id = ?", Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("User with Id = %d, %s %s  removed!\n", Id, FirstName, SecondName)
	time.Sleep(2 * time.Second)
	fmt.Println("Press key 1 and enter for return to the main menu!")
	time.Sleep(2 * time.Second)

	fmt.Print(">>  ")

	for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
		fmt.Println("Incorret number!!!")
		fmt.Print(">>  ")
	}
	time.Sleep(2 * time.Second)
	fmt.Println()
	Menu()
}

func Menu() {
	fmt.Println("1. Look list users")
	fmt.Println("2. Add user")
	fmt.Println("3. Delete user")
	fmt.Println("4. Change user")
	fmt.Println("5. Exit app")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Num)
	fmt.Println()

	time.Sleep(3 * time.Second)

	switch {
	case Num == 1:
		LookListUser()
	case Num == 2:
		AddUser()
	case Num == 3:
		DeleteUser()
	case Num == 4:
		ChangeUser()
	case Num == 5:
		fmt.Println("Good bye!!!")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	db, err := sql.Open("sqlite", "file:mydb.dbUsers.sqlite -- DO NOT DELETE!!!")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = db.ExecContext(context.Background(), `CREATE TABLE IF NOT EXISTS users(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		FirstName TEXT NOT NULL,
		SecondName TEXT NOT NULL,
		Age INTEGER NOT NULL
			)`,
	)

	fmt.Println("Welcome to app Users!")
	fmt.Println()
	time.Sleep(2 * time.Second)

	fmt.Println("Select item from menu: ")
	fmt.Println()
	time.Sleep(2 * time.Second)

	Menu()
}
