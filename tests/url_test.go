package tests

import (
	"github.com/lemavisaitov/applied-informatics_2/internal/fetcher"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLFetcher_Fetch(t *testing.T) {
	fetcherv := &fetcher.URLFetcher{}

	// Создание тестового сервера
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a test page with IP: 10.0.0.1"))
	}))
	defer server.Close()

	// Тест запроса к серверу
	result, err := fetcherv.Fetch(server.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	expected := "This is a test page with IP: 10.0.0.1"
	if result != expected {
		t.Errorf("Expected content '%s', but got '%s'", expected, result)
	}
}
