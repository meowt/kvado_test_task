package grpcServer

import (
	"context"
	"errors"
	"log"
	"testing"

	pb "kvado_test_task/internal/grpcServer/handlers/library"
	"kvado_test_task/internal/grpcServer/mocks"
	"kvado_test_task/pkg/models"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

func TestServer_GetBooksByAuthor_Mocks(t *testing.T) {
	// Test cases
	type fields struct {
		UnimplementedLibrarianServer *pb.UnimplementedLibrarianServer
		GrpcServer                   *grpc.Server
		Storage                      StorageManager
	}
	type args struct {
		ctx context.Context
		in  *pb.BooksByAuthorRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.BooksByAuthorReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				in: &pb.BooksByAuthorRequest{
					AuthorName: "A.S. Pushkin",
				},
			},
			wantErr: false,
		},
		{
			name: "No books",
			args: args{
				in: &pb.BooksByAuthorRequest{
					AuthorName: "Not existing author",
				},
			},
			wantErr: true,
		},
		{
			name: "Internal error",
			args: args{
				in: &pb.BooksByAuthorRequest{
					AuthorName: "internal error",
				},
			},
			wantErr: true,
		},
	}

	// Configuring mocks & server
	storageManager := mocks.NewLibraryManager(t)
	storageManager.
		On("GetBooksByAuthor", mock.AnythingOfType("string")).
		Return(func(author string) ([]models.Book, error) {
			switch author {
			case "Not existing author":
				return []models.Book{}, nil
			case "internal error":
				return nil, errors.New("internal server error")
			default:
				return []models.Book{
					{
						Id:   1,
						Name: "Test book",
					}, {
						Id:   2,
						Name: "Test book2",
					}}, nil
			}
		})
	s := &Server{
		Storage: storageManager,
		Log:     log.Default(),
	}

	// Running test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := s.GetBooksByAuthor(tt.args.ctx, tt.args.in)

			// Processing result
			if err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("GetBooksByAuthor(\"%v\") error = \"%v\", wantErr = %v", tt.args.in.AuthorName, err, tt.wantErr)
					return
				}
				t.Logf("GetBooksByAuthor(\"%v\") error = \"%v\", wantErr = %v", tt.args.in.AuthorName, err, tt.wantErr)
			} else {
				t.Logf("GetBooksByAuthor(\"%v\") res = %v", tt.args.in.AuthorName, res)
			}
		})
	}
}

func TestServer_GetAuthorsByBook_Mocks(t *testing.T) {
	// Test cases
	type fields struct {
		UnimplementedLibrarianServer *pb.UnimplementedLibrarianServer
		GrpcServer                   *grpc.Server
		Storage                      StorageManager
		Log                          *log.Logger
	}
	type args struct {
		ctx context.Context
		in  *pb.AuthorsByBookRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AuthorsByBookReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				in: &pb.AuthorsByBookRequest{
					BookName: "Captain's Daughter",
				},
			},
			wantErr: false,
		},
		{
			name: "No books",
			args: args{
				in: &pb.AuthorsByBookRequest{
					BookName: "Not existing book",
				},
			},
			wantErr: true,
		},
		{
			name: "Internal error",
			args: args{
				in: &pb.AuthorsByBookRequest{
					BookName: "internal error",
				},
			},
			wantErr: true,
		},
	}

	// Configuring mocks & server
	storageManager := mocks.NewLibraryManager(t)
	storageManager.
		On("GetAuthorsByBook", mock.AnythingOfType("string")).
		Return(func(book string) ([]models.Author, error) {
			switch book {
			case "Not existing book":
				return []models.Author{}, nil
			case "internal error":
				return nil, errors.New("internal server error")
			default:
				return []models.Author{
					{
						Id:   1,
						Name: "Test author",
					}, {
						Id:   2,
						Name: "Test author2",
					}}, nil
			}
		})
	s := &Server{
		Storage: storageManager,
		Log:     log.Default(),
	}

	// Running test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := s.GetAuthorsByBook(tt.args.ctx, tt.args.in)

			// Processing result
			if err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("GetBooksByAuthor(\"%v\") error = \"%v\", wantErr = %v", tt.args.in.BookName, err, tt.wantErr)
					return
				}
				t.Logf("GetBooksByAuthor(\"%v\") error = \"%v\", wantErr = %v", tt.args.in.BookName, err, tt.wantErr)
			} else {
				t.Logf("GetBooksByAuthor(\"%v\") res = %v", tt.args.in.BookName, res)
			}
		})
	}
}
