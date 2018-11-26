package rest

import (
	"fmt"
	"github.com/batazor/go-files-store/pkg/rest/files"
	"github.com/batazor/go-files-store/pkg/rest/httpLogger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/micro-company/go-auth/utils"
	"go.uber.org/zap"
	"net/http"
)

var (
	logger *zap.Logger
	err    error

	PORT = utils.Getenv("HTTP_PORT", "7070")
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Print("{\"level\":\"error\",\"msg\":\"Error init logger\"}")
	}
}

func Run() {
	// Routes ==================================================================
	r := chi.NewRouter()

	// CORS ====================================================================
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//Debug:            true,
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	//r.Use(middleware.ContentCharset("utf-8", "image/jpeg"))
	r.Use(httpLogger.NewZapMiddleware("router", logger))
	r.NotFound(NotFoundHandler)

	r.Mount("/files", files.Routes())

	logger.Info("Router lisen to port: 4070")
	err := http.ListenAndServe(":"+PORT, r)
	logger.Error(err.Error())
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	logger.Error(`{"error": "not found page"}`)

	_, err := w.Write([]byte(`{"error": "not found page"}`))
	if err != nil {
		logger.Error(err.Error())
	}
}
