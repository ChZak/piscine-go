package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data > elem {
		BTreeSearchItem(root.Left, elem)
	}
	if root.Data < elem {
		BTreeSearchItem(root.Right, elem)
	}
	return root
}
