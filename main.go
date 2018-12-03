package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        int    `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"email"`
	dateBirth string `json:"dateBirth"`
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

		err = get.Scan(&user.id, &user.firstName, &user.lastName, &user.email, &user.dateBirth)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("id : " + strconv.Itoa(user.id))
		fmt.Println("FirstName : " + user.firstName)
		fmt.Println("LastName : " + user.lastName)
		fmt.Println("Email : " + user.email)
		fmt.Println("DateBirth : " + user.dateBirth)
		fmt.Println("---------------------------\n")
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
