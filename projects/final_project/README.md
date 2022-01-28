# GO TRAINING FINAL PROJECT: BOOKS DATABASE

# NAME

final_project

# SYNOPSIS

http://localhost:8080/books

# DESCRIPTION

The webapp implements a REST API for a database of book records. All of the data is represented in json format.

| HTTP Verb | Entire collection /contacts  | Specific item /contacts/{id} |
|-----------|------------|----------------|
| POST      | 201 (Created), creates new book record; 409 (Conflict), retrieves current record | 405 (Not allowed) |
| GET       | 200 (OK), retrieves all book records | 200 (OK), retrieves book record; 404 (Not found) |
| PUT      | 405 (Not allowed) | 200 (OK), updates book record; 404 (Not found) |
| DELETE   | 405 (Not allowed) | 200 (OK), removes book record; 404 (Not found) | 

*BOOK RECORDS DATABASE*

| Field | Data Type | Description |
|-------|-----------|-------------|
| ID    | string    | Record ID   |
| Title | string    | Book title   |
| Author | string    | Book author   |
| BookType | string    | Type of book (ie. novel)   |
| Genre | string    | Book genre   |
| Rating | string    | Book rating, out of 10   |
| Status | string    | Status of completion   |
| Notes | string    | Additional notes   |

*URL QUERY OPTIONS*
| Query | Description |
|-------|-------------|
| signup | input a new username and password and create new personal database for user |
| login | input existing username and password to access personal database for user |
| logout | logout of existing username database |

# ERROR HANDLING

Always report errors encountered (e.g., incorrect url)