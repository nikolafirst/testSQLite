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
	Date       time.Time
}

var FirstName, SecondName, Confirm string

var Id, Age, Num int

var search = []Users{} // not removed

func LookListUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	fmt.Println("************************  Users list  ****************************")
	fmt.Println()

	rows, err := db.Query("SELECT * FROM users", FirstName, SecondName, Age)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u Users
		err = rows.Scan(&u.Id, &u.FirstName, &u.SecondName, &u.Age, &u.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("User: %s %s, %d year(s), ID: %d; Date created: %s\n", u.FirstName, u.SecondName, u.Age, u.Id, u.Date)
		fmt.Println("-----------------------------------------------------------------------")
	}
	fmt.Println()
	fmt.Println("For return to the main menu press 'y' and enter")
	fmt.Println()
	fmt.Print(">>  ")

	for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
		fmt.Println("Incorret enter!!!")
		fmt.Print(">>  ")
	}
	fmt.Println()
	Menu()
}

func AddUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Added new user:")
	fmt.Println()
	fmt.Println("Enter name: ")
	fmt.Print(">>  ")
	fmt.Scan(&FirstName)
	fmt.Println()
	fmt.Println("Enter second name: ")
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)
	fmt.Println()
	fmt.Println("Enter age: ")
	fmt.Print(">>  ")
	fmt.Scan(&Age)
	fmt.Println()

	_, err = db.Exec("INSERT INTO users (Firstname, Secondname, Age, Date) VALUES (?,?,?,(datetime('now','localtime')))", FirstName, SecondName, Age)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("User added successfully!")
	fmt.Println()
	fmt.Println("Add more user? Press 'y' and enter, or 'n' exit to main menu")
	fmt.Print(">>  ")

	for {
		fmt.Scan(&Confirm)
		fmt.Println()
		if Confirm == "y" {
			time.Sleep(1 * time.Second)
			// fmt.Println()
			AddUser()
			break
		} else if Confirm == "n" {
			fmt.Println("Return to the main menu")
			fmt.Println()
			time.Sleep(1 * time.Second)
			Menu()
			break
		} else {
			fmt.Println("Incorret enter!!!")
			fmt.Println()
			fmt.Println("Select 'y' or 'n' and press enter")
			fmt.Println()
			fmt.Print(">>  ")
		}
	}
}

func SearchUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	fmt.Println("The search is carried out by last name, enter at least 3 characters")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)

	for len(SecondName) < 3 {
		fmt.Println()
		fmt.Println("Last name must be at least 3 characters long, please retry enter:")
		fmt.Println()
		fmt.Print(">>  ")
		fmt.Scan(&SecondName)
	}
	// поиск объекта в базе данных sqlite по первым четырем символам
	query := "SELECT Id, FirstName, SecondName, Age FROM users WHERE SecondName LIKE ?"
	rows, err := db.Query(query, SecondName+"%")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	search = []Users{} // not removed
	found := false
	for rows.Next() {
		var u Users
		err = rows.Scan(&u.Id, &u.FirstName, &u.SecondName, &u.Age)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		search = append(search, u) // not removed
		found = true
		fmt.Println()
		fmt.Printf("User: %s %s, %d year(s), ID: %d\n", u.SecondName, u.FirstName, u.Age, u.Id)
		fmt.Println("------------------------------------------------------------------------")
	}
	fmt.Println()

	if !found {
		fmt.Println()
		fmt.Println("No users found with the last name:", SecondName)
		fmt.Println()
	}

	fmt.Println("For return to the main menu press 'y' and enter")
	fmt.Println()
	fmt.Print(">>  ")

	for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
		fmt.Println("Incorret enter!!!")
		fmt.Print(">>  ")
	}
	fmt.Println()
	Menu()
}

func ChangeUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Your have change user")
	time.Sleep(1 * time.Second)

	fmt.Println("Select id user:")
	fmt.Println()
	fmt.Print(">>  ")
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
	fmt.Println()
	fmt.Println("Press key 1 and enter for return to the main menu!")
	fmt.Println()

	fmt.Print(">>  ")

	for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
		fmt.Println("Incorret number!!!")
		fmt.Print(">>  ")
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
	Menu()
}

func DeleteUser() {
	// var confirm string
	fmt.Println("Your have remove user?(confirm action Y/N)")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Confirm)

	if Confirm == "y" {
		fmt.Println("Select id user:")
		fmt.Println()
		fmt.Print(">>  ")

		db, err := sql.Open("sqlite", "DB.test.sqlite")
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer db.Close()

		fmt.Scan(&Id)
		fmt.Println()

		_, err = db.Exec("DELETE FROM users WHERE Id = ?", Id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("User with Id: %d, removed!\n", Id)
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Println("Press key 1 and enter for return to the main menu!")
		time.Sleep(1 * time.Second)
		fmt.Println()

		fmt.Print(">>  ")

		for fmt.Scan(&Num); Num != 1; fmt.Scan(&Num) {
			fmt.Println("Incorret number!!!")
			fmt.Print(">>  ")
		}
		time.Sleep(1 * time.Second)
		fmt.Println()
		Menu()
	} else if Confirm == "n" {
		fmt.Println()
		fmt.Println("Return to the main menu")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Menu()
	} else {
		fmt.Println()
		fmt.Println("Incorret symbol!!!")
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Println("Return to the main menu")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Menu()
	}
}

func Menu() {
	fmt.Println("1. Look users list")
	fmt.Println("2. Add user")
	fmt.Println("3. Search user")
	fmt.Println("4. Delete user")
	fmt.Println("5. Change user")
	fmt.Println("6. Exit app")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Num)
	fmt.Println()

	switch {
	case Num == 1:
		LookListUser()
	case Num == 2:
		AddUser()
	case Num == 3:
		SearchUser()
	case Num == 4:
		DeleteUser()
	case Num == 5:
		ChangeUser()
	case Num == 6:
		fmt.Println()
		fmt.Println("Exit app?? Press 'y' and enter")
		fmt.Println()
		fmt.Print(">>  ")

		for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
			fmt.Println("Incorret enter!!!")
			fmt.Print(">>  ")
		}
		fmt.Println()
		fmt.Println("Exit app!!!")
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = db.ExecContext(context.Background(), `CREATE TABLE IF NOT EXISTS users(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		FirstName TEXT NOT NULL,
		SecondName TEXT NOT NULL,
		Age INTEGER NOT NULL,
		Date TIMESTAMP
			)`,
	)

	fmt.Println("*************************  Welcome to app  *****************************")
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("********************** Select item from menu  **************************")
	fmt.Println()
	time.Sleep(1 * time.Second)

	Menu()
}
