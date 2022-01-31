//Patrick Jaime Simba

//package main creates functions to read a csv file, ask a user questions to be answered and print the total correct questions the user answered.
package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//sliceFormats checks if there are completely empty spaces in the csv file. Returns the initial list and an error if either the first or second index on a line is empty
func sliceFormat(ls [][]string) ([][]string, error) {
	for _, item := range ls {
		if strings.TrimSpace(item[0]) == "" || strings.TrimSpace(item[1]) == "" {
			return ls, errors.New("Incorrect database format.")
		}
	}
	return ls, nil
}

//readFile takes in the name of a csv file as a string and returns a slice of the file's contents if no errors are picked up.
func readFile(fname string) [][]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("failed to open csv file: %s", err)
	} else {
		fmt.Println("Successfully opened csv file!")
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed to parse csv file: %s", err)
	} else {
		fmt.Println("Successfully parsed csv file!")
	}
	lines, err1 := sliceFormat(lines)
	if err1 != nil {
		log.Fatalf("Incorrect file format: %s", err1)
	}
	return lines
}

//shuffleSlice takes in a csv file in the form of a 2 dimensional slice of strings and randomly shuffles the original slice.
func shuffleSlice(lines [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
}

//checkQuestionCount takes in a 2 dimensional slice of strings and an integer to check if there are not enough elements in the original slice based on the number of questions the user wants to ask.
//Returns a slice of the first probCount elements of the original slice and the new probCount equal to the length of the slice if probCount > len(lines).
func checkQuestionCount(lines [][]string, probCount int) ([][]string, int) {
	if probCount <= len(lines) {
		lines = lines[:probCount]
	} else {
		fmt.Printf("Insufficient number of questions. Only %v questions will be asked.", len(lines))
		probCount = len(lines)
	}
	return lines, probCount
}

//askQuestions iterates through the elements of lines and prints questions for the user. The user inputs answers which askQuestions will determine to be correct or incorrect.
//Returns the number of correct answers that the user inputted.
func askQuestions(lines [][]string, probCount int) int {
	var correctCount int
	for idx, item := range lines {
		var ans string
		fmt.Printf("Q%v: %v = ", idx+1, strings.TrimSpace(item[0]))
		fmt.Scan(&ans)
		if ans == strings.TrimSpace(item[1]) {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Incorrect!")
		}
	}
	return correctCount
}

//main flags the filename and number of questions to be asked and runs all functions in the package to print the total number of correct answers by the user.
func main() {

	//flags:
	//fname -> name of csv file, default is problems.csv
	//probCount -> number of questions to ask, default is 10
	fname := flag.String("csv", "problems.csv", "problems database")
	probCount := flag.Int("n", 10, "integer")
	flag.Parse()

	lines := readFile(*fname)

	shuffleSlice(lines)

	lines, *probCount = checkQuestionCount(lines, *probCount)

	correctCount := askQuestions(lines, *probCount)

	fmt.Printf("You answered %v out of %v questions correctly.\n", correctCount, *probCount)
}
