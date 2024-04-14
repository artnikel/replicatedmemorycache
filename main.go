package main

import (
	"log"
	"net"

	"github.com/artnikel/replicatedmemorycache/internal/handler"
	"github.com/artnikel/replicatedmemorycache/internal/repository"
	"github.com/artnikel/replicatedmemorycache/internal/service"
	"github.com/artnikel/replicatedmemorycache/proto"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewCacheRepository()
	service := service.NewCacheService(repo)
    handler := handler.NewCacheHandler(service)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCacheServiceServer(s, handler)

	log.Println("Cache server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
