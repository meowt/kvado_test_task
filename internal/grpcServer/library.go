package grpcServer

import (
	"context"
	"errors"

	pb "kvado_test_task/internal/grpcServer/handlers/library"
)

func (s *Server) GetBooksByAuthor(ctx context.Context, in *pb.BooksByAuthorRequest) (*pb.BooksByAuthorReply, error) {
	const op = "internal.grpcServer.GetBooksByAuthor"

	// Getting books from storage
	books, err := s.Storage.GetBooksByAuthor(in.AuthorName)
	if err != nil {
		s.Log.Printf("%v:%v\n", op, err)
		return nil, err
	}

	// If no books found, server returns error
	if len(books) == 0 {
		return nil, errors.New("no books found")
	}

	// Extracting only needed information from data
	var bookNames []string
	for _, book := range books {
		bookNames = append(bookNames, book.Name)
	}

	return &pb.BooksByAuthorReply{BookNames: bookNames}, err
}

func (s *Server) GetAuthorsByBook(ctx context.Context, in *pb.AuthorsByBookRequest) (*pb.AuthorsByBookReply, error) {
	const op = "internal.grpcServer.GetAuthorsByBook"

	// Getting books from storage
	authors, err := s.Storage.GetAuthorsByBook(in.BookName)
	if err != nil {
		s.Log.Printf("%v:%v\n", op, err)
		return nil, err
	}

	// If no authors found, server returns error
	if len(authors) == 0 {
		return nil, errors.New("no authors found")
	}

	// Extracting only needed information from data
	var authorNames []string
	for _, author := range authors {
		authorNames = append(authorNames, author.Name)
	}

	return &pb.AuthorsByBookReply{AuthorNames: authorNames}, nil
}
