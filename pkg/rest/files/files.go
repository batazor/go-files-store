package files

import (
	"fmt"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

var (
	logger *zap.Logger
	err    error
)

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Get("/{fileId}", getFile)
	r.Post("/", create)
	r.Delete("/{fileId}", destroy)

	return r
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{}"))
	if err != nil {
		logger.Error(err.Error())
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

	file, _, err := r.FormFile("file")
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

	// TODO: load to fs

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		logger.Error(err.Error())
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