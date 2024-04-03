package repository

import (
	"sync"
	"time"

	"github.com/artnikel/replicatedmemorycache/internal/model"
)

type InMemoryCacheRepository struct {
	cache map[string]model.CacheItem
	mu    sync.RWMutex
}

func NewInMemoryCacheRepository() *InMemoryCacheRepository {
	return &InMemoryCacheRepository{cache: make(map[string]model.CacheItem)}
}

func (repo *InMemoryCacheRepository) Set(key string, value interface{}, duration time.Duration) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	expiration := time.Now().Add(duration).UnixNano()
	repo.cache[key] = model.CacheItem{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	return nil
}

func (repo *InMemoryCacheRepository) Get(key string) (interface{}, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	item, found := repo.cache[key]
	if !found || item.Expiration < time.Now().UnixNano() {
		return nil, nil
	}
	return item.Value, nil
}

func (repo *InMemoryCacheRepository) Delete(key string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.cache, key)
	return nil
}
