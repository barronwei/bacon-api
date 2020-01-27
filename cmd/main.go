package main

import (
	"bacon-api/config"
	"bacon-api/logger"
	"bacon-api/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// delay for shutdown
const (
	delay = 5 * time.Second
)

func main() {
	c, err := config.NewConfig()

	if err != nil {
		log.Fatal("Configuring failed!")
	}

	r := mux.NewRouter()

	setup(r)

	s := http.Server{
		Addr:    ":" + c.Port,
		Handler: r,
	}

	go start(&s)

	block(&s)
}

func setup(r *mux.Router) {
	for _, p := range routes.Router {
		method := logger.Logger(p.Func, p.Name)
		r.
			Name(p.Name).
			Path(p.Path).
			Methods(p.HTTP).
			Handler(method)
	}
}

func start(s *http.Server) {
	log.Println("Listening")
	err := s.ListenAndServe()

	if err != nil {
		log.Fatal("Listening failed!")
	}
}

func block(s *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-ch

	back, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	s.Shutdown(back)
	os.Exit(0)
}
