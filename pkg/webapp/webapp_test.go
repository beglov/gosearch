package webapp

import (
	"gosearch/pkg/engine"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage/btree"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var srv *Service

func TestMain(m *testing.M) {
	srv := New(":8000", engine.New(hash.New(), btree.New()))
	srv.endpoints()

	os.Exit(m.Run())
}

func TestService_indexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)

	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}

func TestService_docsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}
