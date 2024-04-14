package handler

import (
	"context"

	"github.com/artnikel/replicatedmemorycache/internal/service"
	pb "github.com/artnikel/replicatedmemorycache/proto"
)

type CacheHandler struct {
    service *service.CacheService
	pb.UnimplementedCacheServiceServer
}

func NewCacheHandler(service *service.CacheService) *CacheHandler {
    return &CacheHandler{service: service}
}

func (h *CacheHandler) SetData(ctx context.Context, req *pb.SetDataRequest) (*pb.SetDataResponse, error) {
    h.service.SetData(req.Key, req.Value)
    return &pb.SetDataResponse{Success: true}, nil
}

func (h *CacheHandler) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
    value := h.service.GetData(req.Key)
    return &pb.GetDataResponse{Value: value}, nil
}

func (h *CacheHandler) SyncData(ctx context.Context, req *pb.SyncDataRequest) (*pb.SyncDataResponse, error) {
    h.service.SyncData(req.Data)
    return &pb.SyncDataResponse{Success: true}, nil
}