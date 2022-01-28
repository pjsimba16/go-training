//package main implements a REST API where multiple users can create, retrieve, update and delete book records from a JSON database. Each users database will be kept intact as long as the web app is running.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var loggedIn bool = false

type User struct {
	Username string
	Password string
}

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

type Database struct {
	nextID      int
	mu          sync.Mutex
	books       []BookRecord
	users       []User
	currentUser string
}

//main initializes the database and calls on the handler function.
func main() {
	db := &Database{books: []BookRecord{}}
	http.ListenAndServe(":8080", db.handler())
}

//signUp allows user to create a new username and password and creates new json file for that username to save the user's database.
func (db *Database) signUp(username string, password string, w http.ResponseWriter, r *http.Request) {
	for _, user := range db.users {
		if user.Username == username {
			fmt.Fprintln(w, "Unable to sign-up. Username is already taken.")
			return
		}
	}
	newUser := User{Username: username, Password: password}
	db.users = append(db.users, newUser)
	fname := username + ".json"
	_ = ioutil.WriteFile(fname, nil, 0644)
	fmt.Fprintf(w, "Successfully signed up! \nWelcome %s!\n", username)
}

//login allows users to login using a username and password, accesses associated json file, inputs to database.
func (db *Database) login(username string, password string, w http.ResponseWriter, r *http.Request) {
	for _, user := range db.users {
		if loggedIn {
			fmt.Fprintln(w, "You are already logged in. Log out to continue.")
			return
		}
		if user.Username == username && user.Password == password {
			fmt.Fprintf(w, "Successfully logged in! \nHello %s!\n", username)
			db.openJson(username, w, r)
			loggedIn = true
			db.currentUser = username
			return
		}
	}
	fmt.Fprintln(w, "Unable to login. User does not exist in the system. Sign-up to create a new user.")
}

//openJson opens the associated json file with the username and rewrites the current database.
func (db *Database) openJson(username string, w http.ResponseWriter, r *http.Request) {
	fname := username + ".json"
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(file, &db.books)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, "Successfully opened file.")
}

//logout allows a user to logout and save their added/updated/deleted records to their personal database.
func (db *Database) logout(w http.ResponseWriter, r *http.Request) {
	if !loggedIn {
		fmt.Fprintln(w, "Unable to log out. You are not currently logged in.")
		return
	}
	fname := db.currentUser + ".json"
	content, err := json.Marshal(db.books)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fname, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	loggedIn = false
	db.books = nil
	db.nextID = 0
	fmt.Fprintf(w, "User %s was successfully logged out.\n", db.currentUser)
}

//handler calls on the process and processID functions depending on the contents of the URL.
func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/books" {
			fmt.Fprintf(w, "Welcome to your books database!\nSign up and log in with a username and password to access your personal database.\nHappy reading!\n")
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/books/%d", &id); n == 1 {
			db.processID(id, w, r)
		} else if r.URL.Path == "/books/login" {
			username := r.URL.Query().Get("username")
			password := r.URL.Query().Get("password")
			db.login(username, password, w, r)
		} else if r.URL.Path == "/books/signup" {
			username := r.URL.Query().Get("username")
			password := r.URL.Query().Get("password")
			db.signUp(username, password, w, r)
		} else if r.URL.Path == "/books/logout" {
			db.logout(w, r)
		}
	}
}

//process is used when the URL does not contain an ID and contains methods for POST to add a book record, GET to retrieve a book record, PUT and DELETE which print errors
func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var cont BookRecord
		if err := json.NewDecoder(r.Body).Decode(&cont); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if ok, contExist := checkPresence(cont, db.books); ok {
			fmt.Fprintln(w, "409 (Conflict): Book is already in books database.")
			fmt.Fprintln(w, contExist)
			return
		}
		db.mu.Lock()
		cont.ID = db.nextID
		db.nextID++
		db.books = append(db.books, cont)
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"success\": true}")
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.books); err != nil {
			fmt.Fprintln(w, "200 (OK)")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		fmt.Fprintln(w, "405 (Not Allowed)")
		return
	case "DELETE":
		fmt.Fprintln(w, "405 (Not Allowed)")
		return
	}
}

//processID is used when the URL contains an ID and contains methods for GET to retrieve a particular book record, PUT to edit a particular book record, DELETE to delete a book record and POST which does nothing.
func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for _, contact := range db.books {
			if id == contact.ID {
				fmt.Fprintln(w, "200 (OK): Book record retrieved.")
				fmt.Fprintln(w, contact)
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Book record not found.")
		return
	case "PUT":
		ID := r.URL.Query().Get("id")
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")
		bookType := r.URL.Query().Get("type")
		genre := r.URL.Query().Get("genre")
		rating := r.URL.Query().Get("rating")
		status := r.URL.Query().Get("status")
		notes := r.URL.Query().Get("notes")
		for key, book := range db.books {
			if book.ID == id {
				if title != "" {
					db.books[key].Title = title
				}
				if author != "" {
					db.books[key].Author = author
				}
				if bookType != "" {
					db.books[key].BookType = bookType
				}
				if genre != "" {
					db.books[key].Genre = genre
				}
				if rating != "" {
					db.books[key].Rating = rating
				}
				if status != "" {
					db.books[key].Status = status
				}
				if notes != "" {
					db.books[key].Notes = notes
				}
				if ID != "" {
					ID, _ := strconv.Atoi(ID)
					for _, c := range db.books {
						if ID == c.ID {
							fmt.Fprintln(w, "Failed. This ID already exists.")
							return
						}
					}
					db.books[key].ID = ID
				}
				fmt.Fprintln(w, "200 (OK): Book record updated.")
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Book record not found.")
		return
	case "DELETE":
		for key, book := range db.books {
			if id == book.ID {
				db.books = append(db.books[:key], db.books[key+1:]...)
				fmt.Fprintln(w, "200 (OK): Book record deleted.")
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Book record not found.")
		return
	case "POST":
		fmt.Fprintln(w, "405 (Not Allowed)")
		return
	}
}

//checkPresence takes in a BookRecord struct and a slice of BookRecord structs and returns true if the record already exists.
func checkPresence(book BookRecord, books []BookRecord) (bool, BookRecord) {
	result := false
	for _, b := range books {
		if b.Title == book.Title && b.Author == book.Author && b.BookType == book.BookType && b.Genre == book.Genre && b.Rating == book.Rating && b.Status == book.Status && b.Notes == book.Notes {
			result = true
			return result, b
		}
	}
	return result, book
}
