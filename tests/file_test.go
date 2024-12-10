package tests

import (
	"github.com/lemavisaitov/applied-informatics_2/internal/fetcher"
	"os"
	"testing"
)

func TestFileFetcher_Fetch(t *testing.T) {
	fetcherv := &fetcher.FileFetcher{}

	// Создание временного файла
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	content := "This is a test file with IP: 192.168.0.1"
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	tempFile.Close()

	// Тест чтения файла
	result, err := fetcherv.Fetch(tempFile.Name())
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if result != content {
		t.Errorf("Expected content '%s', but got '%s'", content, result)
	}
}
