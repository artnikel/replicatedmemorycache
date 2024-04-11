package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/artnikel/replicatedmemorycache/internal/handler"
	"github.com/artnikel/replicatedmemorycache/internal/repository"
	"github.com/artnikel/replicatedmemorycache/internal/service"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Enter the port at the end")
	}
	port := args[1]
	peers := []string{"http://localhost:8081", "http://localhost:8082"}
	repository := repository.NewKeyValueStore()
	service := service.NewMapDataService(repository, peers)
	handler := handler.NewDataHandler(service)

	http.HandleFunc("/set", handler.Set)
	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/delete", handler.Delete)

	log.Printf("Server started on %s", port)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	cancel()
	log.Println("Shutting down server...")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Server shutdown timeout, force exiting.")
		os.Exit(1)
	case <-ctx.Done():
		log.Println("Server stopped gracefully.")
	}
}
