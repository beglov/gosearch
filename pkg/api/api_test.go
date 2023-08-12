package api

import (
	"bytes"
	"encoding/json"
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage/memstore"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	index := &hash.Index{
		Data: map[string][]int{
			"word1": {0},
			"word2": {1, 2},
		},
	}
	storage := memstore.New()
	docs := []crawler.Document{
		{ID: 0, URL: "url1", Title: "word1"},
		{ID: 1, URL: "url2", Title: "word2"},
		{ID: 2, URL: "url3", Title: "word2"},
	}
	storage.StoreDocs(docs)
	engine := engine.New(index, storage)

	api = New(engine)
	os.Exit(m.Run())
}

func TestAPI_search(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/search/word2", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	wantDocs := []crawler.Document{
		{ID: 1, URL: "url2", Title: "word2"},
		{ID: 2, URL: "url3", Title: "word2"},
	}
	wantBytes, _ := json.Marshal(wantDocs)
	want := string(wantBytes)
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != want {
		t.Errorf("содержимое некорректно: получили %s а хотели %s", got, want)
	}
}

func TestAPI_createDocs(t *testing.T) {
	doc := crawler.Document{
		ID:    999,
		URL:   "url9",
		Title: "word9",
		Body:  "",
	}
	payload, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs/create", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_indexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}

func TestAPI_docsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}
