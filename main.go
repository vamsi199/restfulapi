package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)
type Person struct {
ID string `json:"id,omitempty"`
Name string `json:"name,omitempty"`
Address *Address `json:"address,omitempty"`
}
type Address struct {
City string `json:"city,omitempty"`
State string `json:"state,omitempty"`
}
var people []Person
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)

}
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _,item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(Person{})

}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people,person)
	json.NewEncoder(w).Encode(people)

}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index,item := range people {
		if item.ID == params["id"] {
			people = append(people[:index],people[index+1:]...)
			break
		}
	}

}
func main() {
	router := mux.NewRouter()
	people = append(people,Person{ID:"1",Name: "Vamsi" ,Address: &Address{City:"cleveland",State:"ohio"}})
	people = append(people,Person{ID:"2",Name: "sfjhaksuf" })

	router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}",DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",router))

}
