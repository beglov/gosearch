package webapp

import (
	"encoding/json"
	"gosearch/pkg/engine"
	"net/http"
)

// Service - сетевая служба
type Service struct {
	addr   string
	engine *engine.Service
}

// New - конструктор.
func New(addr string, engine *engine.Service) *Service {
	s := Service{
		addr:   addr,
		engine: engine,
	}
	return &s
}

// Start запускает сетевую службу
func (s *Service) Start() error {
	s.endpoints()

	return http.ListenAndServe(s.addr, nil)
}

func (s *Service) endpoints() {
	http.HandleFunc("/index", s.indexHandler)
	http.HandleFunc("/docs", s.docsHandler)
}

func (s *Service) indexHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(s.engine.Index)
	w.Write(payload)
}

func (s *Service) docsHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(s.engine.Storage)
	w.Write(payload)
}
