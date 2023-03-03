package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//we pick the file csv
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	flag.Parse()
	

	//we gonna read this file 
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided csv file.")
	}
	problems := parseLines(lines)
	
	//checking if correct. and if so we increment in the end 
	correct := 0
	for i, p := range problems{
		//this part prints a problem and checks of a correct answer
		//we could use stdin or stdout the standara input streams instead of scanf
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("correct!")
			correct++
		} 
	}

	fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}

func parseLines( lines [][]string) []problem {
	ret := make([]problem, len(lines))
		for i, line := range lines {
			ret[i] = problem{
				q: line[0],
				//this trim space is going to check if csv is kind of destroyed so that it can trim and parse the correct answer
				a: strings.TrimSpace(line[1]),
			}
		}
   return ret
}

type problem struct {
	q string
	a string
}

 func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
 }