package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id   int    `json:"id,string,omitempty"`
	Name string `json:"name"`
	City string `json:"city"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage ")
	fmt.Fprintf(w, "Welcome to the HomePage.\n")
}

func getAllEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ExistingEmployee ")

	// Connecting to mySQL Database...
	db := dbConn()

	query := `
		select * from employee;
	`
	rows, err := db.Query(query)

	employee := Employee{}
	res := []Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			panic(err.Error())
		}
		//	fmt.Println("ID: ", Employee.Id, " Name: ", Employee.Name)
		res = append(res, employee)
	}

	// Closing DB connnection...
	defer db.Close()

	fmt.Fprintf(w, "Welcome to the Employee.\n")
	json.NewEncoder(w).Encode(res)
}

func getSingleEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SpecificEmployee...")
	vars := mux.Vars(r)
	key := vars["id"]

	// Connecting to mySQL Database...
	db := dbConn()

	query := "select * from employee where id=" + key + ";"
	rows, err := db.Query(query)

	employee := Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			panic(err.Error())
		}
	}

	// Closing DB connnection...
	defer db.Close()

	fmt.Fprintf(w, "Welcome to the SpecificEmployee.\n")
	json.NewEncoder(w).Encode(employee)
}

func addNewEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addNewEmployee ")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee Employee
	json.Unmarshal(reqBody, &employee)

	fmt.Println(employee)
	dbInsert(employee)
	fmt.Fprintf(w, "Welcome to the addNewEmployee.\n")
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateEmployee ")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee Employee
	json.Unmarshal(reqBody, &employee)
	dbUpdate(employee)
	fmt.Fprintf(w, "Welcome to the updateEmployee.\n")
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteEmployee ")
	//reqBody, _ := ioutil.ReadAll(r.Body)
	id := mux.Vars(r)["id"]

	i, err := strconv.Atoi(id)
	errorCheck(err)
	n_rows_affected := dbDelete(int(i))

	fmt.Fprintf(w, "Welcome to the deleteEmployee.\n")
	fmt.Fprintf(w, fmt.Sprint("Number of rows affected: ", n_rows_affected))
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// HomePage
	myRouter.HandleFunc("/", homePage)

	// Show all Employees
	myRouter.HandleFunc("/show", getAllEmployee)

	// Show Employees with given Id
	myRouter.HandleFunc("/show/{id}", getSingleEmployee)

	// Add Employee
	myRouter.HandleFunc("/add", addNewEmployee).Methods("POST")

	// Update
	myRouter.HandleFunc("/update", updateEmployee).Methods("PUT")

	// Delete
	myRouter.HandleFunc("/delete/{id}", deleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4567", myRouter))
}

func main() {
	fmt.Println("API mySQL and Golang...")
	handleRequest()
}
