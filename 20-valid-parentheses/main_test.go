package validparentheses

import (
	"testing"
)

// Test normal case
func TestNormalCase(t *testing.T) {
	result := isValid("()")

	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test nested parentheses
func TestNestedParentheses(t *testing.T) {
	result := isValid("([])")

	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test multiple types of parentheses
func TestMultipleTypes(t *testing.T) {
	result := isValid("{[()]}")

	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test unbalanced parentheses
func TestUnbalancedParentheses(t *testing.T) {
	result := isValid("([)]")

	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test empty string
func TestEmptyString(t *testing.T) {
	result := isValid("")

	if !result {
		t.Errorf("Expected true, got false")
	}
}
