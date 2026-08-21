package longestdiameterbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxD := 0
	var dfsVisit func(*TreeNode) int

	dfsVisit = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftLine := dfsVisit(node.Left)
		rightLine := dfsVisit(node.Right)

		maxD = max(leftLine+rightLine, maxD)

		return max(leftLine, rightLine) + 1 // 1 for ur curr edge
		// so its previous plus current
	}

	dfsVisit(root)

	return maxD
}
