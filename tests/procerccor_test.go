package tests

import (
	"github.com/lemavisaitov/applied-informatics_2/internal/processor"
	"github.com/lemavisaitov/applied-informatics_2/internal/validator"
	"testing"
)

type MockFetcher struct {
	Data string
}

func (m *MockFetcher) Fetch(source string) (string, error) {
	return m.Data, nil
}

func TestProcessor_Process(t *testing.T) {
	validatorv := validator.NewIPv4Validator()
	mockFetcher := &MockFetcher{Data: "Test data with IP: 192.168.0.1 and 10.0.0.1"}

	proc := processor.NewProcessor(mockFetcher, validatorv)
	matches, err := proc.Process("mock_source")

	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	expected := []string{"192.168.0.1", "10.0.0.1"}
	if len(matches) != len(expected) {
		t.Errorf("Expected %d matches, but got %d", len(expected), len(matches))
	}

	for i, match := range matches {
		if match != expected[i] {
			t.Errorf("Expected match '%s', but got '%s'", expected[i], match)
		}
	}
}
