package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

	//Selecting book-author pairs
	rows, err := db.Query("SELECT b.name AS Book, a.name AS Author\nFROM books b\nJOIN book_author_relations rel\nON b.id = rel.book_id\nJOIN authors a\nON a.id = rel.author_id;")
	if err != nil {
		log.Panicf("Select error %v", err)
	}

	var res []struct {
		Book   string
		Author string
	}
	for rows.Next() {
		temp := struct {
			Book   string
			Author string
		}{}
		err := rows.Scan(&temp.Book, &temp.Author)
		if err != nil {
			log.Printf("Scanning error: %v\n", err)
			return
		}
		res = append(res, temp)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, r := range res {
		fmt.Fprintf(w, "Book: %v\t Author: %v\n", r.Book, r.Author)
	}
	w.Flush()
}
