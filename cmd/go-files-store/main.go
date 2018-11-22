package main

import (
	"fmt"
	"github.com/batazor/go-files-store/pkg/minio"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	err    error
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Print("{\"level\":\"error\",\"msg\":\"Error init logger\"}")
	}
}

func main() {
	logger.Info("Run service")
	fmt.Print("Hello world")

	go minio.Connect()

	// Wait forever
	select {}
}
