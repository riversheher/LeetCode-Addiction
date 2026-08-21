package add_two_numbers

import (
	"testing"
)

// Test normal case
func TestNormalCase(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	var expected int

	//[2,4,3]
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	//[5,6,4]
	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	expected = 907

	result := addTwoNumbers(l1, l2)

	if expected != result.toInt() {
		t.Errorf("Expected: %q, Result: %q", expected, result.toInt())
	}
}

// Test order
func TestOrder(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode

	//[2,4,3]
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	//[5,6,4]
	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	resultReverse := addTwoNumbers(l2, l1)

	if result.toInt() != resultReverse.toInt() {
		t.Errorf("Order should not matter when adding: l1 + l2 = %q, l2 + l1 = %q", result.toInt(), resultReverse.toInt())
	}
}

func TestZero(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode

	l1 = &ListNode{
		Val: 0,
	}

	l2 = &ListNode{
		Val: 0,
	}

	result := addTwoNumbers(l1, l2)
	expected := 0

	if result.toInt() != expected {
		t.Errorf("Expected: %q, Result: %q", expected, result.toInt())
	}
}

func TestDifferentLength(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode

	l1 = &ListNode{
		Val: 9,
	}

	//[5,6,4]
	l2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	expected := 9 + 999

	if result.toInt() != expected {
		t.Errorf("Expected: %q, Result: %q", expected, result.toInt())
	}
}
