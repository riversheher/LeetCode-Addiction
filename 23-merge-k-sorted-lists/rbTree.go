package mergeksortedlists

// Properties:
// 1. Every node is either red or black.
// 2. The root is always black.
// 3. Red nodes cannot have red children (no two red nodes can be adjacent).
// 4. Every path from a node to its descendant leaves must have the same number of black nodes. (black height)
// 5. Every leaf (NIL node) is black.
// Black depth is the number of black ancestors from a node to the root.
type RedBlackTree struct {
	Root *Node
}

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
	Color  bool // true for red, false for black
}

func NewTree() *RedBlackTree {
	return &RedBlackTree{
		Root: nil, // Initialize with an empty tree
	}
}

// IsLeftChild returns true if the node is a left child of its parent.
func (n *Node) IsLeftChild() bool {
	if n.Parent == nil {
		return false // Root node has no parent
	}
	return n.Parent.Left == n
}

// Is RightChild returns true if the node is a right child of its parent.
func (n *Node) IsRightChild() bool {
	if n.Parent == nil {
		return false // Root node has no parent
	}
	return n.Parent.Right == n
}

// GetUncle returns the uncle of a given node.
func (n *Node) GetUncle() *Node {
	if n.Parent == nil || n.Parent.Parent == nil {
		return nil // No uncle if parent or grandparent is nil
	}
	if n.Parent == n.Parent.Parent.Left {
		return n.Parent.Parent.Right
	}
	return n.Parent.Parent.Left
}

// Left Rotation rotates the subtree at node x to the left.
// Returns the new root of the subtree after rotation.
func (tree *RedBlackTree) LeftRotate(x *Node) *Node {
	if x.Right == nil {
		return x // Cannot perform left rotation if right child is nil
	}

	y := x.Right // y is the right child of x

	// Set y's left child to x's right child
	x.Right = y.Left

	// Set parents
	y.Parent = x.Parent
	// If y is not the root, update the parent's child pointer
	if y.Parent != nil {
		if x == y.Parent.Left {
			y.Parent.Left = y
		} else {
			y.Parent.Right = y
		}
	} else {
		// If y is the root, update the tree's root
		tree.Root = y
	}
	// Set x's parent to y
	x.Parent = y

	// Set y's left child to x
	y.Left = x

	return y // Return the new root of the subtree
}

// Right rotation rotates the subtree at node x to the right.
// Returns the new root of the subtree
func (tree *RedBlackTree) RightRotate(x *Node) *Node {
	if x.Left == nil {
		return x // Cannot perform right rotation if left child is nil
	}

	y := x.Left // y is the left child of x

	// Set y's right child to x's left child
	x.Left = y.Right

	// Set parents
	y.Parent = x.Parent
	// If y is not the root, update the parent's child pointer
	if y.Parent != nil {
		if x == y.Parent.Right {
			y.Parent.Right = y
		} else {
			y.Parent.Left = y
		}
	} else {
		// If y is the root, update the tree's root
		tree.Root = y
	}
	// Set x's parent to y
	x.Parent = y

	// Set y's right child to x
	y.Right = x

	return y // Return the new root of the subtree
}

// Fix checks the red-black properties of the tree and fixes any violations.
// This function is called after insertion to ensure the tree remains balanced.
func (tree *RedBlackTree) Fix(n *Node) {
	if n == nil {
		return
	}
	if n.Parent == nil {
		// If n is the root, it must be black
		n.Color = false
		return
	}

	if n.Parent.Color {
		// If the parent is red, we have a violation
		uncle := n.GetUncle()
		if uncle != nil && uncle.Color {
			// Case 1: Uncle is red
			// Recolor parent and uncle to black, and grandparent to red
			n.Parent.Color = false
			uncle.Color = false
			if n.Parent.Parent != nil {
				n.Parent.Parent.Color = true
			}
			// Recurse up the tree to fix any further violations
			tree.Fix(n.Parent.Parent)
		} else {
			// Case 2: Uncle is black (or nil)
			if n.IsRightChild() && n.Parent.IsLeftChild() {
				// Case 2a: n is a right child of a left parent, rotate left on parent, then rotate right on grandparent (new parent of n)
				grandparent := n.Parent.Parent

				tree.LeftRotate(n.Parent)
				tree.RightRotate(grandparent)

				// recolour
				grandparent.Color = true // Grandparent becomes red
				n.Color = false          // n becomes black
			} else if n.IsLeftChild() && n.Parent.IsRightChild() {
				// Case 2b: n is a left child of a right parent, rotate right on parent
				grandparent := n.Parent.Parent
				tree.RightRotate(n.Parent)
				tree.LeftRotate(grandparent)

				// recolour
				grandparent.Color = true // Grandparent becomes red
				n.Color = false          // n becomes black
			} else if n.IsLeftChild() && n.Parent.IsLeftChild() {
				// case 2c: n is a left child of a left parent, rotate right on grandparent
				grandparent := n.Parent.Parent
				tree.RightRotate(grandparent)

				grandparent.Color = true         // Grandparent becomes red
				grandparent.Parent.Color = false // Parent of grandparent becomes black

			}
		}
	}
}

func (root *Node) Insert(value int) {
	newNode := &Node{
		Value: value,
		Color: true, // New nodes are always red
	}

	if root == nil {
		// If the tree is empty, set the new node as the root
		root = newNode
		newNode.Color = false // Root must be black
		return
	}

	// Insert node at correct point
	current := root
	for {
		if value <= current.Value {
			// Insert to the left if there is no left child
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				break
			}
			// otherwise, continue down the left subtree
			current = current.Left
		} else {
			//left rotation Insert to the right if there is no right child
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				break
			}
			// otherwise, continue down the right subtree
			current = current.Right
		}
	}

	//Check for red black violations
	if newNode.Parent.Color {
		//Parent is red, need to fix the tree
		if newNode.GetUncle() != nil && newNode.GetUncle().Color {
			//Perform recolours and recurse up the tree
			newNode.Parent.Color = false     // Change parent to black
			newNode.GetUncle().Color = false // Change uncle to black
			if newNode.Parent.Parent != nil {
				newNode.Parent.Parent.Color = true // Change grandparent to red
			}
			// Recurse up the tree
		} else if newNode.Parent.Left != nil && newNode == newNode.Parent.Left {
			//Left child of parent, perform right rotation on grandparent and recolour
		} else if newNode.Parent.Right != nil && newNode == newNode.Parent.Right {
			//Right child of parent, perform left rotation on parent
		}
	}
}
