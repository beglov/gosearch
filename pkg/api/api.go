package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"net/http"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
	engine *engine.Service
}

// New - конструктор.
func New(engine *engine.Service) *API {
	s := API{
		router: mux.NewRouter(),
		engine: engine,
	}
	s.endpoints()
	return &s
}

// Start запускает сетевую службу
func (api *API) Start(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) endpoints() {
	api.router.HandleFunc("/index", api.indexHandler)
	api.router.HandleFunc("/docs", api.docsHandler)

	api.router.HandleFunc("/api/v1/search/{word}", api.search).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/create", api.createDocs).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) search(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	docs := api.engine.Search(word)
	err := json.NewEncoder(w).Encode(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) createDocs(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.engine.Add([]crawler.Document{doc})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) indexHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(api.engine.Index)
	w.Write(payload)
}

func (api *API) docsHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(api.engine.Storage)
	w.Write(payload)
}
