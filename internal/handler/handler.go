package handler

import (
	"fmt"
	"net/http"
)

type DataService interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type DataHandler struct {
	service DataService
}

func NewDataHandler(service DataService) *DataHandler {
	return &DataHandler{
		service: service,
	}
}

func (h *DataHandler) Set(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if err := h.service.Set(key, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data added successfully")
}

func (h *DataHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")

	data, err := h.service.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Value for key %s: %s", key, data)
}
