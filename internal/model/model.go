package model

import (
	"sync"
)

type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration int64
}

type Cache struct {
	Items map[string]CacheItem
	Mu    sync.RWMutex
}
