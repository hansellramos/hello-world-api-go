package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	. "hello-world-api/users"
)

var Users map[int]User

func main() {
	Users = make(map[int]User)
	log.Println("Default users: ", Users)

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser)		.Methods("POST")
	router.HandleFunc("/users", GetUsers)			.Methods("GET")
	router.HandleFunc("/users/{id}", GetUser)		.Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser)	.Methods("DELETE")
	log.Println("Listening on http://localhost:529")
	log.Fatal(http.ListenAndServe(":529", router))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if user, ok := Users[id]; ok {
		delete(Users, id)
		log.Println("User:", user, "removed")
		return
	}

	http.Error(w, "There is no user with given id", http.StatusNotFound)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if user, ok := Users[id]; ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "There is no user with given id", http.StatusNotFound)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0)

	for _, value := range Users {
		userList = append(userList, value)
	}
	json.NewEncoder(w).Encode(userList)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, ok := Users[user.Id]; ok {
		http.Error(w, "Already exists an user with given id", http.StatusConflict)
		return
	}

	Users[user.Id] = user
	log.Println("User:", user, "added")
}
