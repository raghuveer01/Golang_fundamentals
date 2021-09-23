package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Patna@123@tcp(127.0.0.1:3306)/demo_sql")
	errorCheck(err)
	return db
}

func dbInsert(employee Employee) {
	db := dbConn()
	stmt, err := db.Prepare("INSERT INTO employee(id, name, city) VALUES(?, ?, ?)")
	errorCheck(err)

	res, err := stmt.Exec(employee.Id, employee.Name, employee.City)
	errorCheck(err)

	_, err = res.LastInsertId()
	errorCheck(err)

	defer db.Close()
	defer stmt.Close()
}

func dbDelete(id int) int64 {
	db := dbConn()
	stmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	errorCheck(err)

	res, err := stmt.Exec(id)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	fmt.Println("Number of rows affected: ", n_rows_affected)

	defer db.Close()
	defer stmt.Close()
	return n_rows_affected
}

func dbUpdate(employee Employee) {
	db := dbConn()
	stmt, err := db.Prepare("UPDATE employee SET name=? , city=? WHERE id=?")
	errorCheck(err)

	res, err := stmt.Exec(employee.Id, employee.Name, employee.City)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	fmt.Println("Number of rows affected: ", n_rows_affected)

	defer db.Close()
	defer stmt.Close()
}
