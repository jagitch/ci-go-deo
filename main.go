package main

import (
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}
func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Printf("server listen on %s", "http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func Add(a int, b int) int {
	return a + b
}
