package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	h "github.com/guiaramos/go-url-shortener/api"
	mr "github.com/guiaramos/go-url-shortener/repository/mongodb"
	rr "github.com/guiaramos/go-url-shortener/repository/redis"
	"github.com/guiaramos/go-url-shortener/shortener"
)

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		port := httpPort()
		fmt.Printf("Listening on port %s\n", port)
		errs <- http.ListenAndServe(port, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		url := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(url)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		url := os.Getenv("MONGO_URL")
		db := os.Getenv("MONGO_DB")
		timeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(url, db, timeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
