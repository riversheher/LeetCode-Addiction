package validatebst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Limit struct {
	Val int
}

func isValidBST(root *TreeNode) bool {

	if root.Left == nil && root.Right == nil {
		return true
	}

	limit := &Limit{Val: root.Val}

	if root.Left == nil && root.Right != nil {
		return isValidSubTree(root.Right, limit, nil)
	} else if root.Left != nil && root.Right == nil {
		return isValidSubTree(root.Left, nil, limit)
	} else {
		return isValidSubTree(root.Right, limit, nil) && isValidSubTree(root.Left, nil, limit)
	}

}

/*
Checks if the tree is valid by comparing its value with the minimum and maximum allowed.
If Limit is nil, there is no limit.
*/
func isValidSubTree(root *TreeNode, min *Limit, max *Limit) bool {

	if max != nil && root.Val >= max.Val {
		return false
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left != nil && root.Right == nil {
		newMax := &Limit{
			Val: root.Val,
		}
		return isValidSubTree(root.Left, min, newMax)
	} else if root.Right != nil && root.Left == nil {
		newMin := &Limit{
			Val: root.Val,
		}
		return isValidSubTree(root.Right, newMin, max)
	} else {
		newMax := &Limit{
			Val: root.Val,
		}
		newMin := &Limit{
			Val: root.Val,
		}
		return isValidSubTree(root.Left, min, newMax) && isValidSubTree(root.Right, newMin, max)
	}
}
