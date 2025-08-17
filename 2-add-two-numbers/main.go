package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) toInt() int {
	if list == nil {
		return 0
	}

	value := 0
	for i, curr := 1, list; curr != nil; i, curr = i*10, curr.Next {
		value += curr.Val * i
	}

	return value
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var startNode *ListNode

	//Iterate over list nodes, storing carry over
	left := l1
	right := l2
	currNode := startNode
	carry := 0
	for left != nil || right != nil || carry != 0 {
		value := carry
		if left != nil {
			value += left.Val
			left = left.Next
		}
		if right != nil {
			value += right.Val
			right = right.Next
		}

		//Results
		resultNode := &ListNode{
			Val: value % 10,
		}
		carry = value / 10

		//Set results
		if currNode != nil {
			currNode.Next = resultNode
		} else {
			startNode = resultNode
		}

		currNode = resultNode

	}

	return startNode
}
