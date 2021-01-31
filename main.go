package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Id   int    `json:"Id"`
	Name string `json:"name"`
}

var Employees []Employee

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)

}

func handleRequests() {
	http.HandleFunc("/employees", returnAllEmployees)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Hi")

	Employees = []Employee{
		Employee{Id: 1032, Name: "Vishnu"},
		Employee{Id: 1033, Name: "Abhi"},
		Employee{Id: 1034, Name: "vijay"},
	}
	handleRequests()
}
