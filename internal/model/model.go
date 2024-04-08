package model

import (
	"sync"

	"github.com/hashicorp/memberlist"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Server struct {
	DataMap      map[string]string
	PeerList     []*memberlist.Node
	Memberlist   *memberlist.Memberlist
	PeerUpdateMu sync.Mutex
}
