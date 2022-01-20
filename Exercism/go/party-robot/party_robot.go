package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tableNo string
	if table < 10 {
		tableNo = "00" + strconv.Itoa(table)
	} else if table < 100 {
		tableNo = "0" + strconv.Itoa(table)
	} else {
		tableNo = strconv.Itoa(table)
	}
	return Welcome(name) + "\n" +
		"You have been assigned to table " + tableNo + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\n" +
		"You will be sitting next to " + neighbor + "."

}
