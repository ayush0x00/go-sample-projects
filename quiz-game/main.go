package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct{
	question string
	answer string
}

func parseData(file *os.File) []*Problem{
	fileData:= csv.NewReader(file)
	lines, err:= fileData.ReadAll()
	
	if err!=nil{
		exit(fmt.Sprintf("Could not parse data %s", err))
	}
	formattedData:= make([]*Problem,len(lines)) 
	for i, p:= range lines{
		formattedData[i] = &Problem{
			question: p[0],
			answer: strings.TrimSpace(p[1]),
		}
	}

	return formattedData
}

func openFile(csvFileName *string) *os.File{
	file, err := os.Open(*csvFileName)

	if err!=nil{
		exit(fmt.Sprintf("Could not open file %s due to error %s\n",*csvFileName,err))
	}
	return file
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func main(){
	csvFileName:= flag.String("csv","problems.csv","problem file in csv format")
	//timeLimit:= flag.Int("limit", 20, "time limit in seconds to solve the quiz")
	flag.Parse()

	file:= openFile(csvFileName)
	data:= parseData(file)
	
}