package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"kvado_test_task/internal/grpcServer"
	pb "kvado_test_task/internal/grpcServer/handlers/library"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:9000"

func TestLibrary_GetBooksByAuthor(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLibrarianClient(conn)

	//Test cases
	type fields struct {
		UnimplementedLibrarianServer *pb.UnimplementedLibrarianServer
		GrpcServer                   *grpc.Server
		Storage                      grpcServer.StorageManager
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
			name: "Empty author",
			args: args{
				in: &pb.BooksByAuthorRequest{
					AuthorName: "",
				},
			},
			wantErr: true,
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

	// Contact the server and print out its response.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := c.GetBooksByAuthor(ctx, &pb.BooksByAuthorRequest{AuthorName: tt.args.in.AuthorName})

			if err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("GetBooksByAuthor(\"%v\") error = %v, wantErr %v", tt.args.in.AuthorName, err, tt.wantErr)
					return
				}
				t.Logf("GetBooksByAuthor(\"%v\") error = %v, wantErr %v", tt.args.in.AuthorName, err, tt.wantErr)
			} else {
				t.Logf("GetBooksByAuthor(\"%v\") res = %v", tt.args.in.AuthorName, res.GetBookNames())
			}
		})
	}
}

func TestLibrary_GetAuthorsByBook(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLibrarianClient(conn)

	//Test cases
	type fields struct {
		UnimplementedLibrarianServer *pb.UnimplementedLibrarianServer
		GrpcServer                   *grpc.Server
		Storage                      grpcServer.StorageManager
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
					BookName: "12 chairs",
				},
			},
			wantErr: false,
		},
		{
			name: "No authors",
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

	// Contact the server and print out its response.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := c.GetAuthorsByBook(ctx, &pb.AuthorsByBookRequest{BookName: tt.args.in.BookName})

			if err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAuthorsByBook(\"%v\") error = %v, wantErr %v", tt.args.in.BookName, err, tt.wantErr)
					return
				}
				t.Logf("GetAuthorsByBook(\"%v\") error = %v, wantErr %v", tt.args.in.BookName, err, tt.wantErr)
			} else {
				t.Logf("GetAuthorsByBook(\"%v\") res = %v", tt.args.in.BookName, res.GetAuthorNames())
			}
		})
	}
}
