package repository

import (
	"fmt"
	"sync"
)


type KeyValueStore struct {
	dataMap sync.Map
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		dataMap: sync.Map{},
	}
}

func (r *KeyValueStore) Set(key, value string) error {
	r.dataMap.Store(key, value)
	return nil
}

func (r *KeyValueStore) Get(key string) (string, error) {
	if value, ok := r.dataMap.Load(key); ok {
		return value.(string), nil
	}
	return "", fmt.Errorf("key no found")
}

func (r *KeyValueStore) Delete(key string) error {
	if _, ok := r.dataMap.Load(key); !ok {
		return fmt.Errorf("key no found")
	}
	r.dataMap.Delete(key)
	return nil
}
