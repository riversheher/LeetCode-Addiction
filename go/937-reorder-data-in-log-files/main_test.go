package main

import (
	"reflect"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {
	tests := []struct {
		name     string
		logs     []string
		expected []string
	}{
		{
			name: "Example 1 from LeetCode",
			logs: []string{
				"dig1 8 1 5 1",
				"let1 art can",
				"dig2 3 6",
				"let2 own kit dig",
				"let3 art zero",
			},
			expected: []string{
				"let1 art can",
				"let3 art zero",
				"let2 own kit dig",
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
		},
		{
			name: "Example 2 from LeetCode",
			logs: []string{
				"a1 9 2 3 1",
				"g1 act car",
				"zo4 4 7",
				"ab1 off key dog",
				"a8 act zoo",
			},
			expected: []string{
				"g1 act car",
				"a8 act zoo",
				"ab1 off key dog",
				"a1 9 2 3 1",
				"zo4 4 7",
			},
		},
		{
			name: "Same letter-log content, different identifiers",
			logs: []string{
				"let2 art can",
				"let1 art can",
			},
			expected: []string{
				"let1 art can",
				"let2 art can",
			},
		},
		{
			name: "All digit-logs",
			logs: []string{
				"dig1 8 1",
				"dig2 3 6",
			},
			expected: []string{
				"dig1 8 1",
				"dig2 3 6",
			},
		},
		{
			name: "All letter-logs",
			logs: []string{
				"let2 own kit dig",
				"let1 art can",
			},
			expected: []string{
				"let1 art can",
				"let2 own kit dig",
			},
		},
		{
			name: "Single log",
			logs: []string{
				"let1 art can",
			},
			expected: []string{
				"let1 art can",
			},
		},
		{
			name: "Digit logs relative order preservation",
			logs: []string{
				"dig2 3 6",
				"dig1 8 1 5 1",
			},
			expected: []string{
				"dig2 3 6",
				"dig1 8 1 5 1",
			},
		},
		{
			name: "Mixed with letter-logs first in input but should still sort correctly",
			logs: []string{
				"let1 art can",
				"dig1 8 1",
				"let2 own kit dig",
			},
			expected: []string{
				"let1 art can",
				"let2 own kit dig",
				"dig1 8 1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reorderLogFiles(tt.logs)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("reorderLogFiles(%v) = %v; want %v", tt.logs, result, tt.expected)
			}
		})
	}
}
