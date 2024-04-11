package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/artnikel/replicatedmemorycache/internal/model"
)

type DataRepository interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type MapDataService struct {
	repository DataRepository
	peers      []string
}

func NewMapDataService(repository DataRepository, peers []string) *MapDataService {
	return &MapDataService{
		repository: repository,
		peers:      peers,
	}
}

func (s *MapDataService) Set(key, value string) error {
	if err := s.repository.Set(key, value); err != nil {
		return fmt.Errorf("set -> %w", err)
	}
	data := model.Data{Key: key, Value: value}
	s.replicateData(data)
	return nil
}

func (s *MapDataService) Get(key string) (string, error) {
	value, err := s.repository.Get(key)
	if err != nil {
		return "", fmt.Errorf("get -> %w", err)
	}
	return value, nil
}

func (s *MapDataService) Delete(key string) error {
	if err := s.repository.Delete(key); err != nil {
		return fmt.Errorf("delete -> %w", err)
	}
	return nil
}

func (s *MapDataService) replicateData(data model.Data) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling data: %v\n", err)
		return
	}

	for _, peer := range s.peers {
		if !s.isServerAvailable(peer) {
			log.Printf("Server %s is not available, skipping replication\n", peer)
			continue
		}

		_, err := http.Post(peer+"/replicate", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error replicating data to peer %s: %v\n", peer, err)
		}
	}
}

func (s *MapDataService) isServerAvailable(serverURL string) bool {
	client := http.Client{
		Timeout: time.Second * 2, 
	}
	_, err := client.Get(serverURL)
	return err == nil
}
