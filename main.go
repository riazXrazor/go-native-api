package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
type Response struct{
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}



func __commonHeaders(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func __sendResponse(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(Response{
		Data : data,
		Success: true,
	})
}


func __checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadData() []User{
	var users []User
	byteValue,_ := ioutil.ReadFile("users.json")
	json.Unmarshal(byteValue, &users)
	return users
}

func WriteData(data []User){
	d, _ := json.Marshal(data)
	err := ioutil.WriteFile("users.json", d, 0644)
    if err != nil {
        panic(err)
    }
}

func getUsers(w http.ResponseWriter, r *http.Request){
	users := ReadData();
	__commonHeaders(w)
	__sendResponse(w, users)

}


func addUser(w http.ResponseWriter, r *http.Request){
	users := ReadData();
	decoder := json.NewDecoder(r.Body)
	var u User
    err := decoder.Decode(&u)
	len := len(users)
	u.Id = len+1
    if err != nil {
        panic(err)
    }
	__commonHeaders(w)
	nusers := append(users, u)
	__sendResponse(w, nusers)
	defer WriteData(nusers)
}

func updateUser(w http.ResponseWriter, r *http.Request){
	users := ReadData();
	decoder := json.NewDecoder(r.Body)
	var u User
    err := decoder.Decode(&u)
	if err != nil {
        panic(err)
    }
	__commonHeaders(w)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for i := range users {
		if users[i].Id == id {
			users[i] = u
		}
	}

	__sendResponse(w, users)
	defer WriteData(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	users := ReadData();
	__commonHeaders(w)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	nusers := make([]User, 0)
	for i := range users {
		if users[i].Id == id {
			nusers = append(users[:i], users[i+1:]...)
		}
	}

	__sendResponse(w, nusers)
	defer WriteData(nusers)
}


func notFound(w http.ResponseWriter, r *http.Request){
	__commonHeaders(w)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "Not Found !!"}`))
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	api.HandleFunc("/users/add", addUser).Methods(http.MethodPost)
	api.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPut)
	api.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)
	
	api.HandleFunc("", notFound)
	
	log.Fatal(http.ListenAndServe(":8080", r))
}