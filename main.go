package main

import (
	"log"
	"net/http"
)


type server struct{}

func (s *server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	
	switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "get req !!}`))
		case "POST":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "post req!!}`))
		case "PUT":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "put req!!}`))
		case "DELETE":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "del req !!}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Not Found !!}`))
	}


}

func main() {
	s := &server{}
	http.HandleFunc("/", s.ServerHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}