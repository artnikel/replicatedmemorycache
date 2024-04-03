package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type CacheService interface {
	Set(key string, value interface{}, duration time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type CacheHandler struct {
	service CacheService
}

func NewCacheHandler(service CacheService) *CacheHandler {
	return &CacheHandler{service: service}
}

func (h *CacheHandler) HandleGetCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/cache/"):]
	value, err := h.service.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if value == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%v", value)
}

func (h *CacheHandler) HandleSetCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/cache/"):]
	value := r.FormValue("value")
	durationStr := r.FormValue("duration")
	duration, err := strconv.ParseInt(durationStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Set(key, value, time.Duration(duration))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CacheHandler) HandleDeleteCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/cache/"):]
	err := h.service.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
