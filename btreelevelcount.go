package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		var result int = BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right)
		if root.Left != nil && root.Right != nil {
			return result
		} else {
			return result + 1
		}
	}
	return 0
}
