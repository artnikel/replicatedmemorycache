package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/artnikel/replicatedmemorycache/internal/handler"
	"github.com/artnikel/replicatedmemorycache/internal/model"
	"github.com/artnikel/replicatedmemorycache/internal/repository"
	"github.com/artnikel/replicatedmemorycache/internal/service"
	"github.com/hashicorp/memberlist"
)

type DistributedCache struct {
	cache  *model.Cache
	list   *memberlist.Memberlist
	config *memberlist.Config
}

func (dc *DistributedCache) joinCluster(peer string) error {
    _, err := dc.list.Join([]string{peer})
    return err
}

func NewCache() *model.Cache {
    return &model.Cache{
        Items: make(map[string]model.CacheItem),
    }
}

func newDistributedCache(port int) (*DistributedCache, error) {
	cache := NewCache()
	config := memberlist.DefaultLANConfig()
	config.BindPort = port
	config.AdvertisePort = port
	list, err := memberlist.Create(config)
	if err != nil {
		return nil, err
	}
	dc := &DistributedCache{
		cache:  cache,
		list:   list,
		config: config,
	}
	return dc, nil
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	peer := os.Getenv("PEER")

	cacheRepo := repository.NewInMemoryCacheRepository()
	cacheService := service.NewCacheService(cacheRepo)
	cacheHandler := handler.NewCacheHandler(cacheService)

	http.HandleFunc("/cache/", cacheHandler.HandleGetCache)
	http.HandleFunc("/cache/set", cacheHandler.HandleSetCache)
	http.HandleFunc("/cache/delete", cacheHandler.HandleDeleteCache)

	dc, err := newDistributedCache(port)
    if err != nil {
        log.Fatalf("Failed to create distributed cache: %v", err)
    }
    if peer != "" {
        err = dc.joinCluster(peer)
        if err != nil {
            log.Fatalf("Failed to join cluster: %v", err)
        }
    }

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
