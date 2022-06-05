package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	parser "github.com/ayush0x00/Parser"
)

func main(){
	file:= flag.String("htmlFile","SampleHTML/test1.html","HTML file to parse link from")
	flag.Parse()

	doc,err:= os.Open(*file)
	if err!=nil{
		fmt.Printf("Error opening file %s\n",*file)
		log.Fatalf("Error: %s",err)
	}

	parsedLink, errorParsing:= parser.Parse(doc)
	if errorParsing!=nil{
		log.Fatalf("Error while parsing link: %s", errorParsing)
	}

	fmt.Println("+v\n",parsedLink)
}