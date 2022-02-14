# GO TRAINING FINAL PROJECT: BOOKS DATABASE

# NAME

final_project

# SYNOPSIS

http://localhost:8080/books

# DESCRIPTION

The webapp implements a REST API for a database of book records. All of the data is represented in json format.

This webapp can be used to store book records for books that you've read, want to read or are currently reading. While the webapp is running, multiple users can sign up to create their own databases and log in or log out to access or or close their databases. The webapp contains features to create, retrieve, update or delete book entries in each users own databases. Moreover, you can view the current book records as a whole or with filters. Additionally, the webapp has functions that allows users to recommend books from their own databases to other signed up users, which those users can add into their own databases. Similarly, users can create groups with other existing accounts in which they can share book entries within a single database.

Webapp is deployed and ran using Docker and includes a command line interface.

# FILES

server\n
\n
  Dockerfile\n
  go.mod\n
  logs.txt - sample logs \n
  main.go - source code for webapp server\n
  server_demo.mkv - demo video for webapp server\n
  test.rest - testing scripts
  
 client
 
  client_demo.mkv - demo video for client side CLI
  main.exe - source code executable
  main.go - source code for CLI

# REQUIREMENTS

1. http API
2. http client (webapp or CLI)
3. Storage - file-based (csv, json) or SQL database
4. Logging capability, multiple users
5. Deployed using Docker
6. Full documentation (README.md, sufficiently commented source code)
7. Video presentation (min 1 minute, max 3 minutes)

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
| newgroup | create a new group database to be shared with other users |
| addgroupmember | add a member to your group from the pool of existing users |
| addtogroup | add a book entry into your group |
| viewgroup | view the contents of a group that you are a part of |
| recommend | recommend a book entry to another existing user from your own database |
| checkrecs | view the recommendations from other users |
| addrecs | add all recommendations from other users into your database |
| viewallbooks | view all books in current user's database |
| filterby... | filter the books in current user's database by author, booktype, rating, status, or genre |

# ERROR HANDLING

Always report errors encountered (e.g., incorrect url)
