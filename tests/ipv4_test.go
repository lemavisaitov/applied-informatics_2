package tests

import (
	"github.com/lemavisaitov/applied-informatics_2/internal/validator"
	"testing"
)

func TestIPv4Validator_Validate(t *testing.T) {
	v := validator.NewIPv4Validator()

	validIPs := []string{
		"192.168.0.1",
		"255.255.255.255",
		"0.0.0.0",
		"127.0.0.1",
	}

	invalidIPs := []string{
		"256.256.256.256",
		"192.168.1",
		"abc.def.ghi.jkl",
		"192.168.0.999",
		"1234.567.89.0",
	}

	for _, ip := range validIPs {
		if !v.Validate(ip) {
			t.Errorf("Expected valid IPv4 address, but got invalid: %s", ip)
		}
	}

	for _, ip := range invalidIPs {
		if v.Validate(ip) {
			t.Errorf("Expected invalid IPv4 address, but got valid: %s", ip)
		}
	}
}

func TestIPv4Validator_FindMatches(t *testing.T) {
	v := validator.NewIPv4Validator()

	input := "Valid IPs: 192.168.0.1, 10.0.0.1. Invalid: 999.999.999.999, 300.300.300.300."
	expected := []string{"192.168.0.1", "10.0.0.1"}

	matches := v.FindMatches(input)

	if len(matches) != len(expected) {
		t.Errorf("Expected %d matches, but got %d", len(expected), len(matches))
	}

	for i, match := range matches {
		if match != expected[i] {
			t.Errorf("Expected match %s, but got %s", expected[i], match)
		}
	}
}
