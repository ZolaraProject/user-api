package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	s "github.com/ZolaraProject/user-api/userapiserver"
)

const (
	defaultExposePort = "8080"
)

func main() {
	var ok bool
	s.PkiVaultServiceHost, ok = os.LookupEnv("PKI_VAULT_SERVICE_HOST")
	if !ok {
		log.Fatal("Error: could not read $PKI_VAULT_SERVICE_HOST")
	}
	s.PkiVaultServicePort, ok = os.LookupEnv("PKI_VAULT_SERVICE_PORT")
	if !ok {
		log.Fatal("Error: could not read $PKI_VAULT_SERVICE_PORT")
	}

	exposePort, ok := os.LookupEnv("EXPOSE_PORT")
	if !ok {
		exposePort = defaultExposePort
	}

	log.Printf(fmt.Sprintf("Server listens on port %v", exposePort))

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM) // Received after the preStop hook

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", exposePort),
		Handler: s.NewRouter(),
	}

	go server.ListenAndServe()

	select {
	case c := <-termChan:
		log.Printf("Received signal %v, stopping gracefully", c)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		server.Shutdown(ctx)
		log.Printf("Server stopped, exiting. ")
	}
}
