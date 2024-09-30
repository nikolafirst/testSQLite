package main

import (
	// "context"

	"database/sql"
	"fmt"
	"log"
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

var (
	FirstName, SecondName, Confirm string

	Id, Age, Num int

	search []Users // не удалять
)

func lookListUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Ошибка открытия базы данных:", err)
		return
	}
	defer db.Close()

	fmt.Println("************************  Список пользователей  ****************************")
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
			fmt.Println("Ошибка сканирования базы данных:", err)
			return
		}
		fmt.Printf("User: %s %s, %d year(s), ID: %d; Date created: %s\n", u.FirstName, u.SecondName, u.Age, u.Id, u.Date)
		fmt.Println("-----------------------------------------------------------------------")
	}
	fmt.Println()
	fmt.Println("Для возврата в главное меню нажмите" + "у" + "enter")
	fmt.Println()
	fmt.Print(">>  ")

	for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
		fmt.Println("Неправильный ввод!!!")
		fmt.Print(">>  ")
	}
	fmt.Println()
	Menu()
}

func addUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Ошибка открытия базы данных:", err)
		return
	}
	defer db.Close()

	fmt.Println("Внесение в базу данных нового пользователя:")
	fmt.Println()
	fmt.Println("Введите имя: ")
	fmt.Print(">>  ")
	fmt.Scan(&FirstName)
	fmt.Println()
	fmt.Println("Введите фамилию: ")
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)
	fmt.Println()
	fmt.Println("Введите возраст: ")
	fmt.Print(">>  ")
	fmt.Scan(&Age)
	fmt.Println()

	_, err = db.Exec("INSERT INTO users (Firstname, Secondname, Age, Date) VALUES (?,?,?,(datetime('now','localtime')))", FirstName, SecondName, Age)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Пользователь успешно добавлен в базу данных!")
	fmt.Println()
	fmt.Println("Добавить еще пользователя? Нажмите 'y' и enter, или 'n' для выхода в главное меню")
	fmt.Print(">>  ")

	for {
		fmt.Scan(&Confirm)
		fmt.Println()
		if Confirm == "y" {
			time.Sleep(1 * time.Second)
			addUser()
			break
		} else if Confirm == "n" {
			fmt.Println("Возврат в главное меню")
			fmt.Println()
			time.Sleep(1 * time.Second)
			Menu()
			break
		} else {
			fmt.Println("Неправильный ввод")
			fmt.Println()
			fmt.Println("Выбирете 'y' или 'n' и нажмите enter")
			fmt.Println()
			fmt.Print(">>  ")
		}
	}
}

func searchUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Ошибка открытия базы данных:", err)
		return
	}
	defer db.Close()

	fmt.Println("Введите первые три буквы или полностью фамилию:")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)

	for len(SecondName) < 3 {
		fmt.Println()
		fmt.Println("Введено мнее трех символов, пожалуйста повторите ввод:")
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
			fmt.Println("Ошибка сканирования базы данных:", err)
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
		fmt.Println("Данного пользователя не существует: ", SecondName)
		fmt.Println()
	}

	fmt.Println("Повторить поиск? Нажмите 'y' и  enter, или 'n' для перехода в главное меню")
	fmt.Println()
	fmt.Print(">>  ")

	// ConfirmInput()
	for {
		fmt.Scan(&Confirm)
		fmt.Println()
		if Confirm == "y" {
			time.Sleep(1 * time.Second)
			searchUser()
			break
		} else if Confirm == "n" {
			fmt.Println("Возврат в главное меню")
			fmt.Println()
			time.Sleep(1 * time.Second)
			Menu()
			break
		} else {
			fmt.Println("Некорректный ввод!!!")
			fmt.Println()
			fmt.Println("Выбирете 'y' или 'n' и нажмите enter")
			fmt.Println()
			fmt.Print(">>  ")
		}
	}
}

func changeUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	defer db.Close()

	fmt.Println("Вы хотите изменить данные пользователя")
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("Выбирете id пользователя:")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Id)

	fmt.Println("Введите имя: ")
	fmt.Print(">>  ")
	fmt.Scan(&FirstName)
	fmt.Println("Введите фамилию: ")
	fmt.Print(">>  ")
	fmt.Scan(&SecondName)
	fmt.Println("Введите возраст: ")
	fmt.Print(">>  ")
	fmt.Scan(&Age)

	_, err = db.Exec("UPDATE users SET FirstName=?, SecondName=?, Age=? WHERE Id=?", FirstName, SecondName, Age, Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Данные успешно изменены!")
	fmt.Println()
	fmt.Println("Для перехода в главное меню нажмите 'y' и enter")
	fmt.Println()
	fmt.Print(">>  ")

	for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
		fmt.Println("Incorrect enter!!!")
		fmt.Print(">>  ")
	}
	fmt.Println()
	Menu()
}

func removeUser() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println("Ошибка открытия базы данных:", err)
		return
	}
	defer db.Close()

	fmt.Println("Вы хотите удалить пользователя")
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("Выбирете ID пользователя:")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Id)

	result, err := db.Exec("DELETE FROM users WHERE Id=?", Id)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Пользователь с указанным ID не найден")
	}

	fmt.Println("Пользователь успешно удален!")
	fmt.Println()
	fmt.Println("Для возврата в главное меню нажмите 'y' и enter")
	fmt.Println()
	fmt.Print(">>  ")

	for fmt.Scan(&Confirm); Confirm != "y"; fmt.Scan(&Confirm) {
		fmt.Println("Неправильный ввод!!!")
		fmt.Print(">>  ")
	}
	fmt.Println()
	Menu()
}

func Menu() {
	fmt.Println("1. Посмотреть список пользователей")
	fmt.Println("2. Добавить пользователя")
	fmt.Println("3. Поиск пользователя")
	fmt.Println("4. Изменить пользователя")
	fmt.Println("5. Удалить пользователя")
	fmt.Println("6. Выход из программы")
	fmt.Println()
	fmt.Print(">>  ")
	fmt.Scan(&Num)
	fmt.Println()

	switch {
	case Num == 1:
		lookListUser()
	case Num == 2:
		addUser()
	case Num == 3:
		searchUser()
	case Num == 4:
		changeUser()
	case Num == 5:
		removeUser()
	case Num == 6:
		time.Sleep(1 * time.Second)
		fmt.Println("Выход из программы!!!")
		fmt.Println()
	}
}

func main() {
	db, err := sql.Open("sqlite", "DB.test.sqlite")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		FirstName TEXT NOT NULL,
		SecondName TEXT NOT NULL,
		Age INTEGER NOT NULL,
		Date TIMESTAMP
		)`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	fmt.Println("Таблица 'users' успешно создана или уже существует")

	fmt.Println("*************************  Welcome to app  *****************************")
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("********************** Select item from menu  **************************")
	fmt.Println()
	time.Sleep(1 * time.Second)

	Menu()
}
