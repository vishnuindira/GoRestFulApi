package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
func returnSingleEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
}

func handleRequests() {
	// http.HandleFunc("/employees", returnAllEmployees)
	// http.HandleFunc("/", homePage)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	//using  mux for better api usage
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", returnAllEmployees)
	myRouter.HandleFunc("/employees/{id}", returnSingleEmp)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
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
