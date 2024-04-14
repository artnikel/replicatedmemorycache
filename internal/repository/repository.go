package repository

import (
    "sync"
)

type CacheRepository struct {
    data map[string]string
    mu   sync.RWMutex
}

func NewCacheRepository() *CacheRepository {
    return &CacheRepository{data: make(map[string]string)}
}

func (r *CacheRepository) Set(key, value string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.data[key] = value
}

func (r *CacheRepository) Get(key string) string {
    r.mu.RLock()
    defer r.mu.RUnlock()
    return r.data[key]
}

func (r *CacheRepository) Sync(data map[string]string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    for key, value := range data {
        r.data[key] = value
    }
}