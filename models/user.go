package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/hilbertpurba/PBP/tugas-crud-echo/db"
	"github.com/hilbertpurba/PBP/tugas-crud-echo/helpers"
)

type User struct {
	Id      int    `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	Age     int    `form:"age" json:"age"`
	Address string `form:"address" json:"address"`
}

func GetAllUsers() (Response, error) {
	var user User
	var users []User
	var res Response

	con := db.Connect()

	sqlStatement := "SELECT id, name, age, address from users"
	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)
		if err != nil {
			return res, err
		}
		users = append(users, user)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = users

	return res, nil
}

func InsertNewUser(name string, age int, address string) (Response, error) {
	var res Response

	con := db.Connect()

	sqlStatement := "INSERT INTO users (name, age, address) VALUES (?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, age, address)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"lastInsertedId": lastInsertedId,
	}

	return res, nil
}

func UpdateUser(id int, name string, age int, address string) (Response, error) {
	var res Response

	con := db.Connect()
	sqlStatement := "UPDATE users SET name = ?, age = ?, address = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, age, address, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.Connect()
	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}
	return res, nil
}

func Login(email, password string) (bool, error) {
	var user User
	var pwd string

	con := db.Connect()
	sqlStatement := "SELECT * from users WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&user.Id, &user.Name, &user.Age, &user.Address, &email, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckHashedPassword(password, pwd)
	if !match {
		fmt.Println("Hash and password does not match.")
		return false, err
	}

	return true, nil
}
