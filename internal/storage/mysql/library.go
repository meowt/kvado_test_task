package mysql

import (
	"database/sql"

	"kvado_test_task/pkg/models"
)

func (s *Storage) GetBooksByAuthor(author string) (books []models.Book, err error) {
	const op = "internal.storage.mysql.library"

	// Selecting books by author's name
	rows, err := s.db.Query(
		"SELECT books.id, books.name "+
			"FROM books "+
			"JOIN book_author_relations rel "+
			"ON books.id = rel.book_id "+
			"WHERE rel.author_id = (SELECT id FROM authors WHERE authors.name = ?);", author)
	if err != nil && err != sql.ErrNoRows {
		s.log.Printf("%v:%v\n", op, err)
		return
	}

	s.log.Printf("GetBooksByAuthor(\"%v\") queried", author)

	// Scanning data
	for rows.Next() {
		tempBook := models.Book{}
		err = rows.Scan(&tempBook.Id, &tempBook.Name)
		if err != nil {
			s.log.Printf("%v:%v\n", op, err)
			return
		}

		books = append(books, tempBook)
	}
	return
}

func (s *Storage) GetAuthorsByBook(book string) (authors []models.Author, err error) {
	const op = "internal.storage.mysql.library"

	// Selecting authors by book's name
	rows, err := s.db.Query(
		"SELECT authors.id, authors.name "+
			"FROM authors "+
			"JOIN book_author_relations rel "+
			"ON authors.id = rel.author_id "+
			"WHERE rel.book_id = (SELECT id FROM books WHERE books.name = ?);", book)
	if err != nil && err != sql.ErrNoRows {
		s.log.Printf("%v:%v\n", op, err)
		return
	}

	s.log.Printf("GetAuthorsByBook(\"%v\") queried", book)

	// Scanning data
	for rows.Next() {
		tempAuthor := models.Author{}
		err = rows.Scan(&tempAuthor.Id, &tempAuthor.Name)
		if err != nil {
			s.log.Printf("%v:%v\n", op, err)
			return
		}

		authors = append(authors, tempAuthor)
	}
	return
}
