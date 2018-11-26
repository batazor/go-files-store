package main

import (
	"fmt"
	"github.com/batazor/go-files-store/pkg/minio"
	"github.com/batazor/go-files-store/pkg/rest"
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

	go minio.Connect()
	go rest.Run()

	// Wait forever
	select {}
}
