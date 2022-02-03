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
	Username        string
	Password        string
	Notifications   []string
	Recommendations []BookRecord
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
	groups      []Group
	groupBooks  []BookRecord
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
	var notifs []string
	var recs []BookRecord
	newUser := User{Username: username, Password: password, Notifications: notifs, Recommendations: recs}
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
			if len(user.Notifications) > 0 {
				fmt.Fprintln(w, "Here is what you missed while you've been away.")
				for _, i := range user.Notifications {
					fmt.Fprintln(w, i)
				}
			}
			for key, user := range db.users {
				if user.Username == username {
					db.users[key].Notifications = nil
				}
			}
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
		} else if r.URL.Path == "/books/newgroup" {
			groupName := r.URL.Query().Get("groupname")
			db.newGroup(groupName, w, r)
		} else if r.URL.Path == "/books/addgroupmember" {
			groupName := r.URL.Query().Get("groupname")
			memberName := r.URL.Query().Get("member")
			var members []string
			members = append(members, memberName)
			db.addMembers(members, groupName, w, r)
		} else if r.URL.Path == "/books/addtogroup" {
			groupName := r.URL.Query().Get("groupname")
			title := r.URL.Query().Get("title")
			db.addBookToGroup(title, groupName, w, r)
		} else if r.URL.Path == "/books/viewgroup" {
			groupName := r.URL.Query().Get("groupname")
			db.accessGroup(groupName, w, r)
		} else if r.URL.Path == "/books/recommend" {
			username := r.URL.Query().Get("username")
			bookname := r.URL.Query().Get("bookname")
			db.recommend(username, bookname, w, r)
		} else if r.URL.Path == "/books/checkrecs" {
			username := r.URL.Query().Get("username")
			db.checkRecs(username, w, r)
		} else if r.URL.Path == "/books/addrecs" {
			db.addRecs(w, r)
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

//create groups with other users to share books

type Group struct {
	GroupName     string
	Owner         string
	Members       []string
	BooksExisting []BookRecord
	BooksNew      []BookRecord
}

//newGroup takes in a group name and creates a new json file to store book entries for this new group
func (db *Database) newGroup(groupName string, w http.ResponseWriter, r *http.Request) {
	for _, group := range db.groups {
		if group.GroupName == groupName {
			fmt.Fprintln(w, "Another group with this name already exists. Try another group name.")
			return
		}
	}
	var oldBooks, newBooks []BookRecord
	var memberNames []string
	memberNames = append(memberNames, db.currentUser)
	newGroup := Group{GroupName: groupName, Owner: db.currentUser, Members: memberNames, BooksExisting: oldBooks, BooksNew: newBooks}
	db.groups = append(db.groups, newGroup)
	fname := groupName + "bookclub.json"
	_ = ioutil.WriteFile(fname, nil, 0644)
	fmt.Fprintln(w, "Successfully created new group!")
}

//addMembers takes in a slice of member names and a group name to give members access to an existing group which the current user created
func (db *Database) addMembers(memberNames []string, groupName string, w http.ResponseWriter, r *http.Request) {
	realGroup := false
	for key, group := range db.groups {
		if groupName == group.GroupName {
			realGroup = true
		}
		if groupName == group.GroupName && db.currentUser == group.Owner {
			for _, i := range memberNames {
				db.groups[key].Members = append(group.Members, i)
				for key, memb := range db.users {
					if memb.Username == i {
						db.users[key].Notifications = append(db.users[key].Notifications, "You were added to a new group called: "+groupName)
					}
				}
			}
			fmt.Fprintln(w, "Successfully added new members to your group!")
			return
		}
	}
	if realGroup {
		fmt.Fprintln(w, "You cannot add members to this group. You are not the owner of this group.")
		return
	}
	fmt.Fprintln(w, "This group name does not exist.")
}

//accessGroup takes in a group name and prints the contents of the group database if the current user has access
func (db *Database) accessGroup(groupName string, w http.ResponseWriter, r *http.Request) {
	for _, group := range db.groups {
		if group.GroupName == groupName {
			for _, i := range group.Members {
				if i == string(db.currentUser) {
					fmt.Fprintf(w, "You have successfully accessed %s group!\n", groupName)
					db.openGroupJson(groupName, w, r)
					showGroupBooks(db.groupBooks, w, r)
					return
				}
			}
			fmt.Fprintln(w, "You are not in the members list for this group.")
			return
		}
	}
}

//openGroupJson takes in a group name and opens the associated json file
func (db *Database) openGroupJson(groupName string, w http.ResponseWriter, r *http.Request) {
	fname := groupName + "bookclub.json"
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(file, &db.groupBooks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, "Successfully opened file.")
}

//showGroupBooks prints book records from a slice of book recordes
func showGroupBooks(groupBooks []BookRecord, w http.ResponseWriter, r *http.Request) {
	for _, i := range groupBooks {
		fmt.Fprintln(w, i)
	}
}

//closeGroupJson saves the book entries added into the group into the associated json file
func (db *Database) closeGroupJson(groupName string, w http.ResponseWriter, r *http.Request) {
	fname := groupName + "bookclub.json"
	content, err := json.Marshal(db.groupBooks)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fname, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//addBookToGroup adds a new book to a group given that the group exists and the book exists within the current user's database
func (db *Database) addBookToGroup(bookTitle, groupName string, w http.ResponseWriter, r *http.Request) {
	for _, group := range db.groups {
		if group.GroupName == groupName {
			for _, book := range db.books {
				if book.Title == bookTitle {
					db.accessGroup(groupName, w, r)
					db.groupBooks = append(db.groupBooks, book)
					fmt.Fprintf(w, "Successfully added %s to your group\n", bookTitle)
					db.closeGroupJson(groupName, w, r)
					return
				}
			}
			fmt.Fprintf(w, "%s does not exist in your library. Add the book to your library before adding to a group.", bookTitle)
			return
		}
	}
	fmt.Fprintln(w, "This group does not exist.")
}

//recommend books to other users directly

//recommend takes in a username and bookname and sends a book entry recommendation to that user as long as the book is already in the current users database
func (db *Database) recommend(username, bookName string, w http.ResponseWriter, r *http.Request) {
	var recBook BookRecord
	if !loggedIn {
		fmt.Fprintln(w, "Unable to recommend. You are not currently logged in.")
		return
	}
	for key, user := range db.users {
		if user.Username == username {
			bookPresence := db.findBookEntry(bookName, w, r)
			if bookPresence {
				recBook = db.getBook(bookName, w, r)
				allRecs := db.postBook(recBook, username, w, r)
				db.users[key].Recommendations = append(db.users[key].Recommendations, allRecs[0])
				for key, memb := range db.users {
					if memb.Username == username {
						db.users[key].Notifications = append(db.users[key].Notifications, "A new book: "+bookName+" was recommended to you by "+db.currentUser)
					}
				}
				return
			}
		}
	}
}

//findBookEntry searches for the book entry in the current users database and returns true if it is found or false if not
func (db *Database) findBookEntry(bookName string, w http.ResponseWriter, r *http.Request) bool {
	for _, book := range db.books {
		if book.Title == bookName {
			return true
		}
	}
	fmt.Fprintf(w, "Could not find %s in your library. Post the book into your library first before recommending.\n", bookName)
	return false
}

//getBook retrieves and returns a book entry given a book name
func (db *Database) getBook(bookName string, w http.ResponseWriter, r *http.Request) BookRecord {
	var b BookRecord
	for _, book := range db.books {
		if book.Title == bookName {
			return book
		}
	}
	return b
}

//postBook posts a book record into the recommendations slice of a user with the given username
func (db *Database) postBook(book BookRecord, username string, w http.ResponseWriter, r *http.Request) []BookRecord {
	var allRecs []BookRecord
	for _, user := range db.users {
		if user.Username == username {
			user.Recommendations = append(user.Recommendations, book)
			fmt.Fprintf(w, "Successfully posted %s to %s's recommendations!\n", string(book.Title), username)
			fmt.Fprintln(w, user.Recommendations)
			allRecs = user.Recommendations
			return allRecs
		}
	}
	fmt.Fprintf(w, "Could not find %s in your library. Post the book into your library first before recommending.\n", book.Title)
	return allRecs
}

//checkRecs allows the user to check what books have been recommended to them
func (db *Database) checkRecs(username string, w http.ResponseWriter, r *http.Request) {
	for _, user := range db.users {
		if user.Username == username {
			for _, book := range user.Recommendations {
				fmt.Fprintf(w, "books: %v:", book)
			}
			return
		}
	}
}

//addRecs lets the current user add all the books recommended to them by other users into their own database
func (db *Database) addRecs(w http.ResponseWriter, r *http.Request) {
	var count int
	for key, user := range db.users {
		if user.Username == db.currentUser {
			for _, book := range db.users[key].Recommendations {
				db.books = append(db.books, book)
				count++
			}
			fmt.Fprintf(w, "Successfully added %v recommendation(s) into your library.", count)
		}
	}
}
