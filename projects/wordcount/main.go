//package main contains functions to read multiple txt files and count the number of unique words in all the files together.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
)

//readFile takes in the name of a txt file and returns the content of that file as a string
func readFile(fname string) string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return string(content)
}

//saveFiles uses file names taken from the command line using os.Args and calls readFile on each file.
//Returns a slice containing the contents of each file given.
func saveFiles() []string {
	var allText []string
	args := os.Args
	for i := 1; i < len(args); i++ {
		allText = append(allText, readFile(args[i]))
	}
	return allText
}

//findWords takes in a string of contents from a file and separates words from numbers and symbols.
//Returns a slice of all the words in the file.
func findWords(text string) []string {
	var wordList []string
nest:
	for ele := 0; ele < len(text); ele++ {
		var word string
		if unicode.IsLetter(rune(text[ele])) {
			counter := ele
			for unicode.IsLetter(rune(text[counter])) {
				if counter+1 == len(text) {
					break nest
				}
				word += string(text[counter])
				counter++
			}
			wordList = append(wordList, word)
			ele = counter
		}
	}
	return wordList
}

//findAllWords calls on findWords for all file contents, lowercases all words and appends each word into one slice to return.
func findAllWords() []string {
	var allWords []string
	allText := saveFiles()
	for i := 0; i < len(allText); i++ {
		wordList := findWords(allText[i])
		allWords = append(allWords, wordList...)
	}
	for j := 0; j < len(allWords); j++ {
		allWords[j] = strings.ToLower(allWords[j])
	}
	return allWords
}

//SafeCounter is a struct used to safely store a map used to count the instances of each word.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

//addToKey adds a new key to the map or increments the value of an existing key by 1.
func (c *SafeCounter) addToKey(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

//getValue returns the total value associated with a given word.
func (c *SafeCounter) getValue(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

//getKey returns the key given a word.
func (c *SafeCounter) getKey(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return key
}

func (c *SafeCounter) getPair(key string, ch chan string) {
	c.mu.Lock()
	ch <- fmt.Sprintln((key), (c.v[key]))
	c.mu.Unlock()
}

//main calls on all functions and prints a count of each unique word in all files.
func main() {
	allWords := findAllWords()
	counter := SafeCounter{v: make(map[string]int)}
	ch := make(chan string)
	for i := 0; i < len(allWords); i++ {
		go counter.addToKey(allWords[i])
	}
	for key := range counter.v {
		go counter.getPair(key, ch)
		fmt.Printf(<-ch)
	}
}
