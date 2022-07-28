package pkg

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Preorder(root *BinaryTreeNode) (result []int) {
	result = make([]int, 0)

	if root == nil {
		return
	}

	stack := []*BinaryTreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		result = append(result, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return
}

func Inorder(root *BinaryTreeNode) (result []int) {
	result = make([]int, 0)

	if root == nil {
		return
	}

	stack := []*BinaryTreeNode{}
	cursor := root

	for len(stack) != 0 || cursor != nil {
		if cursor != nil {
			stack = append(stack, cursor)
			cursor = cursor.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			result = append(result, top.Val)

			cursor = top.Right
		}
	}

	return
}

func Postorder(root *BinaryTreeNode) (result []int) {
	result = make([]int, 0)

	if root == nil {
		return
	}

	stack := []*BinaryTreeNode{}
	var preNode *BinaryTreeNode
	cursor := root

	for len(stack) > 0 || cursor != nil {
		if cursor != nil {
			stack = append(stack, cursor)
			cursor = cursor.Left
		} else {
			top := stack[len(stack)-1]
			if top.Right != nil {
				if top.Right == preNode {
					stack = stack[0 : len(stack)-1]
					preNode = top
					result = append(result, top.Val)
				} else {
					cursor = top.Right
				}
			} else {
				stack = stack[0 : len(stack)-1]
				preNode = top
				result = append(result, top.Val)
			}
		}
	}

	return
}
