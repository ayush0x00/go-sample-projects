package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func startQuiz(questions []* Problem, limit int){
	timer:= time.NewTimer(time.Duration(limit)* time.Second)
	answerChannel:= make(chan string)

	counter:= 0
	var answer string
	for i,question:= range questions{
		fmt.Printf("Question number %d: %s = ",i+1, question.question)
		go func(){
			fmt.Scanf("%s", &answer)
			answerChannel<- answer
		}()
		
		select{
		case <- timer.C:
			fmt.Println("\nTime's up")
			exit(fmt.Sprintf("Your score is %d out of %d", counter, len(questions)))

		case  answer:= <- answerChannel:
			if answer==question.answer{
				counter++
			}
		}
		
	}
	exit(fmt.Sprintf("Your score is %d out of %d", counter, len(questions)))
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func main(){
	csvFileName:= flag.String("csv","problems.csv","problem file in csv format")
	timeLimit:= flag.Int("limit", 20, "time limit in seconds to solve the quiz")
	flag.Parse()

	file:= openFile(csvFileName)
	data:= parseData(file)
	startQuiz(data, *timeLimit)
}