package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ayush0x00/cyoa"
)

func main(){
	fileName:= flag.String("file","story.json","The CYOA story in JSON format")
	flag.Parse()

	f,errFopen:= os.Open(*fileName)
	if errFopen!=nil{
		fmt.Printf("Error while opening the file %s\n", *fileName)
		log.Fatalf("Error while opening %s",errFopen)
	}

	story,err:= cyoa.JsonDecoder(f)	
	if err!=nil{
		fmt.Println("Error while decoding the file")
		log.Fatalf("Error while decoding: %s", err)
	}

	fmt.Println(story)
}