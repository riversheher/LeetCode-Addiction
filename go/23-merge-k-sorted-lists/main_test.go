package mergeksortedlists

import (
	"testing"
)

func TestLeftRotate(t *testing.T) {
	tree := NewTree()

	root := &Node{
		Value: 10,
		Color: false, // Black
	}

	tree.Root = root

	x := root
	y := &Node{
		Value: 20,
		Color: true,                           // Red
		Left:  &Node{Value: 15, Color: false}, // Black
		Right: &Node{Value: 25, Color: false}, // Black
	}
	x.Right = y
	y.Parent = x

	newRoot := tree.LeftRotate(x)

	if newRoot != y {
		t.Errorf("Expected new root to be %v, got %v", y.Value, newRoot.Value)
	}
	if newRoot.Left != x {
		t.Errorf("Expected new root's left child to be %v, got %v", x.Value, newRoot.Left.Value)
	}
	if x.Right.Value != 15 {
		t.Errorf("Expected x's right child to be 15, got %v", x.Right.Value)
	}
	if x.Parent != y {
		t.Errorf("Expected x's parent to be %v, got %v", y.Value, x.Parent.Value)
	}
	if y.Parent != nil {
		t.Errorf("Expected y's parent to be nil, got %v", y.Parent.Value)
	}
	if tree.Root != y {
		t.Errorf("Expected tree root to be %v, got %v", y.Value, tree.Root.Value)
	}

}

func TestRightRotate(t *testing.T) {
	tree := NewTree()

	root := &Node{
		Value: 10,
		Color: false, // Black
	}

	tree.Root = root

	x := root
	y := &Node{
		Value: 5,
		Color: true,                          // Red
		Left:  &Node{Value: 2, Color: false}, // Black
		Right: &Node{Value: 8, Color: false}, // Black
	}
	x.Left = y
	y.Parent = x

	newRoot := tree.RightRotate(x)

	if newRoot != y {
		t.Errorf("Expected new root to be %v, got %v", y.Value, newRoot.Value)
	}
	if newRoot.Right != x {
		t.Errorf("Expected new root's left child to be %v, got %v", x.Value, newRoot.Left.Value)
	}
	if x.Left.Value != 8 {
		t.Errorf("Expected x's right child to be 15, got %v", x.Right.Value)
	}
	if x.Parent != y {
		t.Errorf("Expected x's parent to be %v, got %v", y.Value, x.Parent.Value)
	}
	if y.Parent != nil {
		t.Errorf("Expected y's parent to be nil, got %v", y.Parent.Value)
	}
	if tree.Root != y {
		t.Errorf("Expected tree root to be %v, got %v", y.Value, tree.Root.Value)
	}

}

func TestFixCase1(t *testing.T) {
	tree := NewTree()

	root := &Node{
		Value: 17,
		Color: false, // Black
	}

	tree.Root = root

	L := &Node{
		Value:  9,
		Color:  false, // Black
		Parent: root,
	}

	R := &Node{
		Value:  19,
		Color:  false, // Black
		Parent: root,
	}

	RR := &Node{
		Value:  75,
		Color:  true, // Red
		Parent: R,
	}

	RRL := &Node{
		Value:  24,
		Color:  true, // Red
		Parent: RR,
	}

	root.Left = L
	root.Right = R
	R.Left = RR
	RR.Left = RRL

	tree.Fix(RRL)

	if tree.Root != root {
		t.Errorf("Expected root to be %v, got %v", root.Value, tree.Root.Value)
	}
	if root.Color != false {
		t.Errorf("Expected root color to be black, got %v", root.Color)
	}
	if root.Left != L {
		t.Errorf("Expected root left child to be %v, got %v", L.Value, root.Left.Value)
	}
	if root.Right != RR {
		t.Errorf("Expected root right child to be %v, got %v", R.Value, root.Right.Value)
	}
	if root.Right.Color != false {
		t.Errorf("Expected RR color to be black, got %v", root.Right.Color)
	}
	if root.Right.Left != R {
		t.Errorf("Expected RR left child to be %v, got %v", R.Value, root.Right.Left.Value)
	}
	if root.Right.Left.Color != true {
		t.Errorf("Expected R color to be red, got %v", root.Right.Left.Color)
	}
}

func TestNoLists(t *testing.T) {
	lists := []*ListNode{}
	expected := (*ListNode)(nil)
	result := mergeKLists(lists)
	if result != expected {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestSingleList(t *testing.T) {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
	}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	result := mergeKLists(lists)
	if !equals(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMultipleLists(t *testing.T) {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	b := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	c := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	lists := []*ListNode{a, b, c}

	expected := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}}

	result := mergeKLists(lists)

	if !equals(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
