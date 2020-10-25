package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", tokenVerifyMiddleware(protectedEndpoint)).Methods("GET")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called sigup"))
	fmt.Println("signup invoked")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked")
}

func tokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("tokenVerifyMiddleware invoked")
	return nil
}
