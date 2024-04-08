package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/memberlist"
)

type DataReplicationService struct {
	peerList []*memberlist.Node
}

func NewDataReplicationService(peerList []*memberlist.Node) *DataReplicationService {
	return &DataReplicationService{
		peerList: peerList,
	}
}

func (s *DataReplicationService) ReplicateData(key, value string) error {
	data := map[string]string{"key": key, "value": value}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, peer := range s.peerList {
		resp, err := http.Post(fmt.Sprintf("http://%s:%d/replicate", peer.Addr.String(), peer.Port), "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("error replicating data to %s: %v", peer.Name, err)
		}
		resp.Body.Close()
	}

	return nil
}
