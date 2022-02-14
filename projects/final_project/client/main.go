//main initializes a client command line interface that works directly with the book listing web app server
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "http://localhost:8080/"

type BookRecord struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	BookType string `json:"booktype"`
	Genre    string `json:"genre"`
	Rating   string `json:"rating"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

//main calls methods to use all functions in the web server script
func main() {
	cmd := flag.String("cmd", "", `signup, login, logout, post, get, view, delete, update, newgroup, addgroupmember, 
	addtogroup, viewgroup, recommend, checkrecs, addrecs, filterbyauthor, filterbybooktype, 
	filterbystatus, filterbygenre, filterbyrating`)
	//bookEntry := flag.Struct("book", "", "book to be added")
	id := flag.Int("id", -1, "ID of record to process")
	flag.Parse()
	switch *cmd {
	case "signup":
		signUp()
	case "login":
		login()
	case "logout":
		logout()
	case "post":
		post()
	case "get":
		get()
	case "view":
		view()
	case "delete":
		delete(*id)
	case "update":
		update(*id)
	case "newgroup":
		newgroup()
	case "addtogroup":
		addtogroup()
	case "viewgroup":
		viewgroup()
	case "addgroupmember":
		addgroupmember()
	case "recommend":
		recommend()
	case "checkrecs":
		checkrecs()
	case "addrecs":
		addrecs()
	case "filterbyauthor":
		filterbyauthor()
	case "filterbybooktype":
		filterbybooktype()
	case "filterbystatus":
		filterbystatus()
	case "filterbygenre":
		filterbygenre()
	case "filterbyrating":
		filterbyrating()
	}
}

//signUp allows a user to sign up with a new account
func signUp() {
	var username, password string
	fmt.Println("Enter username: ")
	fmt.Scan(&username)
	fmt.Println("Enter password: ")
	fmt.Scan(&password)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "signup?username=" + username + "&password=" + password)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//login allows a returning user to login with their credentials
func login() {
	var username, password string
	fmt.Println("Enter username: ")
	fmt.Scan(&username)
	fmt.Println("Enter password: ")
	fmt.Scan(&password)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "login?username=" + username + "&password=" + password)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//logout allows a user to logout of their account
func logout() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "logout")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//post lets a user add book entries into their personal database
func post() {
	inData := createPost()
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	book, err := json.Marshal(inData)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.Post(baseURL+"books", "application/json", bytes.NewBuffer(book))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//createPost asks the user for inputs for each piece of information in the book entry
func createPost() BookRecord {
	var bookSlice []string
	bookInfo := [...]string{"title", "author", "book type (Novel, comic book, etc.)", "genre", "rating (out of 10)", "status", "notes"}
	for i := range bookInfo {
		fmt.Println("Enter " + bookInfo[i])
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		bookSlice = append(bookSlice, input)
	}
	var rec BookRecord
	rec.Title = strings.Replace(bookSlice[0], "\r\n", "", -1)
	rec.Author = strings.Replace(bookSlice[1], "\r\n", "", -1)
	rec.BookType = strings.Replace(bookSlice[2], "\r\n", "", -1)
	rec.Genre = strings.Replace(bookSlice[3], "\r\n", "", -1)
	rec.Rating = strings.Replace(bookSlice[4], "\r\n", "", -1)
	rec.Status = strings.Replace(bookSlice[5], "\r\n", "", -1)
	rec.Notes = strings.Replace(bookSlice[6], "\r\n", "", -1)
	return rec
}

//get calls on the GET method of the server script
func get() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "books")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//view prints all records in a user's personal database
func view() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "viewall")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//delete calls on the DELETE method of the server script and deletes a record based on the input id
func delete(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/%d", baseURL+"books", id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := c.Do(req)
	if err2 != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//update edits a record based on the input id
func update(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/%d", baseURL+"books", id)
	req, err := http.NewRequest("PUT", url+updateHelper(), nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := c.Do(req)
	if err2 != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//updateHelper lets user choose the field to update and enter their updated string
func updateHelper() string {
	var field string
	fmt.Println(`Which field would you like to update? 
	(Choose between: Title, Author, BookType, Genre, Rating, Status, Notes)`)
	fmt.Scan(&field)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your updates to " + field)
	input, _ := inputReader.ReadString('\n')
	return "?" + field + "=" + input
}

//newgroup creates a new group to share book entries with other members
func newgroup() {
	var groupname string
	fmt.Println("Enter group name: ")
	fmt.Scan(&groupname)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "newgroup?groupname=" + groupname)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//addtogroup adds a single book entry into an existing group
func addtogroup() {
	var title, groupname string
	fmt.Println("Enter group name: ")
	fmt.Scan(&groupname)
	fmt.Println("Enter book title from your library: ")
	fmt.Scan(&title)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "addtogroup?groupname=" + groupname + "&title=" + title)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//addgroupmember adds another existing user to your book club
func addgroupmember() {
	var member, groupname string
	fmt.Println("Enter group name: ")
	fmt.Scan(&groupname)
	fmt.Println("Enter username to add to your group: ")
	fmt.Scan(&member)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "addgroupmember?groupname=" + groupname + "&member=" + member)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//viewgroup lets a user see all the entries in an existing group that they are a member of
func viewgroup() {
	var groupname string
	fmt.Println("Enter group name: ")
	fmt.Scan(&groupname)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "viewgroup?groupname=" + groupname)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//recommend lets a user recommend a single book directly to another existing user
func recommend() {
	var title, user string
	fmt.Println("Enter username to recommend to: ")
	fmt.Scan(&user)
	fmt.Println("Enter book title from your library: ")
	fmt.Scan(&title)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "recommend?username=" + user + "&bookname=" + title)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//checkrecs lets a user see all of their current recommendations
func checkrecs() {
	var username string
	fmt.Println("Enter username: ")
	fmt.Scan(&username)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "checkrecs?username=" + username)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//addrecs lets a user add all recommendations into their personal library
func addrecs() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "addrecs")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//filterbyauthor filters books by author and in alphabetical order
func filterbyauthor() {
	var author string
	fmt.Println("Enter author: ")
	fmt.Scan(&author)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "filterbyauthor?author=" + author)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//filterbyauthor filters books by book type and in alphabetical order
func filterbybooktype() {
	var booktype string
	fmt.Println("Enter book type: ")
	fmt.Scan(&booktype)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "filterbybooktype?booktype=" + booktype)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//filterbyauthor filters books by status and in alphabetical order
func filterbystatus() {
	var status string
	fmt.Println("Enter status: ")
	fmt.Scan(&status)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "filterbystatus?status=" + status)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//filterbyauthor filters books by genre and in alphabetical order
func filterbygenre() {
	var genre string
	fmt.Println("Enter genre: ")
	fmt.Scan(&genre)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "filterbygenre?genre=" + genre)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

//filterbyauthor filters books by rating and in alphabetical order
func filterbyrating() {
	var rating string
	fmt.Println("Enter rating: ")
	fmt.Scan(&rating)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL + "filterbyrating?rating=" + rating)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
