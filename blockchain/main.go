package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN string `json:"isbn"`
}

type BookCheckout struct{
	BookId string `json:"book_id"`
	User string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis bool `json:"is_genesis"`
}

type Block struct{
	Pos int
	Timestamp string
	Data BookCheckout
	Hash string
	PrevHash string
}

type Blockchain struct{
	blocks []*Block
}

func newBook(w http.ResponseWriter, r *http.Request){
	var book Book
	if err:=json.NewDecoder(r.Body).Decode(&book); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create the book: %v",err)
		w.Write([] byte("Could bot create new book"))
		return
	}

	h:=md5.New()
	io.WriteString(h,book.ISBN)
	book.ID=fmt.Sprintf("%x",h.Sum(nil))

	resp, err:= json.MarshalIndent(book,""," ")

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not format the book %v", err)
		w.Write([]byte("Could not format the book"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func getBlockChain(w http.ResponseWriter, r *http.Request){

}

func writeBlock(w http.ResponseWriter, r *http.Request){

}

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/",getBlockChain).Methods("GET")
	r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/new",newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000",r))

}