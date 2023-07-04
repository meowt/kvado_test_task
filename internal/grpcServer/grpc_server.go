package grpcServer

import (
	"log"
	"net"

	pb "kvado_test_task/internal/grpcServer/handlers/library"
	"kvado_test_task/internal/storage/mysql"
	"kvado_test_task/pkg/models"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//go:generate go run github.com/vektra/mockery/v2@v2.30.1 --name=LibraryManager
type LibraryManager interface {
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetAuthorsByBook(book string) ([]models.Author, error)
}

type StorageManager interface {
	LibraryManager
}

type Server struct {
	*pb.UnimplementedLibrarianServer
	GrpcServer *grpc.Server
	Storage    StorageManager
	Log        *log.Logger
}

// Start gRPC server with mysql and configured logger
func Start(storage *mysql.Storage, log *log.Logger) {
	lis, err := net.Listen("tcp", viper.GetString("server.address"))
	if err != nil {
		log.Fatalf("Falied to listen on %v", lis.Addr())
	}

	server := &Server{
		GrpcServer: grpc.NewServer(),
		Storage:    storage,
		Log:        log,
	}
	pb.RegisterLibrarianServer(server.GrpcServer, server)

	log.Printf("Server listening at %v", lis.Addr())

	//Serving gRPC server
	if err = server.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("Falied to serve gRPC server over %v", lis.Addr())
	}
}
