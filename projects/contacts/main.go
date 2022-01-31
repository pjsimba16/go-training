//Patrick Jaime Simba

//package main implements a REST API where you can create, retrieve, update and delete contact records from a JSON database.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Contact struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

type Database struct {
	nextID   int
	mu       sync.Mutex
	contacts []Contact
}

//main initializes the database and calls on the handler function.
func main() {
	db := &Database{contacts: []Contact{}}
	http.ListenAndServe(":8080", db.handler())
}

//handler calls on the process and processID functions depending on the contents of the URL.
func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		}
	}
}

//process is used when the URL does not contain an ID and contains methods for POST to add a contact, GET to retrieve a contact, PUT and DELETE which print errors
func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var cont Contact
		if err := json.NewDecoder(r.Body).Decode(&cont); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if ok, contExist := checkPresence(cont, db.contacts); ok {
			fmt.Fprintln(w, "409 (Conflict): Contact already in contacts database.")
			fmt.Fprintln(w, contExist)
			return
		}
		db.mu.Lock()
		cont.ID = db.nextID
		db.nextID++
		db.contacts = append(db.contacts, cont)
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"success\": true}")
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.contacts); err != nil {
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

//processID is used when the URL contains an ID and contains methods for GET to retrieve a particular contact, PUT to edit a particular contact, DELETE to delete a contact and POST which does nothing.
func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for _, contact := range db.contacts {
			if id == contact.ID {
				fmt.Fprintln(w, "200 (OK): Contact Retrieved.")
				fmt.Fprintln(w, contact)
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Contact not found.")
		return
	case "PUT":
		ID := r.URL.Query().Get("id")
		last := r.URL.Query().Get("last")
		first := r.URL.Query().Get("first")
		company := r.URL.Query().Get("company")
		address := r.URL.Query().Get("address")
		country := r.URL.Query().Get("country")
		position := r.URL.Query().Get("position")
		for key, contact := range db.contacts {
			if contact.ID == id {
				if last != "" {
					db.contacts[key].Last = last
				}
				if first != "" {
					db.contacts[key].First = first
				}
				if company != "" {
					db.contacts[key].Company = company
				}
				if address != "" {
					db.contacts[key].Address = address
				}
				if country != "" {
					db.contacts[key].Country = country
				}
				if position != "" {
					db.contacts[key].Position = position
				}
				if ID != "" {
					ID, _ := strconv.Atoi(ID)
					for _, c := range db.contacts {
						if ID == c.ID {
							fmt.Fprintln(w, "Failed. This ID already exists.")
							return
						}
					}
					db.contacts[key].ID = ID
				}
				fmt.Fprintln(w, "200 (OK): Contact updated.")
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Contact not found.")
		return
	case "DELETE":
		for key, contact := range db.contacts {
			if id == contact.ID {
				db.contacts = append(db.contacts[:key], db.contacts[key+1:]...)
				fmt.Fprintln(w, "200 (OK): Contact deleted.")
				return
			}
		}
		fmt.Fprintln(w, "404 (Not Found): Contact not found.")
		return
	case "POST":
		fmt.Fprintln(w, "405 (Not Allowed)")
		return
	}
}

//checkPresence takes in a Contact struct and a slice of Contact structs and returns true if the contact already exists.
func checkPresence(cont Contact, contacts []Contact) (bool, Contact) {
	result := false
	for _, contact := range contacts {
		if contact.Last == cont.Last && contact.First == cont.First && contact.Address == cont.Address && contact.Country == cont.Country && contact.Position == cont.Position && contact.Company == cont.Company {
			result = true
			return result, contact
		}
	}
	return result, cont
}
