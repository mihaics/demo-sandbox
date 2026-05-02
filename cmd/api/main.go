package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mihaics/demo-sandbox/internal/server"
)

const (
	defaultAddr     = ":8080"
	shutdownTimeout = 15 * time.Second
)

func main() {
	addr := os.Getenv("PORT")
	switch {
	case addr == "":
		addr = defaultAddr
	case addr[0] != ':':
		addr = ":" + addr
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           server.New(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server error: %v", err)
		}
	case sig := <-stop:
		log.Printf("received %s, shutting down", sig)
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("graceful shutdown failed: %v", err)
		}
		log.Print("shutdown complete")
	}
}
