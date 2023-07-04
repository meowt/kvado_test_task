package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

// Environmental constants
const (
	envLocal = "local"
	envProd  = "prod"
)

// Configure function
func Configure(env string) (serverLogger, databaseLogger *log.Logger) {
	switch env {
	case envLocal:
		serverLogger = log.New(os.Stdout, "gRPC: ", log.Ldate|log.Ltime|log.Lshortfile)
		databaseLogger = log.New(os.Stdout, "db: ", log.Ldate|log.Ltime|log.Lshortfile)

	case envProd:
		today := time.Now().Format("02-01-06")
		file, err := os.OpenFile(today+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Log's file opening error %v", err)
		}
		serverLogger = log.New(file, "gRPC: ", log.Ldate|log.Ltime|log.Lshortfile)
		serverLogger = log.New(file, "db: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	log.Printf("%v loggers configured", strings.Title(env))
	return
}
