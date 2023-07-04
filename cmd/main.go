package main

import (
	"kvado_test_task/internal/config"
	"kvado_test_task/internal/grpcServer"
	"kvado_test_task/internal/logger"
	"kvado_test_task/internal/storage/mysql"

	"github.com/spf13/viper"
)

func main() {
	//Config loading from .yaml file, using github.com/spf13/viper
	config.MustLoad()

	//Logger configuration depending on the environment
	grpcLog, storageLog := logger.Configure(viper.GetString("env"))

	//Database connection&deployment
	storage := mysql.MustSetup(storageLog)

	//gRPC server starting
	grpcServer.Start(storage, grpcLog)
}
