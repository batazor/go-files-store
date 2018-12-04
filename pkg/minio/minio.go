package minio

import (
	"bytes"
	"fmt"
	"github.com/batazor/go-files-store/pkg/utils"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
)

var (
	SendFile = make(chan File)
)

type Client struct {
	client *minio.Client
}

var (
	logger *zap.Logger
	err    error

	MINIO_ENDPOINT   = utils.Getenv("MINIO_ENDPOINT", "localhost:9001")
	MINIO_ACCESS_KEY = utils.Getenv("MINIO_ACCESS_KEY", "ACCESS_KEY")
	MINIO_SECRET_KEY = utils.Getenv("MINIO_SECRET_KEY", "SECRET_KEY")
	MINIO_SECURE     = utils.Getenv("MINIO_SECURE", "false")
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Print("{\"level\":\"error\",\"msg\":\"Error init logger\"}")
	}
}

func Connect() {
	c := Client{}

	// Initialize minio client object.
	c.client, err = minio.New(MINIO_ENDPOINT, MINIO_ACCESS_KEY, MINIO_SECRET_KEY, MINIO_SECURE == "true")
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Success connect to minio")

	go func() {
		for {
			select {
			case file := <-SendFile:
				c.send(file)
			}
		}
	}()
}

func (c *Client) send(file File) {
	reader := bytes.NewReader(file.Payload)
	_, err := c.client.PutObject("test", file.Name, reader, reader.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info(fmt.Sprintf("Uploaded %s of size: %s - successfully", file.Name, reader.Size()))
}
