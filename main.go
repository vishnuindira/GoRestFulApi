package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Designation string `json:"Designation"`
	Project     string `json:"Project"`
	Lead        string `json:"Lead"`
}

var Employee []Employee

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to home page")
	fmt.Println("hit on home page")
}

func returnAllEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Employee")
	json.NewEncoder(w).Encode(Employee)

}
func handleRequest() {
	http.HandleFunc("/articls", returnAllEmployee)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Employee = []Employee{
		Employee{Id: 1032001, Name: "Vishnu", Designation: "Software engineer", Project: "apple", Lead: "Tovino"},
		Employee{Id: 1032002, Name: "mohanlal", Designation: "Software engineer", Project: "apple", Lead: "Tovino"},
	}
	handleRequest()

}
