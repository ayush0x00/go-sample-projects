package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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

func createBlock(prevBlock *Block, checkoutItem BookCheckout) *Block{
	block:= &Block{}
	block.Pos=prevBlock.Pos+1
	block.PrevHash=prevBlock.Hash
	block.Timestamp=time.Now().String()
	block.generateHash()


	return block
}

func (block *Block) generateHash(){
	bytes,_ := json.Marshal(block.Data)

	data:= fmt.Sprint(block.Pos)+block.Hash+string(bytes)+block.Timestamp
	
	hash:= sha256.New()
	hash.Write([]byte(data))
	block.Hash= hex.EncodeToString(hash.Sum(nil))
}

func (bc *Blockchain) AddBlock(data BookCheckout){
	prevBlock:= bc.blocks[len(bc.blocks)-1]

	currBlock:= createBlock(prevBlock, data)

	if validBlock(currBlock, prevBlock){
		bc.blocks = append(bc.blocks, currBlock)
	}
}

func validBlock(currBlock, prevBlock *Block) bool{
	if prevBlock.Hash != currBlock.PrevHash{
		return false
	}

	if !currBlock.validateHash(currBlock.Hash){
		return false
	}

	if currBlock.Pos!=prevBlock.Pos+1{
		return false
	}

	return true
}

func (block *Block) validateHash (hash string) bool{
	block.generateHash()
	
	return block.Hash == hash
}

var BlockChain *Blockchain

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
	var checkoutItem BookCheckout
	if err:= json.NewDecoder(r.Body).Decode(&checkoutItem); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not decode data %v",err)
		w.Write([]byte("Could not decode checkout data"))
		return
	}

	BlockChain.AddBlock(checkoutItem)
}

func CreateGenesisBlock() *Block{
	return createBlock(& Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain{
	return &Blockchain{[] *Block{CreateGenesisBlock()}}
}

func main(){

	BlockChain = NewBlockchain()
	r:=mux.NewRouter()
	r.HandleFunc("/",getBlockChain).Methods("GET")
	r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/new",newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000",r))

}