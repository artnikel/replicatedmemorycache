package service

import "github.com/artnikel/replicatedmemorycache/internal/repository"

type CacheService struct {
    repository *repository.CacheRepository
}

func NewCacheService(repo *repository.CacheRepository) *CacheService {
    return &CacheService{repository: repo}
}

func (s *CacheService) SetData(key, value string) {
    s.repository.Set(key, value)
}

func (s *CacheService) GetData(key string) string {
    return s.repository.Get(key)
}

func (s *CacheService) SyncData(data map[string]string) {
    s.repository.Sync(data)
}
