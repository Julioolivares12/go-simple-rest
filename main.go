package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//struct de personas
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//struct de direcciones
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//devuelve la lista de personas en formato json
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//devuelve un objeto json de una persona
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	json.NewEncoder(w).Encode(&Person{})

}
func CreatePeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}
func DeletePeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	for position, it := range people {
		if it.ID == params["id"] {
			people = append(people[:position], people[position+1:]...)
			json.NewEncoder(w).Encode("Eliminated")
		}
		break
	}
	json.NewEncoder(w).Encode("not found")
}
func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Ryan", Lastname: "Ray", Address: &Address{
		City: "Dubling", State: "California"}})
	people = append(people, Person{ID: "2", FirstName: "Joe", Lastname: "McMillan"})
	//endpoinsts
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
