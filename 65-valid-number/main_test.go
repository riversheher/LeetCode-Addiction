package validnumber

import (
	"testing"
)

// Test normal case
func TestNormalCase(t *testing.T) {
	result := isNumber("-1E+3")

	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test Bad case
func TestBadCase(t *testing.T) {
	result := isNumber("6+1")

	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test Bad case
func TestBadCase2(t *testing.T) {
	result := isNumber(".+6")

	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test Bad case
func TestBadCase3(t *testing.T) {
	result := isNumber("6E6.5")

	if result {
		t.Errorf("Expected false, got true")
	}
}
