package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
 }

	type Address struct {
		City string `json:"city,omitempty"`
		State string `json:"state,omitempty"`
		
	}
 

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(people)
  }

func GetPerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	fmt.Println(params["id"])

    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{}) 
}

func CreatePerson(w http.ResponseWriter, r *http.Request){
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index, item := range people {
					if item.ID == params["id"] {
									people = append(people[:index], people[index+1:]...)
									break
					}
	}
	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) 

 	for index, item := range people {
					if item.ID == params["id"] {
						var person Person
						_ = json.NewDecoder(r.Body).Decode(&person)

						person.ID = item.ID   
						 people[index] = person
									json.NewEncoder(w).Encode(people[index])
									return
					}
	}
	json.NewEncoder(w).Encode(&Person{})
 
}  

func main() {
	people = append(people, Person{ID: "1", FirstName: "John", LastName: "Doe", Address: &Address{
		City: "Curitiba", State: "Paraná"} })
 	r := mux.NewRouter()
 r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/person/{id}", GetPerson).Methods("GET")
 r.HandleFunc("/person", CreatePerson).Methods("POST")
	r.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE") 
	r.HandleFunc("/person/{id}", UpdatePerson).Methods("PATCH") 

	log.Fatal(http.ListenAndServe(":8000", r))
}

