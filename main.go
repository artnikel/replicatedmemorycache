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
	"github.com/hashicorp/memberlist"
)

func main() {
	repository := repository.NewMapDataRepository()
	peerList := []*memberlist.Node{} 
	replicationService := service.NewDataReplicationService(peerList)
	service := service.NewMapDataService(repository,*replicationService)
	handler := handler.NewDataHandler(service)

	http.HandleFunc("/set", handler.Set)
	http.HandleFunc("/get", handler.Get)

	log.Println("Server started on :8080")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
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
