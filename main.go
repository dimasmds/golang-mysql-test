package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// User model for storing data from database
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	DateBirth string `json:"dateBirth"`
}

func main() {
	fmt.Println("Go MySql Tutorial, How To Use It")
	dbe, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-test")
	if err != nil {
		panic(err.Error())
	}
	// deletePeople(dbe, 1)
	getPeople(dbe)
	// insertPeople(dbe, "user", "Dimas", "Saputra", "dimas@pinbuk.id", "2018-08-08")
}

func insertPeople(dbe *sql.DB, table string, firstName string, lastName string, email string, dateBirth string) {

	query := "INSERT INTO " + table + " VALUES(null,'" + firstName + "','" + lastName + "','" + email + "','" + dateBirth + "')"
	insert, err := dbe.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Data Inserted")
}

func getPeople(dbe *sql.DB) {
	query := "SELECT * FROM user"

	get, err := dbe.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for get.Next() {
		var user User

		err = get.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateBirth)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("id : " + strconv.Itoa(user.ID))
		fmt.Println("FirstName : " + user.FirstName)
		fmt.Println("LastName : " + user.LastName)
		fmt.Println("Email : " + user.Email)
		fmt.Println("DateBirth : " + user.DateBirth)
		fmt.Println("---------------------------")
	}
}

func deletePeople(dbe *sql.DB, id int) {
	query := "DELETE FROM user WHERE id = " + strconv.Itoa(id)
	delete, err := dbe.Query(query)
	if err != nil {
		panic(err.Error())
	}

	delete.Close()
	fmt.Println("User Deleted")
}
