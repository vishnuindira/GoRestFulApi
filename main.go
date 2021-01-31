package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
// defined type Employee 
type Employee struct {
	Id   string `json:"Id"`
	Name string `json:"name"`
}
// declared a global array 
var Employees []Employee

// to get root page (/)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root page")
	fmt.Println("Root page")
}
// to return all employee records (/emplpoyees)
func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)

}

// to return single employee records (/emplpoyees/{empid})
func returnSingleEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)

	for _, emp := range Employees {
		if emp.Id == key {
			json.NewEncoder(w).Encode(emp)
		}
	}
}

// handlers
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
		Employee{Id: "1032", Name: "Vishnu"},
		Employee{Id: "1033", Name: "Abhi"},
		Employee{Id: "1034", Name: "vijay"},
	}
	handleRequests()
}
