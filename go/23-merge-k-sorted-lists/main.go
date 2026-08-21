package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// equal checks if two linked lists are equal.
func equals(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

// mergeKLists merges k sorted linked lists and returns the merged sorted list.
// Solve this problem with a red-black tree, probably best due to number of insertions
// and should be easy to get the sorted list from the tree at the end.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	head := &ListNode{}

	return head

}
