package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	// Name of the service
	Name = "microservice"
	// Version of the service
	Version = "0.0.1"
)

// ListenAddr is the bind address
var ListenAddr = ":9090"

func main() {
	flag.StringVar(&ListenAddr, "b", ListenAddr, "listen address")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/health", ServeHealth).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         ListenAddr,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      router,
	}

	go func() {
		fmt.Printf("%s v%s listening on %s...", Name, Version, ListenAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}

// ServeHealth writes health to the response
func ServeHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "health")
}
