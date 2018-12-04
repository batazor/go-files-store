package files

import (
	"fmt"
	"github.com/batazor/go-files-store/pkg/minio"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	logger *zap.Logger
	err    error
)

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getFileList)
	r.Get("/{fileId}", getFile)
	r.Post("/", create)
	r.Delete("/{fileId}", destroy)

	return r
}

func getFileList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	doneCh := make(chan minio.File)

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	minio.GetFileList <- minio.File{
		Name:   "",
		FileCH: doneCh,
	}

	select {
	case file := <-doneCh:
		println(file.Name)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"fileList":"` + file.Name + `"}`))
		if err != nil {
			logger.Error(err.Error())
		}
	case <-time.After(2 * time.Second):
		logger.Error("getFileList - timeout > 30 seconds")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("{}"))
		if err != nil {
			logger.Error(err.Error())
		}
	}

}

func getFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{}"))
	if err != nil {
		logger.Error(err.Error())
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("file", err)
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(`{"error":"badRequest"}`))
		if err != nil {
			logger.Error(err.Error())
		}
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		_, _ = w.Write([]byte(`{"error":"badRequest"}`))
		logger.Error(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		logger.Error(err.Error())
	}

	minio.SendFile <- minio.File{
		Name:    handler.Filename,
		Payload: fileBytes,
	}
}

func destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{}"))
	if err != nil {
		logger.Error(err.Error())
	}
}
