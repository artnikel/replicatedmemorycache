package service

import "time"

type CacheRepository interface {
	Set(key string, value interface{}, duration time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

// CacheService определяет бизнес-логику для работы с кэшем.
type CacheService struct {
	repo CacheRepository
}

func NewCacheService(repo CacheRepository) *CacheService {
	return &CacheService{repo: repo}
}

// Set добавляет или обновляет элемент в кэше.
func (s *CacheService) Set(key string, value interface{}, duration time.Duration) error {
	return s.repo.Set(key, value, duration)
}

// Get получает значение элемента из кэша по ключу.
func (s *CacheService) Get(key string) (interface{}, error) {
	return s.repo.Get(key)
}

// Delete удаляет элемент из кэша по ключу.
func (s *CacheService) Delete(key string) error {
	return s.repo.Delete(key)
}
