package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ayush0x00/cyoa"
)

func handleError(msg error){
	fmt.Printf("Error while opening the file\n")
	log.Fatalf("Error while opening %s\n",msg)
}

func main(){
	port:= flag.Int("port", 3000, "The port number to start server on")
	fileName:= flag.String("file","story.json","The CYOA story in JSON format")
	flag.Parse()

	f,errFopen:= os.Open(*fileName)
	if errFopen!=nil{
		handleError(errFopen)
	}

	story,err:= cyoa.JsonDecoder(f)	
	if err!=nil{
		handleError(err)
	}

	h:= cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n",*port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port),h))
	
}