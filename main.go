package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	ID          string
	Title       string
	Author      string
	PublishDate string
	ISBN        string
}

type BookCheckout struct{}

type Book struct{}

type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("Get")
	r.HandleFunc("/", writeBlock).Methods("Post")
	r.HandleFunc("/new", newBook).Methods("Post")

	log.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
