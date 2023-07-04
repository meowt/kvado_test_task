package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbConfig = struct {
	user, password, dbname string
}{
	user:     "user",
	password: "secret2",
	dbname:   "library",
}

func main() {
	var (
		authorId, bookId       int64
		isAuthorOld, isBookOld bool
	)
	dbPath := fmt.Sprintf(
		"%v:%v@/%v",
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbname,
	)

	//Opening database connection
	db, err := sql.Open("mysql", dbPath)
	if err != nil {
		log.Panicf("Mysql opening error %v", err)
	}
	defer db.Close()

	//Getting data from cli
	author := flag.String("author", "", "")
	book := flag.String("book", "", "")
	flag.Parse()

	//Inserting author
	res, err := db.Exec("INSERT INTO authors (name) VALUES (?) ", *author)
	if err != nil {
		log.Printf("Insert error: %v", err)
		row := db.QueryRow("SELECT id FROM authors WHERE name = ?", *author)
		err = row.Scan(&authorId)
		if err != nil {
			log.Printf("Scanning error: %v", err)
		}
		isAuthorOld = true
	} else {
		authorId, err = res.LastInsertId()
	}

	//Inserting book
	res, err = db.Exec("INSERT INTO books (name) VALUES (?);", *book)
	if err != nil {
		log.Printf("Insert error: %v", err)
		row := db.QueryRow("SELECT id FROM books WHERE name = ?", *book)
		err = row.Scan(&bookId)
		if err != nil {
			log.Printf("Scanning error: %v", err)
		}
		isBookOld = true
	} else {
		bookId, err = res.LastInsertId()
	}

	if isAuthorOld && isBookOld {
		log.Printf("This pair alredy exists")
		return
	}

	// Creating relation between author&book
	_, err = db.Exec("INSERT INTO book_author_relations (book_id, author_id) VALUES (?, ?);", bookId, authorId)
	if err != nil {
		log.Printf("Insert error: %v", err)
		return
	}
}
